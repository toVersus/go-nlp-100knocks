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

type Node struct {
	Parent *Node
	Child  []*Node
	Pos    string
	Value  string
}

func (n *Node) walkNPString(buf *bytes.Buffer) {
	for i := len(n.Child) - 1; i >= 0; i-- {
		if n.Child[i].Value == "" || n.Child[i].Value == "," || n.Child[i].Value == "." {
			n.Child[i].walkNPString(buf)
			continue
		}
		if n.Pos != "NP" || n.Value != "" {
			n.Child[i].walkNPString(buf)
			continue
		}
		buf.WriteString(n.Child[i].Value + "\n")
		n.Child[i].walkNPString(buf)
	}
}

func New() *Node {
	return &Node{
		Parent: &Node{},
	}
}

func nextNode(parent *Node) *Node {
	return &Node{
		Parent: parent,
	}
}

func parse(str string) (*Node, error) {
	if str[0] != '(' {
		return nil, fmt.Errorf("Initial character must be '('\n  input string: %s", str)
	}
	return New().addChild(str[1:]), nil
}

func (n *Node) addChild(str string) *Node {
	if len(str) == 0 {
		return n
	}
	str = strings.TrimSpace(str)

	switch str[0] {
	case '(':
		node := nextNode(n)
		node = node.addChild(str[1:])
		n.Child = append(n.Child, node)
	case ')':
		n.Parent = n.Parent.addChild(str[1:])
	default:
		for i, s := range str {
			if (s != '(') && (s != ')') {
				continue
			}
			tmp := strings.Split(str[:i], " ")
			n.Pos = tmp[0]
			if len(tmp) == 2 {
				n.Value = tmp[1]
			}
			n = n.addChild(str[i:])
			break
		}
	}
	return n
}

func main() {
	var filePath string
	var sentenceNum int
	flag.StringVar(&filePath, "file", "", "specify a file path")
	flag.StringVar(&filePath, "f", "", "specify a file path")
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

	sexp := root.Document.Sentences[sentenceNum].Parse

	node, err := parse(sexp)

	if err != nil {
		fmt.Print(err)
		os.Exit(1)
	}

	var buf bytes.Buffer
	node.walkNPString(&buf)
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
