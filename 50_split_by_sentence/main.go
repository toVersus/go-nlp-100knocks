package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	"regexp"
)

const (
	upperCaseForLastSent = "A"
)

var sentReg = regexp.MustCompile(`.+?[.;?!]\s[A-Z]`)

func main() {
	var filePath string
	var destFilePath string
	flag.StringVar(&filePath, "file", "", "specify a file path")
	flag.StringVar(&filePath, "f", "", "specify a file path")
	flag.StringVar(&destFilePath, "dest", "", "specify a dest file path")
	flag.StringVar(&destFilePath, "d", "", "specify a dest file path")
	flag.Parse()

	f, err := os.Open(filePath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "cannot find the specified file: %s\n  %s\n", filePath, err)
		os.Exit(1)
	}
	defer f.Close()

	sentText, err := splitBySent(f)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println(sentText)
}

// splitBySent splits the text into sentences following the below pattern matching.
// (. or ; or : or ? or !) → whitespace → upper case
func splitBySent(r io.Reader) (string, error) {
	var sents string
	sc := bufio.NewScanner(r)
	for sc.Scan() {
		var lastChar string
		// add dummy padding at the end of the line
		text := sc.Text() + " " + upperCaseForLastSent

		for _, sent := range sentReg.FindAllString(text, -1) {
			if sent == "" {
				continue
			}
			if len(lastChar) != 0 {
				// add head character at the beggining of each sentence
				sent = lastChar + sent
			}
			// remove whitespace and head character from the end of the sentence
			sents += sent[:len(sent)-2] + "\n"
			lastChar = sent[len(sent)-1:]
		}
	}
	return sents, nil
}
