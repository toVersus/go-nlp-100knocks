package main

import (
	"bufio"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"regexp"
	"strings"
)

func main() {
	var (
		filePath string
		keyword  string
	)
	flag.StringVar(&filePath, "file", "", "specify a file path")
	flag.StringVar(&filePath, "f", "", "specify a file path")
	flag.StringVar(&keyword, "key", "", "specify a keyword for filtering a wiki page")
	flag.Parse()

	articles, err := readJSON(filePath)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	results, err := articles.find(keyword).getUnionFlagURL()
	if err != nil {
		fmt.Printf("Error: %s\n", err)
	}
	for _, page := range results.PageSlice() {
		for _, imageInfo := range page.ImageInfo {
			fmt.Printf("Media Wiki URL: %s\n", imageInfo.URL)
		}
	}
}

// Article represents the body of Wiki page and its Title.
type Article struct {
	Title string `json:"title"`
	Text  string `json:"text"`
}

// Articles represents the slice of Article.
type Articles []Article

// MediaWikiResponse represents response from Media Wiki API endpoint.
type MediaWikiResponse struct {
	BatchComplete string `json:"batchcomplete"`
	Query         struct {
		// Media Wiki API returns the following ugly JSON structure.
		// The workaround is to create a list of pages from the map.
		// "query":{
		// 	"pages":{
		// 		"23473560":{
		// 			"pageid":23473560
		Pages map[string]Page
	}
}

// PageSlice generates a slice from Pages to work around the ugly design of MediaWiki API
func (r *MediaWikiResponse) PageSlice() []Page {
	pages := []Page{}
	for _, page := range r.Query.Pages {
		pages = append(pages, page)
	}
	return pages
}

// Page represents a metadata of MediaWiki page.
type Page struct {
	PageID          int    `json:"pageid"`
	Ns              int    `json:"ns"`
	Title           string `json:"title"`
	ImageRepository string `json:"imagerepository"`
	ImageInfo       []struct {
		URL                 string `json:"url"`
		DescriptionURL      string `json:"descriptionurl"`
		DescriptionShortURL string `json:"descriptionshorturl"`
	}
}

// readJSON reads/parses the json file and initiates Articles instance
func readJSON(path string) (Articles, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("could not open a file specified: %s\n  %s", path, err)
	}
	defer f.Close()

	reader := bufio.NewReader(f)
	article, articles := Article{}, Articles{}
	for {
		buf, readErr := reader.ReadBytes('\n')
		if (readErr != nil) && (readErr != io.EOF) {
			return nil, fmt.Errorf("could not read a file content: %s", readErr)
		}
		if err = json.Unmarshal(buf, &article); err != nil {
			fmt.Printf("could not parse json file: %s", err)
			break
		}

		articles = append(articles, article)

		if readErr == io.EOF {
			break
		}
	}
	return articles, nil
}

// find returns Articles filtered by specified keyword.
func (articles Articles) find(keyword string) Articles {
	var filtered Articles
	for _, a := range articles {
		if strings.Contains(a.Title, keyword) {
			filtered = append(filtered, a)
		}
	}
	return filtered
}

var templateTagReg = regexp.MustCompile(`{{基礎情報 国[\s\S]*?\n}}`)
var sectionSeparatorReg = regexp.MustCompile(`(\|)+?.*(=).*(|)+?`)
var keyValueSeparatorReg = regexp.MustCompile(`(\s)?=(\s)?`)

// getUnionFlagInfo sends the GET request to the endpoint of Media Wiki API
// specifing the title of union flag image and returns the json response.
func (articles Articles) getUnionFlagURL() (*MediaWikiResponse, error) {
	var sep string
	var keyValuePair []string
	dict := make(map[string]string)

	for _, a := range articles {
		for _, line := range templateTagReg.FindAllString(a.Text, -1) {
			for _, text := range sectionSeparatorReg.FindAllString(line, -1) {
				sep = keyValueSeparatorReg.FindString(text)
				keyValuePair = strings.Split(text[1:], sep)
				dict[keyValuePair[0]] = keyValuePair[1]
			}
		}
	}

	endpoint, err := url.ParseRequestURI("https://en.wikipedia.org/w/api.php")
	if err != nil {
		return nil, err
	}

	parameters := url.Values{
		"action": []string{"query"},
		"prop":   []string{"imageinfo"},
		"format": []string{"json"},
		"iiprop": []string{"url"},
		"titles": []string{"File:" + dict["国旗画像"]},
	}

	getURL := endpoint.String() + "?" + parameters.Encode()
	res, err := http.Get(getURL)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	bodyBuf, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	imageInfo := MediaWikiResponse{}
	json.Unmarshal([]byte(bodyBuf), &imageInfo)

	return &imageInfo, err
}
