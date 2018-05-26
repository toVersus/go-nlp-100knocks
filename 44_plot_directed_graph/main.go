package main

import (
	"bufio"
	"bytes"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/awalterschulze/gographviz"
)

func main() {
	var filePath, dotFilepath string
	var sentenceNum int
	flag.StringVar(&filePath, "file", "", "specify a file path")
	flag.StringVar(&filePath, "f", "", "specify a file path")
	flag.StringVar(&dotFilepath, "dot", "", "specify a output dot file path")
	flag.IntVar(&sentenceNum, "n", 1, "spacify number of sentences targeted for dependency analysis")
	flag.Parse()

	f, err := os.Open(filePath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "cannot find the specified file: %s\n  %s\n", filePath, err)
	}
	defer f.Close()

	if err := newChunkPassage(f).createDigraphDotFile(sentenceNum, dotFilepath); err != nil {
		log.Fatal("could not create digraph dot file: ", err)
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

// createDigraphDotFile creates dot file to visualize dependency tree as digraphs
func (chunkPassage ChunkPassage) createDigraphDotFile(sentenceNum int, filepath string) error {
	if len(chunkPassage) < 1 {
		return fmt.Errorf("could not get the meaningful results from dependency analysis")
	}

	var srcStr, dstStr string

	g := gographviz.NewGraph()
	if err := g.SetDir(true); err != nil {
		panic(err)
	}

	nodeAttrs := make(map[string]string)
	nodeAttrs["colorscheme"] = "rdylgn11"
	nodeAttrs["style"] = "\"solid,filled\""
	nodeAttrs["fontcolor"] = "black"
	nodeAttrs["fontname"] = "\"Migu 1M\""
	nodeAttrs["color"] = "10"
	nodeAttrs["fillcolor"] = "7"

	edgeAttrs := make(map[string]string)
	edgeAttrs["color"] = "black"

	for _, chunk := range chunkPassage[sentenceNum-1] {
		var srcWords, dstWords []string
		for _, src := range chunk.morphems {
			if src.pos != "記号" {
				srcWords = append(srcWords, src.surface)
			}
		}
		srcStr = strings.Join(srcWords, "")

		// Add all phrases as node
		if err := g.AddNode("G", srcStr, nodeAttrs); err != nil {
			panic(err)
		}

		if chunk.dst != -1 {
			for _, dst := range chunkPassage[sentenceNum-1][chunk.dst].morphems {
				if dst.pos != "記号" {
					dstWords = append(dstWords, dst.surface)
				}
			}
			dstStr = strings.Join(dstWords, "")

			// Add source-destination pair of phrases (src -> dst)
			if err := g.AddEdge(srcStr, dstStr, true, edgeAttrs); err != nil {
				panic(err)
			}
		}
	}

	f, err := os.Create(filepath)
	if err != nil {
		return fmt.Errorf("could not create a file: %s\n  %s", filepath, err)
	}
	f.WriteString(g.String())
	f.Close()

	return nil
}
