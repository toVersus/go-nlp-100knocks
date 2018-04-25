package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"
)

func main() {
	var filePathFlag string
	flag.StringVar(&filePathFlag, "file", "", "specify a file path")
	flag.StringVar(&filePathFlag, "f", "", "specify a file path")
	flag.Parse()

	if _, err := os.Stat(filePathFlag); err != nil {
		fmt.Fprintf(os.Stderr, "could not find a file: %s\n  %#v", filePathFlag, err)
		os.Exit(1)
	}
	fmt.Println(strings.Join(convertTabToSpace(filePathFlag), "\n"))
}

// convertTabToSpace reads lines of file one by one and converts tab to whitespace.
// the same result obtained from UNIX command:
//   sed -e 's/\t/ /g'
//   tr '\t' ' '
//   expand
func convertTabToSpace(path string) []string {
	var content []string
	f, _ := os.Open(path)
	defer f.Close()

	sc := bufio.NewScanner(f)
	for sc.Scan() {
		content = append(content, strings.Replace(sc.Text(), "\t", " ", -1))
	}

	return content
}
