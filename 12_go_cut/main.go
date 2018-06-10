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
		srcPath, destPath string
		columnNum         int
	)
	flag.StringVar(&srcPath, "src", "", "specify source file path")
	flag.StringVar(&srcPath, "s", "", "specify source file path")
	flag.StringVar(&destPath, "dest", "./col.txt", "specify output file path")
	flag.StringVar(&destPath, "d", "./col.txt", "specify output file path")
	flag.IntVar(&columnNum, "n", 1, "specify a n-th column to be extracted")

	flag.Parse()

	f, err := os.Open(srcPath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "could not open a file: %s\n  %s", srcPath, err)
		os.Exit(1)
	}
	defer f.Close()

	text := cut(f, columnNum)
	if output(destPath, text); err != nil {
		fmt.Fprintf(os.Stderr, "could not create a file: %s\n  %s", srcPath, err)
		os.Exit(1)
	}
}

// cut extracts the portion of text from a file by selecting rows and write down into a new file.
// refer to the following UNIX command:
//   cat foo.txt | sed 's/[\t ]\+/\t/g' | cut -f3
func cut(r io.Reader, columnNum int) string {
	var buf bytes.Buffer
	sc := bufio.NewScanner(r)
	for sc.Scan() {
		ss := strings.Fields(strings.Replace(sc.Text(), "\t", " ", -1))
		if len(ss) < columnNum {
			continue
		}
		buf.WriteString(ss[columnNum-1] + "\n")
	}
	return strings.TrimRight(buf.String(), "\n")
}

// output just creates a file with given contents.
func output(filepath, content string) error {
	f, err := os.Create(filepath)
	if err != nil {
		return fmt.Errorf("could not create a file: %s", err)
	}
	defer f.Close()
	f.WriteString(content)

	return nil
}
