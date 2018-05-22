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

	morphs := newMorpheme(f)
	fmt.Printf("%#v\n", morphs[3])
}

// Morph represents the scheme of Cabocha format.
type Morph struct {
	surface string
	base    string
	pos     string
	pos1    string
}

// Chunk represents the list of Morphs and number of phrases and its depended phrases
type Chunk struct {
	morphems []Morph
	dst      int
	srcs     []int
}

// Passage represents the bunch of phrases
type Passage []Chunk

// newMorpheme implements the constructor of Morpheme.
// Cabocha outputs the following data structure through morphological analysis.
// <Surface>\t<POS>,<POS subtyping1>,<POS subtyping2>,<POS subtyping3>,<Conjugation Form>,<Conjugation>,<Base>,<Furigana>,<Pronunciation>
func newMorpheme(r io.Reader) Passage {
	// Preprocess the line by separating with the tab.
	chunk, passage := Chunk{}, Passage{}
	sc := bufio.NewScanner(r)
	for sc.Scan() {
		if sc.Text() == "" {
			continue
		}

		if (sc.Text() == "EOS") && (len(chunk.morphems) != 0) {
			passage = append(passage, chunk)
			chunk = Chunk{}
		}

		var fields []string
		tmp := strings.Split(sc.Text(), "\t")
		if len(tmp) < 2 {
			continue
		}

		fields = append(tmp[:1], strings.Split(tmp[1], ",")...)
		chunk.morphems = append(chunk.morphems, Morph{
			surface: fields[0],
			base:    fields[7],
			pos:     fields[1],
			pos1:    fields[2],
		})
	}
	return passage
}
