package main

import (
	"encoding/xml"
	"flag"
	"fmt"
	"io"
	"os"
	"strconv"

	"github.com/awalterschulze/gographviz"
)

type Root struct {
	Document *Document `xml:"document"`
}

type Document struct {
	Sentences    Sentences    `xml:"sentences>sentence"`
	Coreferences Coreferences `xml:"coreference>coreference"`
}

type Sentences []*Sentence

type Coreferences []*Coreference

type Sentence struct {
	ID           int          `xml:"id,attr"`
	Dependencies Dependencies `xml:"dependencies"`
	Parse        string       `xml:"parse,omitempty"`
	Tokens       Tokes        `xml:"tokens>token,omitempty"`
}

type Dependencies []*Dependencie

type Tokes []*Token

type Token struct {
	ID                   string `xml:"id,attr"`
	Word                 string `xml:"word,omitempty"`
	Lemma                string `xml:"lemma,omitempty"`
	CharacterOffsetBegin int    `xml:"CharacterOffsetBegin,omitempty"`
	CharacterOffsetEnd   int    `xml:"CharacterOffsetEnd,omitempty"`
	POS                  string `xml:"POS,omitempty"`
	NER                  string `xml:"NER,omitempty"`
	NormalizedNER        string `xml:"NormalizedNER,omitempty"`
	Speaker              string `xml:"Speaker,omitempty"`
	Timex                *Timex `xml:"Timex,omitempty"`
}

type Timex struct {
	Tid   string `xml:"tid,attr"`
	Type  string `xml:"type,attr"`
	Value string `xml:",chardata"`
}

type Governor struct {
	Copy  string `xml:"copy,attr"`
	Idx   int    `xml:"idx,attr"`
	Value string `xml:",chardata"`
}

type Dependent struct {
	Copy  string `xml:"copy,attr"`
	Idx   int    `xml:"idx,attr"`
	Value string `xml:",chardata"`
}

type Dep struct {
	Extra     string     `xml:"extra,attr"`
	Type      string     `xml:"type,attr"`
	Dependent *Dependent `xml:"dependent,omitempty"`
	Governor  *Governor  `xml:"governor,omitempty"`
}

type Dependencie struct {
	Type string `xml:"type,attr"`
	Deps Deps   `xml:"dep,omitempty"`
}

// convertToDiGraphDotFormat returns Digraph dot format with some hardcoded options.
func (dep Dependencie) convertToDiGraphDotFormat(sentenceNum int) (string, error) {
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

	if dep.Type != "collapsed-dependencies" {
		return "", nil
	}

	for _, d := range dep.Deps {
		if d.Type == "punct" {
			continue
		}

		if err := g.AddNode("G", strconv.Quote(d.Dependent.Value), nodeAttrs); err != nil {
			return "", fmt.Errorf("could not add dependent word into the node:\n  %s", err)
		}
		if err := g.AddNode("G", strconv.Quote(d.Governor.Value), nodeAttrs); err != nil {
			return "", fmt.Errorf("could not add governor word into the node:\n  %s", err)
		}

		if err := g.AddEdge(strconv.Quote(d.Dependent.Value), strconv.Quote(d.Governor.Value), true, edgeAttrs); err != nil {
			return "", fmt.Errorf("could not add dependent and governor word into the edge:\n  %s", err)
		}
	}

	return g.String(), nil
}

type Deps []*Dep

type Coreference struct {
	Mentions Mentions `xml:"mention,omitempty"`
}

type Mentions []*Mention

type Mention struct {
	Representative string `xml:"representative,attr"`
	Sentence       int    `xml:"sentence,omitempty"`
	Start          int    `xml:"start,omitempty"`
	End            int    `xml:"end,omitempty"`
	Head           int    `xml:"head,omitempty"`
	Text           string `xml:"text,omitempty"`
}

func main() {
	var filePath, destFilePath string
	var sentenceNum int
	flag.StringVar(&filePath, "file", "", "specify a file path")
	flag.StringVar(&filePath, "f", "", "specify a file path")
	flag.StringVar(&destFilePath, "dest", "", "specify a dest file path")
	flag.StringVar(&destFilePath, "d", "", "specify a dest file path")
	flag.IntVar(&sentenceNum, "n", 1, "specify number of sentence")
	flag.Parse()

	f, err := os.Open(filePath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "cannot find the specified file: %s\n  %s\n", filePath, err)
		os.Exit(1)
	}
	defer f.Close()

	root, err := readXML(f)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	for _, dep := range root.Document.Sentences[sentenceNum].Dependencies {
		text, err := dep.convertToDiGraphDotFormat(sentenceNum)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Failed to convert to digraph dot format: %s", err)
			os.Exit(1)
		}
		if text == "" {
			continue
		}
		output(destFilePath, text)
	}
}

// readXML reads the result of Stanford Core NLP and constructs the Root struct
func readXML(r io.Reader) (*Root, error) {
	root := Root{}

	dec := xml.NewDecoder(r)
	if err := dec.Decode(&root); err != nil {
		return nil, err
	}

	return &root, nil
}

// output just creates a file with given contents.
func output(filepath, content string) error {
	f, err := os.Create(filepath)
	if err != nil {
		return fmt.Errorf("could not create a file: %s", err)
	}
	defer f.Close()
	f.WriteString(content)

	return nil
}
