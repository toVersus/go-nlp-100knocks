package main

import (
	"strings"
	"testing"

	"github.com/go-test/deep"
)

var splitTests = []struct {
	name   string
	text   string
	expect string
}{
	{
		name: "should slice the simple sentence into words",
		text: `Natural language processing (NLP) is a field of computer science, artificial intelligence, and linguistics concerned with the interactions between computers and human (natural) languages.
As such, NLP is related to the area of humani-computer interaction.
Many challenges in NLP involve natural language understanding, that is, enabling computers to derive meaning from human or natural language input, and others involve natural language generation.
`,
		expect: `Natural
language
processing
(NLP)
is
a
field
of
computer
science,
artificial
intelligence,
and
linguistics
concerned
with
the
interactions
between
computers
and
human
(natural)
languages.

As
such,
NLP
is
related
to
the
area
of
humani-computer
interaction.

Many
challenges
in
NLP
involve
natural
language
understanding,
that
is,
enabling
computers
to
derive
meaning
from
human
or
natural
language
input,
and
others
involve
natural
language
generation.

`,
	},
	{
		name: "should slice the sentence containing full symbols into words",
		text: `The subfield of NLP devoted to learning approaches is known as Natural Language Learning (NLL) and its conference CoNLL and peak body SIGNLL are sponsored by ACL, recognizing also their links with Computational Linguistics and Language Acquisition.
When the aims of computational language learning research is to understand more about human language acquisition, or psycholinguistics, NLL overlaps into the related field of Computational Psycholinguistics.
`,
		expect: `The
subfield
of
NLP
devoted
to
learning
approaches
is
known
as
Natural
Language
Learning
(NLL)
and
its
conference
CoNLL
and
peak
body
SIGNLL
are
sponsored
by
ACL,
recognizing
also
their
links
with
Computational
Linguistics
and
Language
Acquisition.

When
the
aims
of
computational
language
learning
research
is
to
understand
more
about
human
language
acquisition,
or
psycholinguistics,
NLL
overlaps
into
the
related
field
of
Computational
Psycholinguistics.

`,
	},
}

func TestSplitByWord(t *testing.T) {
	for _, testcase := range splitTests {
		t.Log(testcase.name)

		r := strings.NewReader(testcase.text)
		result := splitByWord(r)
		if diff := deep.Equal(result, testcase.expect); diff != nil {
			t.Error(diff)
		}
	}
}

var result string

func BenchmarkSplitByWord(b *testing.B) {
	var s string
	for i := 0; i < b.N; i++ {
		for _, testcase := range splitTests {
			b.StopTimer()
			r := strings.NewReader(testcase.text)
			b.StartTimer()

			s = splitByWord(r)
		}
	}
	result = s
}
