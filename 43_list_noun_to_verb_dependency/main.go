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
	var filePath string
	flag.StringVar(&filePath, "file", "", "specify a file path")
	flag.StringVar(&filePath, "f", "", "specify a file path")
	flag.Parse()

	f, err := os.Open(filePath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "cannot find the specified file: %s\n  %s\n", filePath, err)
	}
	defer f.Close()

	fmt.Println(newChunkPassage(f).listNounToVerbDependency())
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

func (chunk Chunk) String() string {
	var buf bytes.Buffer
	for _, morph := range chunk.morphems {
		buf.WriteString(fmt.Sprintf("%+v", morph))
	}
	return strings.TrimRight(buf.String()[2:], "}")
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

func (chunkPassage ChunkPassage) listNounToVerbDependency() []string {
	var words, texts []string
	var hasNounInSrcPhrases, hasVerbInDstPhrases = false, false
	for _, passage := range chunkPassage {
		for _, chunk := range passage {
			if chunk.dst != -1 {
				words = nil
				hasNounInSrcPhrases, hasVerbInDstPhrases = false, false

				for _, srcMorph := range chunk.morphems {
					if srcMorph.pos != "記号" {
						words = append(words, srcMorph.surface)
					}

					if srcMorph.pos == "名詞" {
						hasNounInSrcPhrases = true
					}
				}

				for _, dstMorph := range passage[chunk.dst].morphems {
					if dstMorph.pos != "記号" {
						words = append(words, dstMorph.surface)
					}

					if dstMorph.pos == "動詞" {
						hasVerbInDstPhrases = true
					}
				}

				if (hasNounInSrcPhrases == true) && (hasVerbInDstPhrases == true) {
					texts = append(texts, strings.Join(words, "\t"))
				}
			}
		}
	}

	return texts
}
