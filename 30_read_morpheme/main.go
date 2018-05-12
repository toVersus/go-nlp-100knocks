package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"
)

func main() {
	var filePath string
	flag.StringVar(&filePath, "file", "", "specify a file path")
	flag.StringVar(&filePath, "f", "", "specify a file path")
	flag.Parse()

	if _, err := os.Stat(filePath); err != nil {
		fmt.Fprintf(os.Stderr, "cannot find the specified file: %s\n  %s\n", filePath, err)
	}

	ptr, err := parseMorphemes(filePath)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	fmt.Printf("%#v\n", ptr)
}

// morpheme represents the mapping list of MeCab format.
type morpheme map[string]string

// newMorpheme implements the constructor of Morpheme.
// MeCab outputs the following data structure through morphological analysis.
// <Surface>\t<POS>,<POS subtyping1>,<POS subtyping2>,<POS subtyping3>,<Conjugation Form>,<Conjugation>,<Base>,<Furigana>,<Pronunciation>
func newMorpheme(rawLine string) *morpheme {
	// Preprocess the line by separating with the tab.
	tmpFields := strings.Split(rawLine, "\t")
	fields := append(tmpFields[:1], strings.Split(tmpFields[1], ",")...)
	return &morpheme{
		"surface": fields[0],
		"base":    fields[7],
		"pos":     fields[1],
		"pos1":    fields[2],
	}
}

// morphemes represents lists of Morpheme.
type morphemes []morpheme

// parseMorphemes parses Morpheme information from a file processed by Mecab.
func parseMorphemes(path string) (*morphemes, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("cannot open the specified file: %s\n  %s", path, err)
	}
	defer f.Close()

	morphs := morphemes{}
	sc := bufio.NewScanner(f)
	for sc.Scan() {
		if sc.Text() != "EOS" && sc.Text() != "" {
			morphs = append(morphs, *newMorpheme(sc.Text()))
		}
	}
	return &morphs, nil
}
