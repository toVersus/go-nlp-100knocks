package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"
)

func main() {
	var (
		srcFilePath  string
		destFilePath string
		columnNum    int
	)
	flag.StringVar(&srcFilePath, "src", "", "specify source file path")
	flag.StringVar(&srcFilePath, "s", "", "specify source file path")
	flag.StringVar(&destFilePath, "dest", "./col.txt", "specify output file path")
	flag.StringVar(&destFilePath, "d", "./col.txt", "specify output file path")
	flag.IntVar(&columnNum, "n", 1, "specify a n-th column to be extracted")

	flag.Parse()

	if _, err := os.Stat(srcFilePath); err != nil {
		fmt.Fprintf(os.Stderr, "could not find a file: %s\n  %s\n", srcFilePath, err)
		os.Exit(1)
	}

	cut(srcFilePath, destFilePath, columnNum)
}

// cut extracts the portion of text from a file by selecting rows and write down into a new file.
// refer to the following UNIX command:
//   cat foo.txt | sed 's/[\t ]\+/\t/g' | cut -f3
func cut(srcPath string, destPath string, columnNum int) {
	src, _ := os.Open(srcPath)
	defer src.Close()

	dest, _ := os.Create(destPath)
	defer dest.Close()

	sc := bufio.NewScanner(src)
	w := bufio.NewWriter(dest)
	i := 0
	for sc.Scan() {
		ss := strings.Fields(strings.Replace(sc.Text(), "\t", " ", -1))
		if len(ss) < columnNum {
			fmt.Fprint(w, "")
		} else {
			fmt.Fprintf(w, "%s\n", ss[columnNum-1])
		}
		i++
	}
	w.Flush()
}
