package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
)

func main() {
	var filePath string
	flag.StringVar(&filePath, "file", "", "specify a file path")
	flag.StringVar(&filePath, "f", "", "specify a file path")
	flag.Parse()

	f, err := os.Open(filePath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "could not open a file: %s\n  %s", filePath, err)
		os.Exit(1)
	}
	defer f.Close()

	count, err := countLineByScanner(f)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println(filePath, "has", count, "lines...")
}

// countLine counts lines of input file text.
func countLineByScanner(r io.Reader) (int, error) {
	sc := bufio.NewScanner(r)
	count := 0
	for sc.Scan() {
		count++
	}
	return count, nil
}

// countLine counts lines of input file text.
func countLineByReadLine(r io.Reader) (int, error) {
	reader := bufio.NewReader(r)
	count := 0
	for {
		_, _, err := reader.ReadLine()
		if (err != nil) && (err != io.EOF) {
			return 0, fmt.Errorf("could not read a line: %s", err)
		}
		if err == io.EOF {
			break
		}
		count++
	}

	return count, nil
}
