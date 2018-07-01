package main

import (
	"bufio"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"os"
	"strings"
)

// Article represents the body of Wiki page and its Title.
type Article struct {
	Title string `json:"title"`
	Text  string `json:"text"`
}

// Articles represents the slice of Article.
type Articles []Article

func main() {
	var filePath, keyword string
	flag.StringVar(&filePath, "file", "", "specify a file path")
	flag.StringVar(&filePath, "f", "", "specify a file path")
	flag.StringVar(&keyword, "key", "", "specify a keyword for filtering a wiki page")
	flag.Parse()

	f, err := os.Open(filePath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "could not open a file: %s\n  %s", filePath, err)
		os.Exit(1)
	}
	defer f.Close()

	articles, err := readJSON(f)
	if err != nil {
		fmt.Fprintf(os.Stderr, "could not get the list of articles: %s\n", err)
		os.Exit(1)
	}

	for _, a := range articles.find(keyword) {
		fmt.Println(a.Text)
	}
}

// readJSON reads/parses the json file and initiates Articles instance
func readJSON(r io.Reader) (Articles, error) {
	reader := bufio.NewReader(r)
	article, articles := Article{}, Articles{}
	for {
		buf, readErr := reader.ReadBytes('\n')
		if (readErr != nil) && (readErr != io.EOF) {
			return nil, fmt.Errorf("could not read a file content: %s", readErr)
		}
		if err := json.Unmarshal(buf, &article); err != nil {
			return nil, fmt.Errorf("could not parse json file: %s", err)
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
