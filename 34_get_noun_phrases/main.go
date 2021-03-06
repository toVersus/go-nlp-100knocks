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
	fmt.Println(morphs.getNounPhrases())
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

// getNounPhrases returns the list of noun phrases such as B of A ("AのB" in Japanese).
func (morphes morphemes) getNounPhrases() string {
	nounPhrases := []string{}
	for midIndex, morphe := range morphes {
		if ((*morphe)["surface"] == "の") && (midIndex != 0) {
			prevMorpheme := *morphes[midIndex-1]
			nextMorpheme := *morphes[midIndex+1]
			if (prevMorpheme["pos"] == "名詞") && (nextMorpheme["pos"] == "名詞") {
				nounPhrases = append(nounPhrases, prevMorpheme["surface"]+(*morphe)["surface"]+nextMorpheme["surface"])
			}
		}
	}
	return strings.Join(nounPhrases, "\n")
}
