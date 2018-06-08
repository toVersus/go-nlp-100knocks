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
	var filePath string
	flag.StringVar(&filePath, "file", "", "specify a file path")
	flag.StringVar(&filePath, "f", "", "specify a file path")
	flag.Parse()

	f, err := os.Open(filePath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "could not open a file: %s\n  %s", filePath, err)
	}
	defer f.Close()

	fmt.Println(convertTabToSpace(f))
}

// convertTabToSpace converts tab in the lines to whitespace one by one.
// The following UNIX command returns the same output:
//   sed -e 's/\t/ /g'
//   tr '\t' ' '
//   expand
func convertTabToSpace(r io.Reader) string {
	var buf bytes.Buffer
	sc := bufio.NewScanner(r)
	for sc.Scan() {
		buf.WriteString(strings.Replace(sc.Text(), "\t", " ", -1) + "\n")
	}

	return strings.TrimRight(buf.String(), "\n")
}
