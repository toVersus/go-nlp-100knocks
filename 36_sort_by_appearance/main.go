package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	"sort"
	"strings"
)

func main() {
	var filePath string
	flag.StringVar(&filePath, "file", "", "specify a file path")
	flag.StringVar(&filePath, "f", "", "specify a file path")
	flag.Parse()

	f, err := os.Open(filePath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "cannot find the specified file: %s\n  %s\n", filePath, err)
	}
	defer f.Close()

	morphs, err := newMorpheme(f)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	ranking := morphs.sortByAppearance()
	fmt.Printf("Surface: %s, Counts: %d\n", ranking[1].Key, ranking[1].Count)
}

// morpheme represents the mapping list of MeCab format.
type morpheme map[string]string

// morphemes represents lists of Morpheme.
type morphemes []*morpheme

// newMorpheme implements the constructor of Morpheme.
// MeCab outputs the following data structure through morphological analysis.
// <Surface>\t<POS>,<POS subtyping1>,<POS subtyping2>,<POS subtyping3>,<Conjugation Form>,<Conjugation>,<Base>,<Furigana>,<Pronunciation>
func newMorpheme(r io.Reader) (morphemes, error) {
	morphs := morphemes{}
	sc := bufio.NewScanner(r)
	for sc.Scan() {
		if sc.Text() == "EOS" || sc.Text() == "" {
			continue
		}
		// separation of the surface field and the rest of the fields.
		surf := strings.Split(sc.Text(), "\t")
		// separation of the rest of the fields.
		other := strings.Split(surf[1], ",")
		morphs = append(morphs, &morpheme{
			"surface": surf[0],
			"base":    other[6],
			"pos":     other[0],
			"pos1":    other[1],
		})
	}
	return morphs, nil
}

// CountSorter is used for sorting the key by counts
type CountSorter struct {
	Key   string
	Count int
}

// CountSorters represents the slice of Countsorter
type CountSorters []CountSorter

// sortByAppearance returns the surface sorted by its appearance in Morphemes
func (morphes *morphemes) sortByAppearance() CountSorters {
	// counter counts frequency of word
	counter := map[string]int{}
	var sortedWords CountSorters
	var word string

	// Use separate slice container to memorize stable iteration order
	// See Iteration order section in the following URL
	// https://blog.golang.org/go-maps-in-action
	var keys []string

	for _, morph := range *morphes {
		word = (*morph)["surface"]
		if _, ok := counter[word]; ok {
			counter[word]++
		} else {
			// Memorize original order of words
			keys = append(keys, word)
			counter[word] = 1
		}
	}

	for _, key := range keys {
		// Assign sorter in fixed order using memorized keys
		sortedWords = append(sortedWords, CountSorter{
			Key:   key,
			Count: counter[key],
		})
	}

	sort.SliceStable(sortedWords, func(i, j int) bool {
		return sortedWords[i].Count > sortedWords[j].Count
	})

	return sortedWords
}
