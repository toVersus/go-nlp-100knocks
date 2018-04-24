package main

import (
	"math/rand"
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
		expect: "This is a tienstg stgrni.",
	},
}

func TestShuffle(t *testing.T) {
	for _, testcase := range shuffleTests {
		t.Log(testcase.name)

		ss := strings.Fields(testcase.str)
		typoglycemia := make([]string, len(ss))
		r := rand.New(rand.NewSource(testcase.seed))
		for i, word := range ss {
			if len(word) <= 4 {
				typoglycemia[i] = word
				continue
			}
			b := []byte(word[1:])
			r.Shuffle(len(b)-1, func(i, j int) {
				b[i], b[j] = b[j], b[i]
			})
			typoglycemia[i] = string(word[0]) + string(b)
		}
		result := strings.Join(typoglycemia, " ")
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

var selfImplShuffleTests = []struct {
	name   string
	str    string
	seed   int64
	expect string
}{
	{
		name:   "should predict the shuffled string.",
		str:    "This is a testing string.",
		seed:   10,
		expect: "This is a titseng snirtg.",
	},
}

func TestSelfImplShuffle(t *testing.T) {
	for _, testcase := range selfImplShuffleTests {
		t.Log(testcase.name)

		ss := strings.Fields(testcase.str)
		typoglycemia := make([]string, len(ss))
		for i, word := range ss {
			if len(word) <= 4 {
				typoglycemia[i] = word
				continue
			}
			b := []byte(word[1:])
			shuffle(len(b)-1, testcase.seed, func(i, j int) {
				b[i], b[j] = b[j], b[i]
			})
			typoglycemia[i] = string(word[0]) + string(b)
		}
		result := strings.Join(typoglycemia, " ")
		if !reflect.DeepEqual(result, testcase.expect) {
			t.Errorf("result => %#v\n expect => %#v\n", result, testcase.expect)
		}
	}
}

func BenchmarkSelfImplShuffle(b *testing.B) {
	b.ResetTimer()
	for _, testcase := range shuffleTests {
		for i := 0; i < b.N; i++ {
			typoglycemiaBySelfImplShuffle(testcase.str)
		}
	}
}
