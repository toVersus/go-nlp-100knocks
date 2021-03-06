package main

import (
	"strings"
	"testing"

	"github.com/go-test/deep"
)

var digGraphDotTests = []struct {
	name   string
	text   string
	expect string
}{
	{
		name: "should slice the simple sentence into words",
		text: `<?xml version="1.0" encoding="UTF-8"?>
<?xml-stylesheet href="CoreNLP-to-HTML.xsl" type="text/xsl"?>
<root>
	<document>
	<docId>test.txt</docId>
	<sentences>
		<sentence id="1">
		<tokens>
			<token id="1">
			<word>I</word>
			<lemma>I</lemma>
			<CharacterOffsetBegin>0</CharacterOffsetBegin>
			<CharacterOffsetEnd>1</CharacterOffsetEnd>
			<POS>PRP</POS>
			<NER>O</NER>
			<Speaker>PER0</Speaker>
			</token>
			<token id="2">
			<word>am</word>
			<lemma>be</lemma>
			<CharacterOffsetBegin>2</CharacterOffsetBegin>
			<CharacterOffsetEnd>4</CharacterOffsetEnd>
			<POS>VBP</POS>
			<NER>O</NER>
			<Speaker>PER0</Speaker>
			</token>
			<token id="3">
			<word>Bob</word>
			<lemma>Bob</lemma>
			<CharacterOffsetBegin>5</CharacterOffsetBegin>
			<CharacterOffsetEnd>8</CharacterOffsetEnd>
			<POS>NNP</POS>
			<NER>PERSON</NER>
			<Speaker>PER0</Speaker>
			</token>
			<token id="4">
			<word>.</word>
			<lemma>.</lemma>
			<CharacterOffsetBegin>8</CharacterOffsetBegin>
			<CharacterOffsetEnd>9</CharacterOffsetEnd>
			<POS>.</POS>
			<NER>O</NER>
			<Speaker>PER0</Speaker>
			</token>
		</tokens>
		<parse>(ROOT (S (NP (PRP I)) (VP (VBP am) (NP (NNP Bob))) (. .))) </parse>
		<dependencies type="basic-dependencies">
			<dep type="root">
			<governor idx="0">ROOT</governor>
			<dependent idx="3">Bob</dependent>
			</dep>
			<dep type="nsubj">
			<governor idx="3">Bob</governor>
			<dependent idx="1">I</dependent>
			</dep>
			<dep type="cop">
			<governor idx="3">Bob</governor>
			<dependent idx="2">am</dependent>
			</dep>
			<dep type="punct">
			<governor idx="3">Bob</governor>
			<dependent idx="4">.</dependent>
			</dep>
		</dependencies>
		<dependencies type="collapsed-dependencies">
			<dep type="root">
			<governor idx="0">ROOT</governor>
			<dependent idx="3">Bob</dependent>
			</dep>
			<dep type="nsubj">
			<governor idx="3">Bob</governor>
			<dependent idx="1">I</dependent>
			</dep>
			<dep type="cop">
			<governor idx="3">Bob</governor>
			<dependent idx="2">am</dependent>
			</dep>
			<dep type="punct">
			<governor idx="3">Bob</governor>
			<dependent idx="4">.</dependent>
			</dep>
		</dependencies>
		<dependencies type="collapsed-ccprocessed-dependencies">
			<dep type="root">
			<governor idx="0">ROOT</governor>
			<dependent idx="3">Bob</dependent>
			</dep>
			<dep type="nsubj">
			<governor idx="3">Bob</governor>
			<dependent idx="1">I</dependent>
			</dep>
			<dep type="cop">
			<governor idx="3">Bob</governor>
			<dependent idx="2">am</dependent>
			</dep>
			<dep type="punct">
			<governor idx="3">Bob</governor>
			<dependent idx="4">.</dependent>
			</dep>
		</dependencies>
		<dependencies type="enhanced-dependencies">
			<dep type="root">
			<governor idx="0">ROOT</governor>
			<dependent idx="3">Bob</dependent>
			</dep>
			<dep type="nsubj">
			<governor idx="3">Bob</governor>
			<dependent idx="1">I</dependent>
			</dep>
			<dep type="cop">
			<governor idx="3">Bob</governor>
			<dependent idx="2">am</dependent>
			</dep>
			<dep type="punct">
			<governor idx="3">Bob</governor>
			<dependent idx="4">.</dependent>
			</dep>
		</dependencies>
		<dependencies type="enhanced-plus-plus-dependencies">
			<dep type="root">
			<governor idx="0">ROOT</governor>
			<dependent idx="3">Bob</dependent>
			</dep>
			<dep type="nsubj">
			<governor idx="3">Bob</governor>
			<dependent idx="1">I</dependent>
			</dep>
			<dep type="cop">
			<governor idx="3">Bob</governor>
			<dependent idx="2">am</dependent>
			</dep>
			<dep type="punct">
			<governor idx="3">Bob</governor>
			<dependent idx="4">.</dependent>
			</dep>
		</dependencies>
		</sentence>
	</sentences>
	<coreference>
		<coreference>
		<mention representative="true">
			<sentence>1</sentence>
			<start>3</start>
			<end>4</end>
			<head>3</head>
			<text>Bob</text>
		</mention>
		<mention>
			<sentence>1</sentence>
			<start>1</start>
			<end>2</end>
			<head>1</head>
			<text>I</text>
		</mention>
		</coreference>
	</coreference>
	</document>
</root>`,
		expect: `digraph  {
	"Bob"->"ROOT"[ color=black ];
	"I"->"Bob"[ color=black ];
	"am"->"Bob"[ color=black ];
	"Bob" [ color=10, colorscheme=rdylgn11, fillcolor=7, fontcolor=black, fontname="Migu 1M", style="solid,filled" ];
	"I" [ color=10, colorscheme=rdylgn11, fillcolor=7, fontcolor=black, fontname="Migu 1M", style="solid,filled" ];
	"ROOT" [ color=10, colorscheme=rdylgn11, fillcolor=7, fontcolor=black, fontname="Migu 1M", style="solid,filled" ];
	"am" [ color=10, colorscheme=rdylgn11, fillcolor=7, fontcolor=black, fontname="Migu 1M", style="solid,filled" ];

}
`,
	},
}

func TestConvertToDiGraphDotFormat(t *testing.T) {
	for _, testcase := range digGraphDotTests {
		t.Log(testcase.name)

		r := strings.NewReader(testcase.text)
		root, err := readXML(r)
		if err != nil {
			t.Error(err)
		}

		var result string
		for _, dep := range root.Document.Sentences[0].Dependencies {
			result, err = dep.convertToDiGraphDotFormat(0)
			if err != nil {
				t.Error(err)
			}
			if result != "" {
				break
			}
		}
		if diff := deep.Equal(result, testcase.expect); diff != nil {
			t.Error(diff)
		}
	}
}
