package main

import (
	"strings"
	"testing"

	"github.com/go-test/deep"
)

var stemmingTests = []struct {
	name   string
	text   string
	expect string
}{
	{
		name: "should slice the simple sentence into words",
		text: `Natural
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
languages.`,
		expect: `Natural	natur
language	languag
processing	process
(NLP)	(nlp)
is	is
a	a
field	field
of	of
computer	comput
science,	science,
artificial	artifici
intelligence,	intelligence,
and	and
linguistics	linguist
concerned	concern
with	with
the	the
interactions	interact
between	between
computers	comput
and	and
human	human
(natural)	(natural)
languages.	languages.`,
	},
}

func TestStemming(t *testing.T) {
	for _, testcase := range stemmingTests {
		t.Log(testcase.name)

		r := strings.NewReader(testcase.text)
		result := stemming(r)
		if diff := deep.Equal(result, testcase.expect); diff != nil {
			t.Error(diff)
		}
	}
}
