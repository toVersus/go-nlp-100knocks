package main

import (
	"reflect"
	"strings"
	"testing"
)

var countLenTests = []struct {
	name   string
	str    string
	expect []int
}{
	{
		name:   "should return int array, which element is equal to pi-digit number",
		str:    "Now I need a drink, alcoholic of course, after the heavy lectures involving quantum mechanics.",
		expect: []int{3, 1, 4, 1, 5, 9, 2, 6, 5, 3, 5, 8, 9, 7, 9},
	},
	{
		name:   "should return zero int array",
		str:    "",
		expect: []int{0},
	},
}

func TestCountWordLenByCounter(t *testing.T) {
	for _, testcase := range countLenTests {
		if result := countWordLenByCounter(testcase.str); !reflect.DeepEqual(result, testcase.expect) {
			t.Errorf("result => %#v\n expect => %#v\n", result, testcase.expect)
		}
	}
}

func BenchmarkCountWordLenByCounter(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for _, testcase := range countLenTests {
			countWordLenByCounter(testcase.str)
		}
	}
}

func TestCountWordLenByRecursiveFunc(t *testing.T) {
	for _, testcase := range countLenTests {
		t.Log(testcase.name)
		testcase.str = strings.Replace(testcase.str, ".", "", -1)
		testcase.str = strings.Replace(testcase.str, ",", "", -1)
		if result := countWordLenByRecursiveFunc(strings.Split(testcase.str, " ")); !reflect.DeepEqual(result, testcase.expect) {
			t.Errorf("result => %#v\n expect => %#v\n", result, testcase.expect)
		}

	}
}

func BenchmarkCountWordLenByRecursiveFunc(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for _, testcase := range countLenTests {
			testcase.str = strings.Replace(testcase.str, ".", "", -1)
			testcase.str = strings.Replace(testcase.str, ",", "", -1)
			countWordLenByRecursiveFunc(strings.Split(testcase.str, " "))
		}
	}
}

func TestCountWordLen(t *testing.T) {
	for _, testcase := range countLenTests {
		t.Log(testcase.name)
		testcase.str = strings.Replace(testcase.str, ".", "", -1)
		testcase.str = strings.Replace(testcase.str, ",", "", -1)
		if result := countWordLen(testcase.str); !reflect.DeepEqual(result, testcase.expect) {
			t.Errorf("result => %#v\n expect => %#v\n", result, testcase.expect)
		}
	}
}

func BenchmarkCountWordLen(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for _, testcase := range countLenTests {
			testcase.str = strings.Replace(testcase.str, ".", "", -1)
			testcase.str = strings.Replace(testcase.str, ",", "", -1)
			countWordLen(testcase.str)
		}
	}
}

func TestCountWordLenCallByReference(t *testing.T) {
	for _, testcase := range countLenTests {
		i := []int{}
		t.Log(testcase.name)
		testcase.str = strings.Replace(testcase.str, ".", "", -1)
		testcase.str = strings.Replace(testcase.str, ",", "", -1)
		if result := countWordLenCallByReference(i, strings.Split(testcase.str, " ")); !reflect.DeepEqual(result, testcase.expect) {
			t.Errorf("result => %#v\n expect => %#v\n", result, testcase.expect)
		}
	}
}

func BenchmarkCountWordLenCallByReference(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		j := []int{}
		for _, testcase := range countLenTests {
			testcase.str = strings.Replace(testcase.str, ".", "", -1)
			testcase.str = strings.Replace(testcase.str, ",", "", -1)
			countWordLenCallByReference(j, strings.Split(testcase.str, " "))
		}
	}
}
