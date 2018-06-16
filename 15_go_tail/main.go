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

	f, err := os.Open(filePath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "could not open a file: %s\n  %s", filePath, err)
		os.Exit(1)
	}
	defer f.Close()

	fmt.Println(tail(f, lineNum))
}

// tail prints the last n-th lines of input file to the sprcified file descripter
func tail(r io.Reader, lineNum int) (string, error) {
	var (
		count        int
		currentIndex int
		isPrefix     bool
	)
	reader := bufio.NewReaderSize(r, 4096)
	out := make([]string, lineNum)
	for {
		err := func() error {
			var (
				readLineErr error
				tmp, line   []byte
			)
			isPrefix = true
			for isPrefix && readLineErr == nil {
				tmp, isPrefix, readLineErr = reader.ReadLine()
				line = append(line, tmp...)
			}
			if readLineErr != nil {
				return readLineErr
			}

			// Overwrite the elements of array over and over
			currentIndex = count % lineNum
			out[currentIndex] = fmt.Sprintln(string(line))

			count++
			return nil
		}()
		if err == io.EOF {
			break
		} else if err != nil {
			return "", err
		}
	}

	return strings.Join(out[currentIndex+1:], "") + strings.Join(out[:currentIndex+1], ""), nil
}
