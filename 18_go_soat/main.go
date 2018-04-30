package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	var (
		filePath  string
		columnNum int
	)
	flag.StringVar(&filePath, "file", "", "specify a file path")
	flag.StringVar(&filePath, "f", "", "specify a file path")
	flag.IntVar(&columnNum, "number", 1, "specify number of row to be sorted")
	flag.IntVar(&columnNum, "n", 1, "specify number of row to be sorted")
	flag.Parse()

	if _, err := os.Stat(filePath); err != nil {
		fmt.Fprintf(os.Stderr, "could not find a file: %s\n  %s\n", filePath, err)
		os.Exit(1)
	}

	if columnNum <= 0 {
		fmt.Println("please specify a natural number")
		os.Exit(1)
	}

	if err := Sort(filePath, columnNum, os.Stdout); err != nil {
		fmt.Printf("could not sort lines of text in a file by specified row: %s\n", err)
	}
}

// Record represents scheme of data structure
type Record struct {
	Pref string
	City string
	Temp float64
	Date string
}

// Table represents compound record
type Table []Record

// NewTable returns a new table reading from specified file
func NewTable(fp *os.File) Table {
	var (
		fields []string
		table  Table
		temp   float64
	)
	sc := bufio.NewScanner(fp)
	for sc.Scan() {
		fields = strings.Fields(strings.Replace(sc.Text(), "\t", " ", -1))
		temp, _ = strconv.ParseFloat(fields[2], 64)
		table = append(table, Record{
			fields[0], fields[1], temp, fields[3],
		})
	}
	return table
}

// String concatenates and outputs each element of table
func (r Record) String() string {
	return fmt.Sprintf("%s\t%s\t%g\t%s", r.Pref, r.City, r.Temp, r.Date)
}

// Sort sorts the lines of a text in ascending order by specified column
func Sort(path string, columnNum int, file *os.File) error {
	f, err := os.Open(path)
	if err != nil {
		return err
	}
	defer f.Close()

	table := NewTable(f)

	comparators := []func(int, int) bool{
		func(i, j int) bool { return table[i].Pref > table[j].Pref },
		func(i, j int) bool { return table[i].City > table[j].City },
		func(i, j int) bool { return table[i].Temp > table[j].Temp },
		func(i, j int) bool { return table[i].Date > table[j].Date },
	}
	sort.SliceStable(table, comparators[columnNum-1])

	w := bufio.NewWriter(file)
	defer w.Flush()
	for i, t := range table {
		fmt.Fprint(w, t.String())
		if i == len(table)-1 {
			break
		}
		fmt.Fprintln(w, "")
	}

	return nil
}
