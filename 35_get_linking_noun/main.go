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
	fmt.Println(morphs.getLinkingNounPhrases())
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

// getLinkingNounPhrases returns the linking noun phrases.
func (morphes morphemes) getLinkingNounPhrases() string {
	var (
		nounPieces         []string
		linkingNoun        string
		linkingnounPhrases []string
	)
	for currentIndex, morphe := range morphes {
		if (*morphe)["pos"] == "名詞" {
			nounPieces = append(nounPieces, (*morphe)["surface"])

			// In case that the sentence is ended with noun phrases
			if currentIndex == len(morphes)-1 {
				linkingNoun = strings.Join(nounPieces, "")
				linkingnounPhrases = append(linkingnounPhrases, linkingNoun)
			}
		} else {
			// Only concatenate noun pieces in case of linking noun
			if len(nounPieces) > 1 {
				linkingNoun = strings.Join(nounPieces, "")
				linkingnounPhrases = append(linkingnounPhrases, linkingNoun)
			}
			nounPieces = nil
		}
	}
	return strings.Join(linkingnounPhrases, "\n")
}
