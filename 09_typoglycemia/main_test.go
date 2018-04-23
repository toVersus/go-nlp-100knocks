package main

import (
	"reflect"
	"strings"
	"testing"
)

var shuffleTests = []struct {
	name   string
	str    string
	seed   int64
	expect string
}{
	{
		name:   "should predict the shuffled string.",
		str:    "This is a testing string.",
		seed:   10,
		expect: "testing a is This string.",
	},
}

func TestShuffle(t *testing.T) {
	for _, testcase := range shuffleTests {
		t.Log(testcase.name)
		words := strings.Fields(testcase.str)
		shuffle(len(words), testcase.seed, func(i, j int) {
			words[i], words[j] = words[j], words[i]
		})
		result := strings.Join(words, " ")
		if !reflect.DeepEqual(result, testcase.expect) {
			t.Errorf("result => %#v\n expect => %#v\n", result, testcase.expect)
		}
	}
}

func BenchmarkTypoglycemia(b *testing.B) {
	b.ResetTimer()
	for _, testcase := range shuffleTests {
		for i := 0; i < b.N; i++ {
			typoglycemia(testcase.str)
		}
	}
}
