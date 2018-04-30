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

// highTemp represents the details on high temperature in Japan
type highTemp struct {
	pref  string
	city  string
	value float64
	date  string
}

type highTempMap []highTemp

func newHighTempMap(fp *os.File) highTempMap {
	var (
		fields  []string
		tempMap highTempMap
		temp    float64
	)
	sc := bufio.NewScanner(fp)
	for sc.Scan() {
		fields = strings.Fields(strings.Replace(sc.Text(), "\t", " ", -1))
		temp, _ = strconv.ParseFloat(fields[2], 64)
		tempMap = append(tempMap, highTemp{
			pref: fields[0], city: fields[1], value: temp, date: fields[3],
		})
	}
	return tempMap
}

// String concatenates and outputs each element of high temperature map
func (ht *highTemp) String() string {
	return fmt.Sprintf("%s\t%s\t%g\t%s", ht.pref, ht.city, ht.value, ht.date)
}

// Sort sorts the high temperature map in stable order as well as in ascending order.
func Sort(path string, columnNum int, file *os.File) error {
	f, err := os.Open(path)
	if err != nil {
		return err
	}
	defer f.Close()

	tempMap := newHighTempMap(f)

	comparators := []func(int, int) bool{
		func(i, j int) bool { return tempMap[i].pref > tempMap[j].pref },
		func(i, j int) bool { return tempMap[i].city > tempMap[j].city },
		func(i, j int) bool { return tempMap[i].value > tempMap[j].value },
		func(i, j int) bool { return tempMap[i].date > tempMap[j].date },
	}
	sort.SliceStable(tempMap, comparators[columnNum-1])

	w := bufio.NewWriter(file)
	defer w.Flush()
	for i, t := range tempMap {
		fmt.Fprint(w, t.String())
		if i == len(tempMap)-1 {
			break
		}
		fmt.Fprintln(w, "")
	}

	return nil
}
