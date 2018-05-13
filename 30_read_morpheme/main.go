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

	ptr, err := newMorpheme(filePath)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	fmt.Printf("%#v\n", ptr)
}

// morpheme represents the mapping list of MeCab format.
type morpheme map[string]string

// morphemes represents lists of Morpheme.
type morphemes []morpheme

// newMorpheme implements the constructor of Morpheme.
// MeCab outputs the following data structure through morphological analysis.
// <Surface>\t<POS>,<POS subtyping1>,<POS subtyping2>,<POS subtyping3>,<Conjugation Form>,<Conjugation>,<Base>,<Furigana>,<Pronunciation>
func newMorpheme(path string) (morphemes, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("cannot open the specified file: %s\n  %s", path, err)
	}
	defer f.Close()

	morphs := morphemes{}
	sc := bufio.NewScanner(f)
	for sc.Scan() {
		if sc.Text() == "EOS" || sc.Text() == "" {
			continue
		}
		// separation of the surface field and the rest of the fields.
		surf := strings.Split(sc.Text(), "\t")
		// separation of the rest of the fields.
		other := strings.Split(surf[1], ",")
		morphs = append(morphs, morpheme{
			"surface": surf[0],
			"base":    other[6],
			"pos":     other[0],
			"pos1":    other[1],
		})
	}
	return morphs, nil
}
