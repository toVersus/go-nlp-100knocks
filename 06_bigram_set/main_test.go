package main

import (
	"reflect"
	"testing"
)

var newBigramTests = []struct {
	name   string
	str    string
	expect Bigram
}{
	{
		name: "should return new bigram",
		str:  "paraparaparadise",
		expect: Bigram{
			"pa": struct{}{},
			"ar": struct{}{},
			"ra": struct{}{},
			"ap": struct{}{},
			"ad": struct{}{},
			"di": struct{}{},
			"is": struct{}{},
			"se": struct{}{},
		},
	},
}

func TestNewBigram(t *testing.T) {
	for _, testcase := range newBigramTests {
		t.Log(testcase.name)
		result := NewBigram(testcase.str)
		for i := range result {
			if !reflect.DeepEqual(result[i], testcase.expect[i]) {
				t.Errorf("result => %#v\n expect => %#v\n", result, testcase.expect)
			}
		}
	}
}

func BenchmarkNewBiGram(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for _, testcase := range newBigramTests {
			NewBigram(testcase.str)
		}
	}
}

var addTests = []struct {
	name   string
	bg     Bigram
	str    string
	expect Bigram
}{
	{
		name: "successfully add new element into the set",
		bg: Bigram{
			"pa": struct{}{},
		},
		str: "ar",
		expect: Bigram{
			"pa": struct{}{},
			"ar": struct{}{},
		},
	},
	{
		name: "successfully add new element into empty set",
		bg:   Bigram{},
		str:  "ar",
		expect: Bigram{
			"ar": struct{}{},
		},
	},
}

func TestAdd(t *testing.T) {
	for _, testcase := range addTests {
		t.Log(testcase.name)
		// Extract original type using type assertion
		result := testcase.bg.Add(testcase.str).(Bigram)
		for i := range result {
			if !reflect.DeepEqual(result[i], testcase.expect[i]) {
				t.Errorf("result => %#v\n expect => %#v\n", result, testcase.expect)
			}
		}
	}
}

func BenchmarkAdd(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for _, testcase := range addTests {
			testcase.bg.Add(testcase.str)
		}
	}
}

var containTests = []struct {
	name   string
	bg     Bigram
	str    string
	expect bool
}{
	{
		name: "should return true",
		bg: Bigram{
			"pa": struct{}{},
			"ar": struct{}{},
		},
		str:    "pa",
		expect: true,
	},
	{
		name: "should return false",
		bg: Bigram{
			"pa": struct{}{},
			"ar": struct{}{},
		},
		str:    "aa",
		expect: false,
	},
}

func TestContains(t *testing.T) {
	for _, testcase := range containTests {
		t.Log(testcase.name)
		if result := testcase.bg.Contains(testcase.str); !reflect.DeepEqual(result, testcase.expect) {
			t.Errorf("result => %#v\n expect => %#v\n", result, testcase.expect)
		}
	}
}

func BenchmarkContains(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for _, testcase := range containTests {
			testcase.bg.Contains(testcase.str)
		}
	}
}

var unionTests = []struct {
	name   string
	bg     Bigram
	other  Bigram
	expect Bigram
}{
	{
		name: "successfully combined two maps",
		bg: Bigram{
			"pa": struct{}{},
			"ar": struct{}{},
		},
		other: Bigram{
			"ap": struct{}{},
			"ra": struct{}{},
		},
		expect: Bigram{
			"pa": struct{}{},
			"ar": struct{}{},
			"ap": struct{}{},
			"ra": struct{}{},
		},
	},
	{
		name: "successfully combined an empty map with a map",
		bg:   Bigram{},
		other: Bigram{
			"ap": struct{}{},
			"ra": struct{}{},
		},
		expect: Bigram{
			"ap": struct{}{},
			"ra": struct{}{},
		},
	},
}

func TestUnion(t *testing.T) {
	for _, testcase := range unionTests {
		t.Log(testcase.name)
		result := testcase.bg.Union(testcase.other).(Bigram)
		for i := range result {
			if !reflect.DeepEqual(result[i], testcase.expect[i]) {
				t.Errorf("result => %#v\n expect => %#v\n", result, testcase.expect)
			}
		}
	}
}

func BenchmarkUnion(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for _, testcase := range unionTests {
			testcase.bg.Union(testcase.other)
		}
	}
}

var intersectTests = []struct {
	name   string
	bg     Bigram
	other  Bigram
	expect Bigram
}{
	{
		name: "Successfully intersect maps",
		bg: Bigram{
			"pa": struct{}{},
			"ar": struct{}{},
		},
		other: Bigram{
			"ap": struct{}{},
			"ra": struct{}{},
		},
		expect: Bigram{
			"": struct{}{},
		},
	},
	{
		name: "Successfully intersect maps",
		bg: Bigram{
			"pa": struct{}{},
			"ar": struct{}{},
		},
		other: Bigram{
			"pa": struct{}{},
			"ar": struct{}{},
		},
		expect: Bigram{
			"pa": struct{}{},
			"ar": struct{}{},
		},
	},
}

func TestIntersect(t *testing.T) {
	for _, testcase := range intersectTests {
		t.Log(testcase.name)
		result := testcase.bg.Intersect(testcase.other).(Bigram)
		for i := range result {
			if !reflect.DeepEqual(result[i], testcase.expect[i]) {
				t.Errorf("result => %#v\n expect => %#v\n", result, testcase.expect)
			}
		}
	}
}

func BenchmarkIntersect(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for _, testcase := range intersectTests {
			testcase.bg.Intersect(testcase.other)
		}
	}
}

var differenceTests = []struct {
	name   string
	bg     Bigram
	other  Bigram
	expect Bigram
}{
	{
		name: "should return empty map",
		bg: Bigram{
			"pa": struct{}{},
			"ar": struct{}{},
		},
		other: Bigram{
			"pa": struct{}{},
			"ar": struct{}{},
		},
		expect: Bigram{
			"": struct{}{},
		},
	},
	{
		name: "should return difference between A and B",
		bg: Bigram{
			"pa": struct{}{},
			"ar": struct{}{},
		},
		other: Bigram{
			"pa": struct{}{},
			"ra": struct{}{},
		},
		expect: Bigram{
			"ar": struct{}{},
		},
	},
}

func TestDifference(t *testing.T) {
	for _, testcase := range differenceTests {
		t.Log(testcase.name)
		result := testcase.bg.Intersect(testcase.other).(Bigram)
		for i := range result {
			if !reflect.DeepEqual(result[i], testcase.expect[i]) {
				t.Errorf("result => %#v\n expect => %#v\n", result, testcase.expect)
			}
		}
	}
}

func BenchmarkDifference(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for _, testcase := range differenceTests {
			testcase.bg.Difference(testcase.other)
		}
	}
}
