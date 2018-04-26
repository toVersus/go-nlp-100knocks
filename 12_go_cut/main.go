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
		srcPath, destPath string
		columnNum         int
	)
	flag.StringVar(&srcPath, "src", "", "specify source file path")
	flag.StringVar(&srcPath, "s", "", "specify source file path")
	flag.StringVar(&destPath, "dest", "./col.txt", "specify output file path")
	flag.StringVar(&destPath, "d", "./col.txt", "specify output file path")
	flag.IntVar(&columnNum, "n", 1, "specify a n-th column to be extracted")

	flag.Parse()

	if _, err := os.Stat(srcPath); err != nil {
		fmt.Fprintf(os.Stderr, "could not find a file: %s\n  %s\n", srcPath, err)
		os.Exit(1)
	}

	cut(srcPath, destPath, columnNum)
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
	defer w.Flush()

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
}
