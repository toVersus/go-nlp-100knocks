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
		filePath string
		rowNum   int
	)
	flag.StringVar(&filePath, "file", "", "specify a file path")
	flag.StringVar(&filePath, "f", "", "specify a file path")
	flag.IntVar(&rowNum, "number", 1, "specify n-th row to be extracted")
	flag.IntVar(&rowNum, "n", 1, "specify n-th row to be extracted")
	flag.Parse()

	if _, err := os.Stat(filePath); err != nil {
		fmt.Fprintf(os.Stderr, "could not find a file: %s\n  %s\n", filePath, err)
		os.Exit(1)
	}

	if rowNum < 1 {
		fmt.Println("please specify a positive number")
		os.Exit(1)
	}

	if err := uniq(filePath, rowNum, *os.Stdout); err != nil {
		fmt.Printf("could not filter repeated lines: %s\n", err)
	}

}

// Item represents Set-type container.
type Item map[string]struct{}

// Set implements fundamental set operators.
type Set interface {
	Add(s string) Set
	Contains(s string) bool
	Strings() []string
	Union(other Item) Set
}

// Add adds element into set-typed container.
func (item Item) Add(s string) Set {
	item[s] = struct{}{}
	return item
}

// Contains returns true in case that element of set-typed container matches input string.
func (item Item) Contains(s string) bool {
	if _, ok := item[s]; !ok {
		return false
	}
	return true
}

// Strings outputs set-typed container as string slice.
func (item Item) Strings() []string {
	ss := make([]string, 0, len(item))

	for elem := range item {
		ss = append(ss, fmt.Sprintf("%s", elem))
	}
	return ss
}

// Union returns the set of all elemments in the collection.
func (item Item) Union(other Item) Set {
	union := Item{}

	for i := range item {
		union.Add(i)
	}

	for j := range other {
		union.Add(j)
	}

	return union
}

// uniq filters out repeated lines of specified row in a file
func uniq(path string, rowNum int, file os.File) error {
	f, err := os.Open(path)
	if err != nil {
		return err
	}
	defer f.Close()

	items := Item{}
	sc := bufio.NewScanner(f)
	for sc.Scan() {
		ss := strings.Fields(strings.Replace(sc.Text(), "\t", " ", -1))
		if len(ss) < rowNum {
			continue
		}
		items.Add(ss[rowNum-1])
	}

	w := bufio.NewWriter(&file)
	for _, item := range items.Strings() {
		fmt.Fprintf(w, "%s\n", item)
	}
	w.Flush()

	return nil
}
