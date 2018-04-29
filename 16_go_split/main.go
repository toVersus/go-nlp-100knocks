package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

func main() {
	var (
		filePath string
		fileNum  int
	)
	flag.StringVar(&filePath, "file", "", "specify a file path")
	flag.StringVar(&filePath, "f", "", "specify a file path")
	flag.IntVar(&fileNum, "n", 1, "specify number of output files")
	flag.IntVar(&fileNum, "number", 1, "specify number of output files")
	flag.Parse()

	if _, err := os.Stat(filePath); err != nil {
		fmt.Printf("could not find the specified file: %s\n  %s\n", filePath, err)
		os.Exit(1)
	}

	if fileNum < 1 {
		fmt.Println("please specify a positive number")
		os.Exit(1)
	}

	if err := split(filePath, fileNum); err != nil {
		fmt.Printf("%s\n", err)
		os.Exit(1)
	}
}

// split splits a file into pieces and outputs specified number of files named <INPUT FILE NAME>_1.<INPUT FILE EXT>...
func split(path string, fileNum int) error {
	f, err := os.Open(path)
	if err != nil {
		return fmt.Errorf("could not open the file: %s\n  %s", path, err)
	}
	defer f.Close()

	sc := bufio.NewScanner(f)
	var out []string
	count := 0
	for sc.Scan() {
		out = append(out, sc.Text())
		count++
	}

	// Create output files
	ws := make([]*bufio.Writer, fileNum)
	fds := make([]*os.File, fileNum)
	ext := filepath.Ext(path)
	for i := 0; i < fileNum; i++ {
		fds[i], _ = os.Create(strings.TrimSuffix(path, ext) + "_" + strconv.Itoa(i+1) + ext)
		ws[i] = bufio.NewWriter(fds[i])
		defer fds[i].Close()
	}

	// Relational expression between count (lines of text) and fileNum (number of output file)
	x := count / fileNum
	y := count % fileNum
	// count = (x + 1) * y + x * (fileNum - y)
	// 0 <= i < y     -> [i*(x+1):(i+1)*(x+1)]
	// i == y         -> [(y+1)*(x+1):(y+1)*(x+1)+x]
	// y < i < fileNum  -> [(y+1)*(x+1)+x:(y+1)*(x+1)+x*(i-y)]
	// TODO: You’re overthinking it
	var line string
	var tmp int
	for i := 0; i < fileNum; i++ {
		if i < y {
			line = strings.TrimRight(strings.Join(out[i*(x+1):(i+1)*(x+1)], "\n"), "\n")
			tmp = (i + 1) * (x + 1)
		} else {
			line = strings.TrimRight(strings.Join(out[tmp:tmp+x], "\n"), "\n")
			tmp = tmp + x
		}
		if _, err := ws[i].WriteString(line); err != nil {
			return err
		}
		ws[i].Flush()
	}

	return nil
}
