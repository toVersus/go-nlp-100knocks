package main

import (
	"reflect"
	"testing"
)

var isInArrayTests = []struct {
	name   string
	a      []int
	b      int
	expect bool
}{
	{
		name:   "should return true when int array contains the specified int value",
		a:      []int{1, 3, 5, 7, 9},
		b:      1,
		expect: true,
	},
	{
		name:   "should NOT return true when int array doesn't contain the specified int value",
		a:      []int{1, 3, 5, 7, 9},
		b:      2,
		expect: false,
	},
	{
		name:   "should NOT return true when int array doesn't contain the specified int value",
		a:      []int{5, 3, 1, 4, 9},
		b:      2,
		expect: false,
	},
}

func TestIsInArray(t *testing.T) {
	for _, testcase := range isInArrayTests {
		t.Log(testcase.name)
		if result := isInArray(testcase.a, testcase.b); !reflect.DeepEqual(result, testcase.expect) {
			t.Errorf("result => %#v\n expect => %#v\n", result, testcase.expect)
		}
	}
}

var listElementTests = []struct {
	name   string
	str    string
	a      []int
	expect map[int]string
}{
	{
		name: "should return each symbol of an element with map array from 1 to 20",
		str:  "Hi He Lied Because Boron Could Not Oxidize Fluorine. New Nations Might Also Sign Peace Security Clause. Arthur King Can.",
		a:    []int{8, 7, 16, 15, 1, 9, 5, 6, 19},
		expect: map[int]string{
			1:  "H",
			2:  "He",
			3:  "Li",
			4:  "Be",
			5:  "B",
			6:  "C",
			7:  "N",
			8:  "O",
			9:  "F",
			10: "Ne",
			11: "Na",
			12: "Mi",
			13: "Al",
			14: "Si",
			15: "P",
			16: "S",
			17: "Cl",
			18: "Ar",
			19: "K",
			20: "Ca",
		},
	},
}

func TestGetElemSymbol(t *testing.T) {
	for _, testcase := range listElementTests {
		t.Log(testcase.name)
		result := getElemSymbol(testcase.str, testcase.a)
		for i := range result {
			if !reflect.DeepEqual(result[i+1], testcase.expect[i+1]) {
				t.Errorf("result[%d] => %#v\n expect[%d] => %#v\n", i+1, result[i+1], i+1, testcase.expect[i+1])
			}
		}
	}
}

func BenchmarkGetElemSymbol(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for _, testcase := range listElementTests {
			getElemSymbol(testcase.str, testcase.a)
		}
	}
}

func BenchmarkListElement3(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for _, testcase := range listElementTests {
			getElemSymbolBySet(testcase.str, testcase.a)
		}
	}
}
