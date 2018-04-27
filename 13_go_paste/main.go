package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
)

func main() {
	var (
		filePath1    string
		filePath2    string
		destFilePath string
	)

	flag.StringVar(&filePath1, "file1", "", "specify a file1 path")
	flag.StringVar(&filePath1, "f1", "", "specify a file1 path")
	flag.StringVar(&filePath2, "file2", "", "specify a file2 path")
	flag.StringVar(&filePath2, "f2", "", "specify a file2 path")
	flag.StringVar(&destFilePath, "dest", "", "specify a destination file path")
	flag.StringVar(&destFilePath, "d", "", "specify a destination file path")
	flag.Parse()

	if _, err := os.Stat(filePath1); err != nil {
		fmt.Fprintf(os.Stderr, "could not find a file: %s\n  %s\n", filePath1, err)
		os.Exit(1)
	}
	if _, err := os.Stat(filePath2); err != nil {
		fmt.Fprintf(os.Stderr, "could not find a file: %s\n  %s\n", filePath2, err)
		os.Exit(1)
	}

	if err := pasteByChannel(filePath1, filePath2, destFilePath); err != nil {
		fmt.Fprint(os.Stderr, err)
		os.Exit(1)
	}
}

// pasteByChannel writes lines consisting of the sequentially corresponding lines from each file,
// separated by tab, to the specified destination file.
func pasteByChannel(file1 string, file2 string, dest string) error {
	rx1 := make(chan string)
	rx2 := make(chan string)
	destFile, err := os.Create(dest)
	if err != nil {
		return fmt.Errorf("cannot create a file: %s\n  %s", dest, err)
	}

	defer destFile.Close()
	w := bufio.NewWriter(destFile)
	read := func(path string, yield chan string) {
		f, _ := os.Open(path)
		sc := bufio.NewScanner(f)
		for sc.Scan() {
			yield <- sc.Text()
		}
		f.Close()
		close(yield)
	}

	go read(file1, rx1)
	go read(file2, rx2)
	for {
		str1, ok1 := <-rx1
		str2, ok2 := <-rx2
		if !ok1 || !ok2 {
			w.Flush()
			return nil
		}
		fmt.Fprint(w, str1, "\t", str2, "\n")
	}
}

// paste reads the line of each file and creates a new file by concatenating them with tab-delimited.
func paste(filePath1 string, filePath2 string, destPath string) error {
	f1, _ := os.Open(filePath1)
	defer f1.Close()
	f2, _ := os.Open(filePath2)
	defer f2.Close()

	destFile, err := os.Create(destPath)
	if err != nil {
		return fmt.Errorf("cannot create a file: %s\n  %s", destPath, err)
	}
	defer destFile.Close()

	r1, r2 := bufio.NewReader(f1), bufio.NewReader(f2)
	w := bufio.NewWriter(destFile)
	defer w.Flush()
	for {
		line1, _, err := r1.ReadLine()
		if (err != nil) && (err != io.EOF) {
			return err
		}
		if err == io.EOF {
			break
		}

		line2, _, err := r2.ReadLine()
		if (err != nil) && (err != io.EOF) {
			return err
		}
		if err == io.EOF {
			break
		}
		fmt.Fprint(w, string(line1), "\t", string(line2), "\n")
	}
	return nil
}
