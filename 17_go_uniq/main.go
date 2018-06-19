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
		rowNum   int
	)
	flag.StringVar(&filePath, "file", "", "specify a file path")
	flag.StringVar(&filePath, "f", "", "specify a file path")
	flag.IntVar(&rowNum, "number", 1, "specify n-th row to be extracted")
	flag.IntVar(&rowNum, "n", 1, "specify n-th row to be extracted")
	flag.Parse()

	if rowNum < 1 {
		fmt.Fprint(os.Stderr, "please specify a positive number")
		os.Exit(1)
	}

	f, err := os.Open(filePath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "could not open a file: %s\n  %s", filePath, err)
		os.Exit(1)
	}
	defer f.Close()

	item, err := uniq(f, rowNum)
	if err != nil {
		fmt.Fprint(os.Stderr, "could not filter repeated lines: ", err)
		os.Exit(1)
	}
	fmt.Println(strings.Join(item.Strings(), "\n"))
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

	for key := range item {
		ss = append(ss, fmt.Sprintf("%s", key))
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
func uniq(r io.Reader, rowNum int) (Item, error) {
	item := Item{}
	sc := bufio.NewScanner(r)
	for sc.Scan() {
		ss := strings.Fields(strings.Replace(sc.Text(), "\t", " ", -1))
		if len(ss) < rowNum {
			continue
		}
		item.Add(ss[rowNum-1])
	}

	return item, nil
}
