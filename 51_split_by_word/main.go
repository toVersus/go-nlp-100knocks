package main

import (
	"bufio"
	"bytes"
	"flag"
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	var filePath, destFilePath string
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

	wordText := splitByWord(f)

	dest, err := os.Create(destFilePath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "could not create a file: %s\n  %s\n", destFilePath, err)
		os.Exit(1)
	}
	defer dest.Close()

	dest.WriteString(wordText)
}

// splitByWord splits the text into words by detecting whitespace as a token of word separator.
func splitByWord(r io.Reader) string {
	var buf bytes.Buffer
	sc := bufio.NewScanner(r)
	for sc.Scan() {
		words := strings.Fields(sc.Text())
		buf.WriteString(strings.Join(words, "\n") + "\n\n")
	}
	return buf.String()
}
