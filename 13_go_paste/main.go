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
		filePath1 string
		filePath2 string
		destPath  string
	)

	flag.StringVar(&filePath1, "file1", "", "specify a file1 path")
	flag.StringVar(&filePath1, "f1", "", "specify a file1 path")
	flag.StringVar(&filePath2, "file2", "", "specify a file2 path")
	flag.StringVar(&filePath2, "f2", "", "specify a file2 path")
	flag.StringVar(&destPath, "dest", "", "specify a destination file path")
	flag.StringVar(&destPath, "d", "", "specify a destination file path")
	flag.Parse()

	f1, err := os.Open(filePath1)
	if err != nil {
		fmt.Fprintf(os.Stderr, "could not open a file: %s\n  %s", filePath1, err)
		os.Exit(1)
	}
	defer f1.Close()

	f2, err := os.Open(filePath2)
	if err != nil {
		fmt.Fprintf(os.Stderr, "could not open a file: %s\n  %s", filePath2, err)
		os.Exit(1)
	}
	defer f2.Close()

	text := pasteByChannel(f1, f2)
	if output(destPath, text); err != nil {
		fmt.Fprintf(os.Stderr, "could not create a file: %s\n  %s", destPath, err)
		os.Exit(1)
	}
}

// pasteByChannel returns the sequence of lines one after the other by using channel.
func pasteByChannel(reader1, reader2 io.Reader) string {
	rx1 := make(chan string)
	rx2 := make(chan string)
	read := func(r io.Reader, yield chan string) {
		sc := bufio.NewScanner(r)
		for sc.Scan() {
			yield <- sc.Text()
		}
		close(yield)
	}

	var buf bytes.Buffer
	go read(reader1, rx1)
	go read(reader2, rx2)
	for {
		str1, ok1 := <-rx1
		str2, ok2 := <-rx2
		if !ok1 || !ok2 {
			return strings.TrimRight(buf.String(), "\n")
		}
		buf.WriteString(str1 + "\t" + str2 + "\n")
	}
}

// paste reads the line of each file and returns new text by concatenating them with tab-delimited.
func paste(reader1, reader2 io.Reader) (string, error) {
	r1, r2 := bufio.NewReader(reader1), bufio.NewReader(reader2)
	var buf bytes.Buffer
	for {
		line1, _, err := r1.ReadLine()
		if (err != nil) && (err != io.EOF) {
			return "", err
		}
		if err == io.EOF {
			break
		}

		line2, _, err := r2.ReadLine()
		if (err != nil) && (err != io.EOF) {
			return "", err
		}
		if err == io.EOF {
			break
		}
		buf.WriteString(string(line1) + "\t" + string(line2) + "\n")
	}
	return strings.TrimRight(buf.String(), "\n"), nil
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
