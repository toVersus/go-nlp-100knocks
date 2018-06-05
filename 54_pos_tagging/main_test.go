package main

import (
	"bytes"
	"reflect"
	"strings"
	"testing"
)

var posTaggingTests = []struct {
	name   string
	text   string
	expect string
}{
	{
		name: "should return pos representation with tab-delimited format",
		text: `<?xml version="1.0" encoding="UTF-8"?>
<?xml-stylesheet href="CoreNLP-to-HTML.xsl" type="text/xsl"?>
<root>
	<document>
	<docId>test.txt</docId>
	<sentences>
		<sentence id="1">
		<tokens>
			<token id="1">
			<word>This</word>
			<lemma>this</lemma>
			<CharacterOffsetBegin>0</CharacterOffsetBegin>
			<CharacterOffsetEnd>4</CharacterOffsetEnd>
			<POS>DT</POS>
			<NER>O</NER>
			<Speaker>PER0</Speaker>
			</token>
			<token id="2">
			<word>is</word>
			<lemma>be</lemma>
			<CharacterOffsetBegin>5</CharacterOffsetBegin>
			<CharacterOffsetEnd>7</CharacterOffsetEnd>
			<POS>VBZ</POS>
			<NER>O</NER>
			<Speaker>PER0</Speaker>
			</token>
			<token id="3">
			<word>a</word>
			<lemma>a</lemma>
			<CharacterOffsetBegin>8</CharacterOffsetBegin>
			<CharacterOffsetEnd>9</CharacterOffsetEnd>
			<POS>DT</POS>
			<NER>O</NER>
			<Speaker>PER0</Speaker>
			</token>
			<token id="4">
			<word>pen</word>
			<lemma>pen</lemma>
			<CharacterOffsetBegin>10</CharacterOffsetBegin>
			<CharacterOffsetEnd>13</CharacterOffsetEnd>
			<POS>NN</POS>
			<NER>O</NER>
			<Speaker>PER0</Speaker>
			</token>
			<token id="5">
			<word>.</word>
			<lemma>.</lemma>
			<CharacterOffsetBegin>13</CharacterOffsetBegin>
			<CharacterOffsetEnd>14</CharacterOffsetEnd>
			<POS>.</POS>
			<NER>O</NER>
			<Speaker>PER0</Speaker>
			</token>
		</tokens>
		<parse>(ROOT (S (NP (DT This)) (VP (VBZ is) (NP (DT a) (NN pen))) (. .))) </parse>
		<dependencies type="basic-dependencies">
			<dep type="root">
			<governor idx="0">ROOT</governor>
			<dependent idx="4">pen</dependent>
			</dep>
			<dep type="nsubj">
			<governor idx="4">pen</governor>
			<dependent idx="1">This</dependent>
			</dep>
			<dep type="cop">
			<governor idx="4">pen</governor>
			<dependent idx="2">is</dependent>
			</dep>
			<dep type="det">
			<governor idx="4">pen</governor>
			<dependent idx="3">a</dependent>
			</dep>
			<dep type="punct">
			<governor idx="4">pen</governor>
			<dependent idx="5">.</dependent>
			</dep>
		</dependencies>
		<dependencies type="collapsed-dependencies">
			<dep type="root">
			<governor idx="0">ROOT</governor>
			<dependent idx="4">pen</dependent>
			</dep>
			<dep type="nsubj">
			<governor idx="4">pen</governor>
			<dependent idx="1">This</dependent>
			</dep>
			<dep type="cop">
			<governor idx="4">pen</governor>
			<dependent idx="2">is</dependent>
			</dep>
			<dep type="det">
			<governor idx="4">pen</governor>
			<dependent idx="3">a</dependent>
			</dep>
			<dep type="punct">
			<governor idx="4">pen</governor>
			<dependent idx="5">.</dependent>
			</dep>
		</dependencies>
		<dependencies type="collapsed-ccprocessed-dependencies">
			<dep type="root">
			<governor idx="0">ROOT</governor>
			<dependent idx="4">pen</dependent>
			</dep>
			<dep type="nsubj">
			<governor idx="4">pen</governor>
			<dependent idx="1">This</dependent>
			</dep>
			<dep type="cop">
			<governor idx="4">pen</governor>
			<dependent idx="2">is</dependent>
			</dep>
			<dep type="det">
			<governor idx="4">pen</governor>
			<dependent idx="3">a</dependent>
			</dep>
			<dep type="punct">
			<governor idx="4">pen</governor>
			<dependent idx="5">.</dependent>
			</dep>
		</dependencies>
		<dependencies type="enhanced-dependencies">
			<dep type="root">
			<governor idx="0">ROOT</governor>
			<dependent idx="4">pen</dependent>
			</dep>
			<dep type="nsubj">
			<governor idx="4">pen</governor>
			<dependent idx="1">This</dependent>
			</dep>
			<dep type="cop">
			<governor idx="4">pen</governor>
			<dependent idx="2">is</dependent>
			</dep>
			<dep type="det">
			<governor idx="4">pen</governor>
			<dependent idx="3">a</dependent>
			</dep>
			<dep type="punct">
			<governor idx="4">pen</governor>
			<dependent idx="5">.</dependent>
			</dep>
		</dependencies>
		<dependencies type="enhanced-plus-plus-dependencies">
			<dep type="root">
			<governor idx="0">ROOT</governor>
			<dependent idx="4">pen</dependent>
			</dep>
			<dep type="nsubj">
			<governor idx="4">pen</governor>
			<dependent idx="1">This</dependent>
			</dep>
			<dep type="cop">
			<governor idx="4">pen</governor>
			<dependent idx="2">is</dependent>
			</dep>
			<dep type="det">
			<governor idx="4">pen</governor>
			<dependent idx="3">a</dependent>
			</dep>
			<dep type="punct">
			<governor idx="4">pen</governor>
			<dependent idx="5">.</dependent>
			</dep>
		</dependencies>
		</sentence>
	</sentences>
	<coreference>
		<coreference>
		<mention representative="true">
			<sentence>1</sentence>
			<start>3</start>
			<end>5</end>
			<head>4</head>
			<text>a pen</text>
		</mention>
		<mention>
			<sentence>1</sentence>
			<start>1</start>
			<end>2</end>
			<head>1</head>
			<text>This</text>
		</mention>
		</coreference>
	</coreference>
	</document>
</root>`,
		expect: "This	this	DT\nis	be	VBZ\na	a	DT\npen	pen	NN\n.	.	.",
	},
}

func TestPosTagging(t *testing.T) {
	for _, testcase := range posTaggingTests {
		t.Log(testcase.name)

		r := strings.NewReader(testcase.text)
		root, err := readXML(r)
		if err != nil {
			t.Error(err)
		}

		var buf bytes.Buffer
		for _, sent := range root.Document.Sentences.Sentence {
			for _, token := range sent.Tokens.Token {
				buf.WriteString(token.String() + "\n")
			}
		}

		result := strings.TrimRight(buf.String(), "\n")
		if !reflect.DeepEqual(result, testcase.expect) {
			t.Errorf("node result => %#v\n  should contain => %#v\n", result, testcase.expect)
		}
	}
}
