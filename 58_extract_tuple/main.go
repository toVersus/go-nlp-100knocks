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

func (dep Dependencie) getVerb() []string {
	if dep.Type != "collapsed-dependencies" {
		return nil
	}

	nsubj := map[int]string{}
	for _, d := range dep.Deps {
		if d.Type != "nsubj" {
			continue
		}
		nsubj[d.Governor.Idx] = d.Governor.Value
	}

	// To store the keys in slice in sorted order
	var govIdxs []int
	for _, d := range dep.Deps {
		if d.Type != "dobj" {
			continue
		}

		if _, ok := nsubj[d.Governor.Idx]; ok {
			govIdxs = append(govIdxs, d.Governor.Idx)
		}
	}

	// Store verb in ascending order of index
	var verbs []string
	for _, i := range govIdxs {
		verbs = append(verbs, nsubj[i])
	}

	return verbs
}

func (dep Dependencie) getSubject(verbs []string) []string {
	var subjects []string
	for _, verb := range verbs {
		for _, d := range dep.Deps {
			if d.Type != "nsubj" {
				continue
			}
			if d.Governor.Value != verb {
				continue
			}
			subjects = append(subjects, d.Dependent.Value)
		}
	}
	return subjects
}

func (dep Dependencie) getObject(verbs []string) []string {
	var objects []string
	for _, verb := range verbs {
		for _, d := range dep.Deps {
			if d.Type != "dobj" {
				continue
			}
			if d.Governor.Value != verb {
				continue
			}
			objects = append(objects, d.Dependent.Value)
		}
	}
	return objects
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
		fmt.Fprintf(os.Stderr, "cannot find the specified file: %s\n  %s\n", filePath, err)
		os.Exit(1)
	}

	var buf bytes.Buffer
	for _, dep := range root.Document.Sentences[sentenceNum].Dependencies {
		verbs := dep.getVerb()
		if len(verbs) == 0 {
			continue
		}

		subjects := dep.getSubject(verbs)
		objects := dep.getObject(verbs)

		for i := 0; i < len(verbs); i++ {
			buf.WriteString(subjects[i] + "\t" + verbs[i] + "\t" + objects[i] + "\n")
		}
	}

	fmt.Println(strings.TrimRight(buf.String(), "\n"))
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
