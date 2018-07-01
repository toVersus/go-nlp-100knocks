package main

import (
	"bufio"
	"bytes"
	"flag"
	"fmt"
	"io"
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
	flag.IntVar(&columnNum, "number", 1, "specify a column to be sorted")
	flag.IntVar(&columnNum, "n", 1, "specify a column to be sorted")
	flag.Parse()

	f, err := os.Open(filePath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "could not open a file: %s\n  %s", filePath, err)
		os.Exit(1)
	}
	defer f.Close()

	if columnNum <= 0 {
		fmt.Fprintf(os.Stderr, "please specify a natural number, got=%d\n", columnNum)
		os.Exit(1)
	}

	if err := sortByVotes(f, columnNum); err != nil {
		fmt.Printf("could ntable sort lines by counts of overlap elements: %s\n  %s", filePath, err)
	}
}

// highTemp represents the details on high temperature in Japan
type highTemp struct {
	pref  string
	city  string
	value float64
	date  string
	vote  int
}

type highTempMap []*highTemp

func newHighTempMap(r io.Reader) highTempMap {
	var tempMap highTempMap
	sc := bufio.NewScanner(r)
	for sc.Scan() {
		fields := strings.Fields(strings.Replace(sc.Text(), "\t", " ", -1))
		temp, _ := strconv.ParseFloat(fields[2], 64)
		tempMap = append(tempMap, &highTemp{
			pref: fields[0], city: fields[1], value: temp, date: fields[3],
		})
	}
	return tempMap
}

// String concatenates and outputs each element of high temperature map
func (ht *highTemp) String() string {
	return fmt.Sprintf("%s\t%s\t%g\t%s", ht.pref, ht.city, ht.value, ht.date)
}

func (htm highTempMap) String() string {
	var out bytes.Buffer
	for i, t := range htm {
		out.WriteString(t.String())
		if i == len(htm)-1 {
			break
		}
		out.WriteString("\n")
	}

	return out.String()
}

// countVotes sorts slice of high temperature map and then counts the overlapped map
func (htm highTempMap) countVotes(columnNum int) highTempMap {
	comparators := []func(int, int) bool{
		func(i, j int) bool { return htm[i].pref > htm[j].pref },
		func(i, j int) bool { return htm[i].city > htm[j].city },
		func(i, j int) bool { return htm[i].value > htm[j].value },
		func(i, j int) bool { return htm[i].date > htm[j].date },
	}
	sort.SliceStable(htm, comparators[columnNum-1])

	matchers := []func(int, int) bool{
		func(i, j int) bool { return htm[i].pref == htm[j].pref },
		func(i, j int) bool { return htm[i].city == htm[j].city },
		func(i, j int) bool { return htm[i].value == htm[j].value },
		func(i, j int) bool { return htm[i].date == htm[j].date },
	}
	for i := 0; i < len(htm); i++ {
		for j := 0; j < i; j++ {
			if matchers[columnNum-1](i, j) {
				htm[i].vote++
				htm[j].vote++
			}
		}
	}
	return htm
}

// sortByVotes sorts the record by overlapped counts
func sortByVotes(r io.Reader, columnNum int) highTempMap {
	table := newHighTempMap(r).countVotes(columnNum)

	sort.SliceStable(table, func(i, j int) bool {
		return table[i].vote > table[j].vote
	})

	return table
}
