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
	var (
		lineNum  int
		filePath string
	)

	flag.StringVar(&filePath, "file", "", "specify a file path")
	flag.StringVar(&filePath, "f", "", "specify a file path")
	flag.IntVar(&lineNum, "lines", 1, "specify the first numbers of line")
	flag.IntVar(&lineNum, "n", 1, "specify the first numbers of line")
	flag.Parse()

	f, err := os.Open(filePath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "could not open a file: %s\n  %s", filePath, err)
		os.Exit(1)
	}
	defer f.Close()

	fmt.Println(head(f, lineNum))
}

// head reads the first n lines of text from input file
func head(r io.Reader, lineNum int) string {
	var buf bytes.Buffer
	sc := bufio.NewScanner(r)
	for sc.Scan() {
		if lineNum <= 0 {
			break
		}
		buf.WriteString(sc.Text() + "\n")
		lineNum--
	}

	return strings.TrimRight(buf.String(), "\n")
}
