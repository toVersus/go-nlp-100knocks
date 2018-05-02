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

// NewArticle is a constructer of Article.
func newArticle(a Article) Article {
	return Article{a.Title, a.Text}
}

// Articles represents the slice of Article.
type Articles []Article

func main() {
	var (
		filePath string
		keyword  string
	)
	flag.StringVar(&filePath, "file", "", "specify a file path")
	flag.StringVar(&filePath, "f", "", "specify a file path")
	flag.StringVar(&keyword, "key", "", "specify a keyword for filtering a wiki page")
	flag.Parse()

	results := readJSON(filePath).find(keyword)
	for _, result := range results {
		fmt.Println(result.Text)
	}
}

// readJSON reads and parse the json file and returns Articles
func readJSON(path string) Articles {
	var (
		buf      []byte
		article  Article
		articles Articles
		readErr  error
		err      error
	)

	file, err := os.Open(path)
	if err != nil {
		fmt.Printf("could not open a file specified: %s\n  %s\n", path, err)
	}
	defer file.Close()

	reader := bufio.NewReader(file)
	for {
		buf, readErr = reader.ReadBytes('\n')
		if (readErr != nil) && (readErr != io.EOF) {
			panic(err)
		}

		if err = json.Unmarshal(buf, &article); err != nil {
			fmt.Print("could not parse json file.")
			break
		}

		articles = append(articles, article)

		if readErr == io.EOF {
			break
		}
	}
	return articles
}

// Find returns filetered Articles by input string
func (articles Articles) find(keyword string) Articles {
	var filtered Articles
	for _, a := range articles {
		if strings.Contains(a.Title, keyword) {
			filtered = append(filtered, a)
		}
	}
	return filtered
}
