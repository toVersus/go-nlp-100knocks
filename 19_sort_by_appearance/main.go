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
	flag.IntVar(&columnNum, "number", 1, "specify a column to be sorted")
	flag.IntVar(&columnNum, "n", 1, "specify a column to be sorted")
	flag.Parse()

	if err := sortByVotes(filePath, columnNum, *os.Stdout); err != nil {
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

// countVotes sorts slice of high temperature map and then counts the overlapped map
func countVotes(columnNum int, tempMap highTempMap) highTempMap {
	comparators := []func(int, int) bool{
		func(i, j int) bool { return tempMap[i].pref > tempMap[j].pref },
		func(i, j int) bool { return tempMap[i].city > tempMap[j].city },
		func(i, j int) bool { return tempMap[i].value > tempMap[j].value },
		func(i, j int) bool { return tempMap[i].date > tempMap[j].date },
	}
	sort.SliceStable(tempMap, comparators[columnNum-1])

	matchers := []func(int, int) bool{
		func(i, j int) bool { return tempMap[i].pref == tempMap[j].pref },
		func(i, j int) bool { return tempMap[i].city == tempMap[j].city },
		func(i, j int) bool { return tempMap[i].value == tempMap[j].value },
		func(i, j int) bool { return tempMap[i].date == tempMap[j].date },
	}
	for i := 0; i < len(tempMap); i++ {
		for j := 0; j < i; j++ {
			if matchers[columnNum-1](i, j) {
				tempMap[i].vote++
				tempMap[j].vote++
			}
		}
	}
	return tempMap
}

// sortByVotes sorts the record by overlapped counts
func sortByVotes(path string, columnNum int, file os.File) error {
	f, err := os.Open(path)
	if err != nil {
		return err
	}
	defer f.Close()

	table := countVotes(columnNum, newHighTempMap(f))

	sort.SliceStable(table, func(i, j int) bool {
		return table[i].vote > table[j].vote
	})

	w := bufio.NewWriter(&file)
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
