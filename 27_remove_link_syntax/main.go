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

var templateFieldReg = regexp.MustCompile(`(?:|\s)([^|]+)(?:\s=\s)(.+)(?:\n?|)`)
var markupSyntaxReg = regexp.MustCompile(`'{2,5}`)
var internalLinkReg = regexp.MustCompile(`\[\[(?:[^|]+\|)?([^|]*)\]\]`)

// getTemplate returns key value pair of template by removing emphatic style and internal link.
func (articles Articles) getTemplate() map[string]string {
	dict := make(map[string]string)
	for _, a := range articles {
		for _, line := range templateFieldReg.FindAllStringSubmatch(a.Text, -1) {
			// Remove emphatic style (Italics, bold, and both) symbols.
			// https://en.wikipedia.org/wiki/Help:Cheatsheet
			line[2] = markupSyntaxReg.ReplaceAllString(line[2], "")

			// Remove internal link (link to a section) synbols.
			// https://en.wikipedia.org/wiki/Help:Cheatsheet
			// TODO: Do not remove other parts except internal link
			if internalLinkReg.MatchString(line[2]) {
				for _, word := range internalLinkReg.FindAllStringSubmatch(line[2], -1) {
					word[1] = strings.Replace(word[1], "]]", "", -1)
					line[2] = strings.Replace(word[1], "[[", "", -1)
				}
			}
			dict[line[1]] = line[2]
		}
	}
	return dict
}
