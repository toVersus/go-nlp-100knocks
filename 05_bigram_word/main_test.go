package main

import (
	"reflect"
	"testing"
)

var biGramWordTests = []struct {
	name   string
	str    string
	expect []string
}{
	{
		name:   "should return bi-gram word",
		str:    "I am an NLPer",
		expect: []string{"I am", "am an", "an NLPer"},
	},
	{
		name:   "should return bi-gram word",
		str:    "I am an NLPer I am an NLPer I am an NLPer",
		expect: []string{"I am", "am an", "an NLPer", "NLPer I", "I am", "am an", "an NLPer", "NLPer I", "I am", "am an", "an NLPer"},
	},
	{
		name:   "should return empty strings",
		str:    "",
		expect: []string{""},
	},
	{
		name:   "should return strings only containing one element",
		str:    "foobar",
		expect: []string{"foobar"},
	},
}

func TestBigramWordByCopy(t *testing.T) {
	for _, testcase := range biGramWordTests {
		t.Log(testcase.name)
		if result := bigramWordByCopy(testcase.str); !reflect.DeepEqual(result, testcase.expect) {
			t.Errorf("result => %#v\n expect => %#v\n", result, testcase.expect)
		}
	}
}

func BenchmarkBigramWordByCopy(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for _, testcase := range biGramWordTests {
			bigramWordByCopy(testcase.str)
		}
	}
}

func TestBigramWordByConcat(t *testing.T) {
	for _, testcase := range biGramWordTests {
		t.Log(testcase.name)
		if result := bigramWordByConcat(testcase.str); !reflect.DeepEqual(result, testcase.expect) {
			t.Errorf("result => %#v\n expect => %#v\n", result, testcase.expect)
		}
	}
}

func BenchmarkBiGramWord(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for _, testcase := range biGramWordTests {
			bigramWordByConcat(testcase.str)
		}
	}
}

func TestBigramWordByJoin(t *testing.T) {
	for _, testcase := range biGramWordTests {
		t.Log(testcase.name)
		if result := bigramWordByJoin(testcase.str); !reflect.DeepEqual(result, testcase.expect) {
			t.Errorf("result => %#v\n expect => %#v\n", result, testcase.expect)
		}
	}
}

func BenchmarkBigramWordByJoin(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for _, testcase := range biGramWordTests {
			bigramWordByJoin(testcase.str)
		}
	}
}

func TestBigramWordByAppend(t *testing.T) {
	for _, testcase := range biGramWordTests {
		t.Log(testcase.name)
		if result := bigramWordByAppend(testcase.str); !reflect.DeepEqual(result, testcase.expect) {
			t.Errorf("result => %#v\n expect => %#v\n", result, testcase.expect)
		}
	}
}

func BenchmarkBigramWordByAppend(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for _, testcase := range biGramWordTests {
			bigramWordByAppend(testcase.str)
		}
	}
}
