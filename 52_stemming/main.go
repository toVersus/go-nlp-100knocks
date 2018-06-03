package main

import (
	"bufio"
	"bytes"
	"flag"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/reiver/go-porterstemmer"
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

	text := stemming(f)

	dest, err := os.Create(destFilePath)
	if err != nil {
		fmt.Printf("could not create a file: %s\n  %s\n", destFilePath, err)
		os.Exit(1)
	}
	defer dest.Close()
	dest.WriteString(text)
}

// stemming reduces inflected word to its root form and append it to the original word form.
func stemming(r io.Reader) string {
	var buf bytes.Buffer
	sc := bufio.NewScanner(r)
	for sc.Scan() {
		if sc.Text() == "" {
			buf.WriteString("\n")
			continue
		}
		stem := porterstemmer.StemString(sc.Text())
		buf.WriteString(sc.Text() + "\t" + stem + "\n")
	}
	return strings.TrimRight(buf.String(), "\n")
}
