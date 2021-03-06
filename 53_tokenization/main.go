package main

import (
	"encoding/xml"
	"flag"
	"fmt"
	"io"
	"os"
)

type Root struct {
	Document *Document `xml:"document"`
}

type Document struct {
	Sentences    *Sentences    `xml:"sentences,omitempty"`
	Coreferences *Coreferences `xml:"coreference,omitempty"`
}

type Sentences struct {
	Sentence []*Sentence `xml:"sentence,omitempty"`
}

type Sentence struct {
	ID           string          `xml:"id,attr"`
	Dependencies []*Dependencies `xml:"dependencies"`
	Parse        *Parse          `xml:"parse,omitempty"`
	Tokens       Tokens          `xml:"tokens,omitempty"`
}

type Tokens struct {
	Token []*Token `xml:"token,omitempty"`
}

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
	Value string
}

type Parse struct {
	Value string
}

type Governor struct {
	Copy  string `xml:"copy,attr"`
	Idx   string `xml:"idx,attr"`
	Value string
}

type Dependent struct {
	Copy  string `xml:"copy,attr"`
	Idx   string `xml:"idx,attr"`
	Value string
}

type Dep struct {
	Extra     string     `xml:"extra,attr"`
	Type      string     `xml:"type,attr"`
	Dependent *Dependent `xml:"dependent,omitempty"`
	Governor  *Governor  `xml:"governor,omitempty"`
}

type Dependencies struct {
	Type string `xml:"type,attr"`
	Dep  []*Dep `xml:"dep,omitempty"`
}

type Coreferences struct {
	Coreference []*Coreference `xml:"coreference,omitempty"`
}

type Coreference struct {
	Mention []*Mention `xml:"mention,omitempty"`
}

type Mention struct {
	Representative string `xml:"representative,attr"`
	Sentence       int    `xml:"sentence,omitempty"`
	Start          int    `xml:"start,omitempty"`
	End            int    `xml:"end,omitempty"`
	Head           int    `xml:"head,omitempty"`
	Text           string `xml:"text,omitempty"`
}

func main() {
	var filePath string
	flag.StringVar(&filePath, "file", "", "specify a file path")
	flag.StringVar(&filePath, "f", "", "specify a file path")
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

	for _, s := range root.Document.Sentences.Sentence {
		for _, t := range s.Tokens.Token {
			fmt.Println(t.Word)
		}
	}
}

// readXML reads the result of Stanford Core NLP and initiate the Root struct
func readXML(r io.Reader) (*Root, error) {
	root := Root{}

	dec := xml.NewDecoder(r)
	if err := dec.Decode(&root); err != nil {
		return nil, err
	}

	return &root, nil
}
