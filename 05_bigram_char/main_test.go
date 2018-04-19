package main

import (
	"reflect"
	"testing"
)

var biGramCharTests = []struct {
	name   string
	str    string
	expect []string
}{
	{
		name:   "should return bi-gram character",
		str:    "I am an NLPer",
		expect: []string{"Ia", "am", "ma", "an", "nN", "NL", "LP", "Pe", "er"},
	},
	{
		name:   "should return empty strings",
		str:    "",
		expect: []string{""},
	},
	{
		name:   "should return strings only containing one element",
		str:    "f",
		expect: []string{"f"},
	},
}

func TestBiGramChar(t *testing.T) {
	for _, testcase := range biGramCharTests {
		t.Log(testcase.name)
		if result := biGramChar(testcase.str); !reflect.DeepEqual(result, testcase.expect) {
			t.Errorf("result => %#v\n expect => %#v\n", result, testcase.expect)
		}
	}
}

func BenchmarkBiGramChar(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for _, testcase := range biGramCharTests {
			biGramChar(testcase.str)
		}
	}
}
