package main

import (
	"reflect"
	"sort"
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

var elemSymbolTests = []struct {
	name   string
	str    string
	a      []int
	expect []string
}{
	{
		name:   "should return each symbol of an element with map array from 1 to 20",
		str:    "Hi He Lied Because Boron Could Not Oxidize Fluorine. New Nations Might Also Sign Peace Security Clause. Arthur King Can.",
		a:      []int{8, 7, 16, 15, 1, 9, 5, 6, 19},
		expect: []string{"H", "He", "Li", "Be", "B", "C", "N", "O", "F", "Ne", "Na", "Mi", "Al", "Si", "P", "S", "Cl", "Ar", "K", "Ca"},
	},
}

func TestGetElemSymbol(t *testing.T) {
	for _, testcase := range elemSymbolTests {
		t.Log(testcase.name)
		symbol := getElemSymbol(testcase.str, testcase.a)

		keys := []int{}
		for key := range symbol {
			keys = append(keys, key)
		}
		sort.Ints(keys)

		result := []string{}
		for _, idx := range keys {
			result = append(result, symbol[idx])
		}

		if !reflect.DeepEqual(result, testcase.expect) {
			t.Errorf("result => %#v\n expect => %#v\n", result, testcase.expect)
		}
	}
}

func BenchmarkGetElemSymbol(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for _, testcase := range elemSymbolTests {
			getElemSymbol(testcase.str, testcase.a)
		}
	}
}

func TestGetElemSymbolWithExtraWork(t *testing.T) {
	for _, testcase := range elemSymbolTests {
		t.Log(testcase.name)
		symbol := getElemSymbolWithExtraWork(testcase.str, testcase.a)

		keys := []int{}
		for key := range symbol {
			keys = append(keys, key)
		}
		sort.Ints(keys)

		result := []string{}
		for _, idx := range keys {
			result = append(result, symbol[idx])
		}

		if !reflect.DeepEqual(result, testcase.expect) {
			t.Errorf("result => %#v\n expect => %#v\n", result, testcase.expect)
		}
	}
}

func BenchmarkGetElemSymbolWithExtraWork(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for _, testcase := range elemSymbolTests {
			getElemSymbolWithExtraWork(testcase.str, testcase.a)
		}
	}
}

func TestGetElemSymbolBySet(t *testing.T) {
	for _, testcase := range elemSymbolTests {
		t.Log(testcase.name)
		symbol := getElemSymbolBySet(testcase.str, testcase.a)

		keys := []int{}
		for key := range symbol {
			keys = append(keys, key)
		}
		sort.Ints(keys)

		result := []string{}
		for _, idx := range keys {
			result = append(result, symbol[idx])
		}

		if !reflect.DeepEqual(result, testcase.expect) {
			t.Errorf("result => %#v\n expect => %#v\n", result, testcase.expect)
		}
	}
}

func BenchmarkGetElemSymbolBySet(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for _, testcase := range elemSymbolTests {
			getElemSymbolBySet(testcase.str, testcase.a)
		}
	}
}
