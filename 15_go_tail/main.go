package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	var (
		filePath string
		lineNum  int
	)

	flag.StringVar(&filePath, "file", "", "specify a file path")
	flag.StringVar(&filePath, "f", "", "specify a file path")
	flag.IntVar(&lineNum, "lines", 1, "specify a number of line from the tail")
	flag.IntVar(&lineNum, "n", 1, "specify a number of lines from the tail")
	flag.Parse()

	if _, err := os.Stat(filePath); err != nil {
		fmt.Fprintf(os.Stderr, "could not find a file: %s\n  %s\n", filePath, err)
		os.Exit(1)
	}

	tail(filePath, lineNum, *os.Stdout)
}

// tail prints the last n-th lines of input file to the sprcified file descripter
func tail(path string, lineNum int, file os.File) error {
	f, err := os.Open(path)
	if err != nil {
		return err
	}
	defer f.Close()

	var (
		count        int
		currentIndex int
		isPrefix     bool
	)
	w := bufio.NewWriter(&file)
	r := bufio.NewReaderSize(f, 4096)
	out := make([]string, lineNum)
	for {
		err := func() error {
			var tmp, line []byte
			isPrefix = true
			for isPrefix && err == nil {
				tmp, isPrefix, err = r.ReadLine()
				line = append(line, tmp...)
			}
			if err != nil {
				return err
			}

			// Overwrite the elements of array over and over
			currentIndex = count % lineNum
			out[currentIndex] = fmt.Sprintf("%s\n", string(line))

			count++
			return nil
		}()
		if err == io.EOF {
			break
		} else if err != nil {
			return err
		}
	}

	// Put the elements of array in order
	fmt.Fprintf(w, "%s", strings.Join(out[currentIndex+1:], ""))
	fmt.Fprintf(w, "%s", strings.Join(out[:currentIndex+1], ""))

	w.Flush()

	return nil
}
