package main

import (
	"bufio"
	"bytes"
	"flag"
	"fmt"
	"io"
	"os"
	"sort"
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

	if _, err := os.Stat(filePath); err != nil {
		fmt.Printf("cannot find the specified file: %s\n  %s\n", filePath, err)
	}

	content := newChunkPassage(f).stringifyNounDstPath()
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

// Morphs represents list of Morph.
type Morphs []*Morph

// String returns the surface of Morphs trimming the symbol.
func (morphs Morphs) String() string {
	var buf bytes.Buffer
	for _, morph := range morphs {
		if morph.pos == "記号" {
			continue
		}
		buf.WriteString(morph.surface)
	}
	return buf.String()
}

// StringWithMask returns the masked surface of Morphs trimming the symbol.
func (morphs Morphs) StringWithMask(mask string) string {
	var buf bytes.Buffer
	isFirst := true
	for _, morph := range morphs {
		if morph.pos == "記号" {
			continue
		}
		if morph.pos == "名詞" && isFirst {
			buf.WriteString(mask)
			isFirst = false
			continue
		}
		buf.WriteString(morph.surface)
	}
	return buf.String()
}

// Chunk represents the list of Morphs and number of phrases and its depended phrases.
type Chunk struct {
	morphems Morphs
	dst      int
	srcs     []int
}

// Passage represents the bunch of phrases.
type Passage []Chunk

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
		// Catch the line of Morpheme analysis.
		if morph := newMorph(sc.Text()); morph != nil {
			chunk.morphems = append(chunk.morphems, morph)
			continue
		}

		// Catch the line of Dependency analysis
		if strings.HasPrefix(sc.Text(), "*") {
			// Append the chunk assigned the result of morpheme analysis in previous phrase.
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

		// Assign the src index number.
		for index, srcIndex := range dict {
			if srcIndex == -1 {
				break
			}
			passage[srcIndex].srcs = append(passage[srcIndex].srcs, index)
		}

		chunkPassage = append(chunkPassage, passage)

		// Initialize the variables.
		dict = dict[0:0]
		passage = nil
		chunk = Chunk{nil, 0, nil}
	}

	return &chunkPassage
}

// stringifyNounDstPath tracks all the depended phrases from parsed noun phrases.
func (chunkPassage ChunkPassage) stringifyNounDstPath() string {
	var buf bytes.Buffer
	for _, passage := range chunkPassage {
		indexes := passage.getNounChunkIndex()
		for i := 0; i < len(indexes); i++ {
			for j := i + 1; j < len(indexes); j++ {
				commonIndex, isGettingOnSamePath := passage.getCommonLeaf(indexes[i], indexes[j])
				if commonIndex == -1 {
					continue
				}
				if isGettingOnSamePath {
					passage.bufferPhrasesOnSameRoute(&buf, indexes[i], commonIndex)
					continue
				}
				passage.bufferPhrasesOnOtherRoute(&buf, indexes[i], indexes[j], commonIndex)
			}
		}
	}
	return buf.String()
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

func (passage Passage) getPathToRoot(index int, path treePath) treePath {
	path.Add(index)
	if passage[index].dst == -1 {
		return path
	}
	path.Add(passage[index].dst)
	return passage.getPathToRoot(passage[index].dst, path)
}

// getCommonLeaf compares the tree path of chunks selected by input indexes,
// and returns index of common leaf and check whether pair of paths are on the same route or not.
// When there is no common leaf, it returns -1 and also false.
func (passage Passage) getCommonLeaf(index, nextIndex int) (int, bool) {
	indexPath, nextIndexPath := treePath{}, treePath{}
	indexPath = passage.getPathToRoot(index, indexPath)
	nextIndexPath = passage.getPathToRoot(nextIndex, nextIndexPath)

	inter := indexPath.Intersect(nextIndexPath)
	if len(inter) == 0 {
		return -1, false
	}

	var keys []int
	for k := range inter {
		keys = append(keys, k)
	}
	sort.Ints(keys)
	return keys[0], indexPath.Include(nextIndexPath)
}

func (passage Passage) getChunkStrOnRoute(chunkStrs []string, index, commonIndex int) []string {
	nextIndex := passage[index].dst
	if nextIndex == commonIndex {
		return chunkStrs
	}
	chunkStrs = append(chunkStrs, passage[nextIndex].morphems.String())
	return passage.getChunkStrOnRoute(chunkStrs, nextIndex, commonIndex)
}

func (passage Passage) getNounChunkIndex() []int {
	var index []int
	for currentIndexInPassage, chunk := range passage {
		for _, morph := range chunk.morphems {
			if morph.pos != "名詞" {
				continue
			}
			index = append(index, currentIndexInPassage)
			break
		}
	}
	return index
}

// ChunkPassage represents the bunch of passages.
type ChunkPassage []Passage

func (passage Passage) bufferPhrasesOnSameRoute(buf *bytes.Buffer, index, commonIndex int) {
	headNoun := passage[index].morphems.StringWithMask("X")
	buf.WriteString(headNoun + " -> ")
	var chunkStrs []string
	chunkStrs = passage.getChunkStrOnRoute(chunkStrs, index, commonIndex)
	if len(chunkStrs) != 0 {
		buf.WriteString(strings.Join(chunkStrs, " -> ") + " -> ")
	}
	buf.WriteString("Y\n")
}

func (passage Passage) bufferPhrasesOnOtherRoute(buf *bytes.Buffer, index, nextIndex, commonIndex int) {
	passage.bufferSidePhrases(buf, index, commonIndex, "X")
	passage.bufferSidePhrases(buf, nextIndex, commonIndex, "Y")
	buf.WriteString(passage[commonIndex].morphems.String() + "\n")
}

func (passage Passage) bufferSidePhrases(buf *bytes.Buffer, index, commonIndex int, mask string) {
	headIndexNoun := passage[index].morphems.StringWithMask(mask)
	buf.WriteString(headIndexNoun)
	var indexchunkStrs []string
	indexchunkStrs = passage.getChunkStrOnRoute(indexchunkStrs, index, commonIndex)
	if len(indexchunkStrs) == 1 {
		buf.WriteString(" -> " + strings.Join(indexchunkStrs, ""))
	} else if len(indexchunkStrs) > 1 {
		buf.WriteString(" -> " + strings.Join(indexchunkStrs, " -> "))
	}
	buf.WriteString(" | ")
}

type treePath map[int]struct{}

func (path treePath) Add(n int) treePath {
	path[n] = struct{}{}
	return path
}

func (path treePath) Contains(n int) bool {
	if _, ok := path[n]; ok {
		return true
	}
	return false
}

func (path treePath) Include(other treePath) bool {
	for i := range other {
		if _, ok := path[i]; !ok {
			return false
		}
	}
	return true
}

func (path treePath) Intersect(other treePath) treePath {
	intersect := make(treePath)
	for i := range path {
		if !other.Contains(i) {
			continue
		}
		intersect.Add(i)
	}
	return intersect
}

func (path treePath) Difference(other treePath) treePath {
	differ := make(treePath)
	for i := range path {
		if other.Contains(i) {
			continue
		}
		differ.Add(i)
	}
	return differ
}
