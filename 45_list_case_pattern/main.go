package main

import (
	"bufio"
	"bytes"
	"flag"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func main() {
	var filePath, dstFilePath string
	flag.StringVar(&filePath, "file", "", "specify a file path")
	flag.StringVar(&filePath, "f", "", "specify a file path")
	flag.StringVar(&dstFilePath, "dest", "", "specify a output file path")
	flag.StringVar(&dstFilePath, "d", "", "specify a output file path")
	flag.Parse()

	f, err := os.Open(filePath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "cannot find the specified file: %s\n  %s\n", filePath, err)
		os.Exit(1)
	}
	defer f.Close()

	content := newChunkPassage(f).stringifyCasePattern()
	if err := output(dstFilePath, content); err != nil {
		fmt.Fprintf(os.Stderr, "cannot find the specified file: %s\n  %s\n", filePath, err)
		os.Exit(1)
	}
}

// Morph represents the scheme of Cabocha format.
type Morph struct {
	surface string
	base    string
	pos     string
	pos1    string
}

// newMorph implements the constructor of Morpheme.
// Cabocha outputs the following data structure through morphological analysis.
// <Surface>\t<POS>,<POS subtyping1>,<POS subtyping2>,<POS subtyping3>,<Conjugation Form>,<Conjugation>,<Base>,<Furigana>,<Pronunciation>
func newMorph(line string) *Morph {
	// Preprocess the line by separating with the tab.
	var fields []string
	tmp := strings.Split(line, "\t")
	if len(tmp) < 2 {
		return nil
	}
	fields = append(tmp[:1], strings.Split(tmp[1], ",")...)

	return &Morph{
		surface: fields[0],
		base:    fields[7],
		pos:     fields[1],
		pos1:    fields[2],
	}
}

// Chunk represents the list of Morphs and number of phrases and its depended phrases
type Chunk struct {
	morphems []*Morph
	dst      int
	srcs     []int
}

// Passage represents the bunch of phrases
type Passage []Chunk

// ChunkPassage represents the bunch of passages
type ChunkPassage []Passage

// newChunkPassage reads results of Morpheme analysis and returns the lines of Morphemes one by one.
func newChunkPassage(r io.Reader) *ChunkPassage {
	var (
		chunkPassage ChunkPassage
		dict         []int
		passage      Passage
		chunk        Chunk
	)

	sc := bufio.NewScanner(r)
	for sc.Scan() {
		// Catch the line of Morpheme analysis
		if morph := newMorph(sc.Text()); morph != nil {
			chunk.morphems = append(chunk.morphems, morph)
			continue
		}

		// Catch the line of Dependency analysis
		if strings.HasPrefix(sc.Text(), "*") {
			// Append the chunk assigned the result of morpheme analysis in previous phrase
			if chunk.morphems != nil {
				passage = append(passage, chunk)
			}

			// tmp = ["*", "1", "2D", "0/1", "2.397100"]
			tmp := strings.Fields(sc.Text())
			chunk.dst, _ = strconv.Atoi(strings.Split(tmp[2], "D")[0]) // "2D" -> 1
			dict = append(dict, chunk.dst)
			chunk.morphems = nil
			continue
		}

		// Catch the continuous "EOS" string case:
		if chunk.morphems == nil {
			continue
		}

		passage = append(passage, chunk)

		// Assign the src index number
		for index, srcIndex := range dict {
			if srcIndex == -1 {
				break
			}
			passage[srcIndex].srcs = append(passage[srcIndex].srcs, index)
		}

		chunkPassage = append(chunkPassage, passage)

		// Initialize the variables
		dict = dict[0:0]
		passage = nil
		chunk = Chunk{nil, 0, nil}
	}

	return &chunkPassage
}

// stringifyCasePattern returns formatted case pattern.
func (chunkPassage ChunkPassage) stringifyCasePattern() string {
	var verb, particle string
	var buf bytes.Buffer
	for _, passage := range chunkPassage {
		for i, chunk := range passage {
			if chunk.srcs == nil {
				continue
			}

			verb = getVerb(chunk.morphems)
			if verb == "" {
				continue
			}

			particle = getParticles(i, passage)
			if particle != "" {
				buf.WriteString(verb + "\t" + particle + "\n")
			}
		}
	}
	return strings.TrimRight(buf.String(), "\n")
}

func output(filepath, content string) error {
	f, err := os.Create(filepath)
	if err != nil {
		return fmt.Errorf("could not create a file: %s", err)
	}
	defer f.Close()
	f.WriteString(content)

	return nil
}

func getVerb(morphs []*Morph) string {
	for _, morph := range morphs {
		if morph.pos == "動詞" {
			return morph.base
		}
	}
	return ""
}

func getParticles(currentIdx int, passage []Chunk) string {
	var tmp string
	var buf bytes.Buffer
	for _, src := range passage[currentIdx].srcs {
		for _, morph := range passage[src].morphems {
			if morph.pos != "助詞" {
				continue
			}
			// Extract only last particle in the phrase
			tmp = morph.surface

		}
		if tmp == "" {
			continue
		}
		buf.WriteString(tmp + " ")
	}
	return strings.TrimRight(buf.String(), " ")
}
