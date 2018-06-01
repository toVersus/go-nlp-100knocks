package main

import (
	"strings"
	"testing"

	"github.com/go-test/deep"
)

var splitBySentTests = []struct {
	name   string
	text   string
	expect string
}{
	{
		name: "should split the simple text by sentence",
		text: `Natural language processing
From Wikipedia, the free encyclopedia

Natural language processing (NLP) is a field of computer science, artificial intelligence, and linguistics concerned with the interactions between computers and human (natural) languages. As such, NLP is related to the area of humani-computer interaction. Many challenges in NLP involve natural language understanding, that is, enabling computers to derive meaning from human or natural language input, and others involve natural language generation.
`,
		expect: `Natural language processing (NLP) is a field of computer science, artificial intelligence, and linguistics concerned with the interactions between computers and human (natural) languages.
As such, NLP is related to the area of humani-computer interaction.
Many challenges in NLP involve natural language understanding, that is, enabling computers to derive meaning from human or natural language input, and others involve natural language generation.
`,
	},
	{
		name: "should split the text containing full symbols by sentence",
		text: `The subfield of NLP devoted to learning approaches is known as Natural Language Learning (NLL) and its conference CoNLL and peak body SIGNLL are sponsored by ACL, recognizing also their links with Computational Linguistics and Language Acquisition. When the aims of computational language learning research is to understand more about human language acquisition, or psycholinguistics, NLL overlaps into the related field of Computational Psycholinguistics.
`,
		expect: `The subfield of NLP devoted to learning approaches is known as Natural Language Learning (NLL) and its conference CoNLL and peak body SIGNLL are sponsored by ACL, recognizing also their links with Computational Linguistics and Language Acquisition.
When the aims of computational language learning research is to understand more about human language acquisition, or psycholinguistics, NLL overlaps into the related field of Computational Psycholinguistics.
`,
	},
}

func TestSplitBySent(t *testing.T) {
	for _, testcase := range splitBySentTests {
		t.Log(testcase.name)

		r := strings.NewReader(testcase.text)
		results, err := splitBySent(r)
		if err != nil {
			t.Error(err)
		}
		if diff := deep.Equal(results, testcase.expect); diff != nil {
			t.Error(diff)
		}
	}
}
