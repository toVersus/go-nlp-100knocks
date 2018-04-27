package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
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

	if err := head(filePath, lineNum, *os.Stdout); err != nil {
		fmt.Printf("could not read the n lines of text: %s\n", err)
	}
}

// head reads the first n lines of text from input file
func head(path string, lineNum int, file os.File) error {
	f, err := os.Open(path)
	if err != nil {
		return err
	}
	defer f.Close()

	w := bufio.NewWriter(&file)
	defer w.Flush()
	sc := bufio.NewScanner(f)
	for sc.Scan() {
		if lineNum <= 0 {
			break
		}
		fmt.Fprintf(w, "%s\n", sc.Text())
		lineNum--
	}

	return nil
}
