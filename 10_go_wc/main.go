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

	if _, err := os.Stat(filePath); err != nil {
		fmt.Fprintf(os.Stderr, "could not find a file:\n  %s\n", err)
		os.Exit(1)
	}

	count, err := countLineByScanner(filePath)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println(filePath, "has", count, "lines...")
}

// countLine counts lines of input file text.
func countLineByScanner(path string) (int, error) {
	f, err := os.Open(path)
	if err != nil {
		return -1, fmt.Errorf("could not open a file: %s\n  %s", path, err)
	}
	defer f.Close()

	sc := bufio.NewScanner(f)
	count := 0
	for sc.Scan() {
		count++
	}
	return count, nil
}

// countLine counts lines of input file text.
func countLineByReadLine(path string) (int, error) {
	f, err := os.Open(path)
	if err != nil {
		return -1, fmt.Errorf("could not open a file: %s\n  %s", path, err)
	}
	defer f.Close()

	r := bufio.NewReader(f)
	count := 0
	for {
		_, _, err := r.ReadLine()
		if (err != nil) && (err != io.EOF) {
			fmt.Println(err)
			os.Exit(1)
		}
		if err == io.EOF {
			break
		}
		count++
	}

	return count, nil
}
