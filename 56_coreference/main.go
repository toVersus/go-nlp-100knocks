package main

import (
	"bytes"
	"encoding/xml"
	"flag"
	"fmt"
	"io"
	"os"
	"strings"
)

type Root struct {
	Document *Document `xml:"document"`
}

type Document struct {
	Sentences    []*Sentence    `xml:"sentences>sentence"`
	Coreferences []*Coreference `xml:"coreference>coreference"`
}

// updateToRepresentiveMention replaces mention to representative mention in sentence
// and prints out by formatting the following rules.
// [Representative Mention] (Mention) ... .
func (doc Document) updateToRepresentiveMention() string {
	for _, coref := range doc.Coreferences {
		var repl string
		for _, mention := range coref.Mention {
			sentID, start, end := mention.Sentence-1, mention.Start, mention.End
			if mention.Representative == "true" {
				repl = mention.Text
				continue
			}
			doc.Sentences[sentID].Tokens[start-1].Word = "[" + repl + "] (" + doc.Sentences[sentID].Tokens[sentID].Word
			doc.Sentences[sentID].Tokens[end-2].Word += ")"
		}
	}

	var buf bytes.Buffer
	for _, s := range doc.Sentences {
		buf.WriteString(s.String() + "\n")
	}
	return strings.TrimRight(buf.String(), "\n")
}

type Sentence struct {
	ID           int             `xml:"id,attr"`
	Dependencies []*Dependencies `xml:"dependencies"`
	Parse        string          `xml:"parse,omitempty"`
	Tokens       []*Token        `xml:"tokens>token,omitempty"`
}

func (sent Sentence) String() string {
	var buf bytes.Buffer
	for _, token := range sent.Tokens {
		buf.WriteString(token.Word + " ")
	}
	return strings.TrimRight(buf.String(), " ")
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

func (token *Token) getWordTaggedByPos() string {
	return fmt.Sprintf("%s\t%s\t%s", token.Word, token.Lemma, token.POS)
}

func (token *Token) getPersonName() string {
	if token.NER == "PERSON" {
		return token.Word
	}
	return ""
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

type Dependencies struct {
	Type string `xml:"type,attr"`
	Dep  []*Dep `xml:"dep,omitempty"`
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
	var filePath, destFilePath string
	flag.StringVar(&filePath, "file", "", "specify a file path")
	flag.StringVar(&filePath, "f", "", "specify a file path")
	flag.StringVar(&destFilePath, "dest", "", "specify a dest file path")
	flag.StringVar(&destFilePath, "d", "", "specify a dest file path")
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

	text := root.Document.updateToRepresentiveMention()
	fmt.Print(text)
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
