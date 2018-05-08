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

	results := articles.find(keyword).getTemplate()
	fmt.Printf("%#v\n", results)
}

// Article represents the body of Wiki page and its Title.
type Article struct {
	Title string `json:"title"`
	Text  string `json:"text"`
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

var templateTagReg = regexp.MustCompile(`{{基礎情報 国[\s\S]*?\n}}`)
var sectionSeparatorReg = regexp.MustCompile(`(\|)+?.*(=).*(|)+?`)
var keyValueSeparatorReg = regexp.MustCompile(`(\s)?=(\s)?`)
var emphaticStyleReg = regexp.MustCompile(`'{2,5}`)

// getTemplate returns name and value of template field removing emphatic style.
func (articles Articles) getTemplate() map[string]string {
	var sep string
	var keyValuePair []string
	dict := make(map[string]string)

	for _, a := range articles {
		for _, line := range templateTagReg.FindAllString(a.Text, -1) {
			for _, text := range sectionSeparatorReg.FindAllString(line, -1) {
				sep = keyValueSeparatorReg.FindString(text)
				keyValuePair = strings.Split(text[1:], sep)

				// Remove emphatic style (Italics, bold, and both) symbols.
				// https://en.wikipedia.org/wiki/Help:Cheatsheet
				dict[keyValuePair[0]] = emphaticStyleReg.ReplaceAllString(keyValuePair[1], "")
			}
		}
	}
	return dict
}
