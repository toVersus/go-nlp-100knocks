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
	flag.IntVar(&columnNum, "number", 1, "specify number of row to be sorted")
	flag.IntVar(&columnNum, "n", 1, "specify number of row to be sorted")
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

	fmt.Println(sortByColumnNum(f, columnNum).String())
}

// highTemp represents the information about the highest temperature in Japan
type highTemp struct {
	pref  string
	city  string
	value float64
	date  string
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

// sortByColumnNum sorts the high temperature map in stable order as well as in ascending order.
func sortByColumnNum(r io.Reader, columnNum int) highTempMap {
	tempMap := newHighTempMap(r)

	comparators := []func(int, int) bool{
		func(i, j int) bool { return tempMap[i].pref > tempMap[j].pref },
		func(i, j int) bool { return tempMap[i].city > tempMap[j].city },
		func(i, j int) bool { return tempMap[i].value > tempMap[j].value },
		func(i, j int) bool { return tempMap[i].date > tempMap[j].date },
	}
	sort.SliceStable(tempMap, comparators[columnNum-1])

	return tempMap
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
