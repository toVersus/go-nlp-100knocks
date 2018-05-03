package main

import (
	"bufio"
	"encoding/json"
	"flag"
	"fmt"
	"io"
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

	for _, result := range articles.find(keyword).selectCategory() {
		fmt.Printf("%#v\n", result)
	}
}

// Article represents the body of Wiki page and its Title.
type Article struct {
	Title string `json:"title"`
	Text  string `json:"text"`
}

// newArticle is a constructer of Article.
func newArticle(a Article) Article {
	return Article{a.Title, a.Text}
}

// Articles represents the slice of Article.
type Articles []Article

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

var reg = regexp.MustCompile(`\[\[Category:[^]]+\]\]`)

// selectCategory returns Category description.
func (articles Articles) selectCategory() []string {
	var selected []string
	for _, a := range articles {
		for _, str := range reg.FindAllString(a.Text, -1) {
			selected = append(selected, str)
		}
	}
	return selected
}
