package main

import (
	"bufio"
	"bytes"
	"flag"
	"fmt"
	"io"
	"os"
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
	fmt.Println(morphs.filterByPos("動詞").stringify("base"))
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

// filterByPos returns the morphemes filtered by specified keyword
func (morphs morphemes) filterByPos(keyword string) morphemes {
	var filtered morphemes
	for _, m := range morphs {
		if (*m)["pos"] == keyword {
			filtered = append(filtered, m)
		}
	}
	return filtered
}

func (morphs morphemes) stringify(field string) string {
	buf := bytes.Buffer{}
	for _, morph := range morphs {
		if morph == nil {
			continue
		}
		buf.WriteString((*morph)[field] + "\n")
	}
	return strings.TrimRight(buf.String(), "\n")
}
