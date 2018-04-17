package main

import (
	"reflect"
	"testing"
)

var catStringTests = []struct {
	name   string
	str1   string
	str2   string
	expect string
}{
	{
		name:   "should return a string concatenating two strings",
		str1:   "acegikmoqsuwyacegikmoqsuwyacegikmoqsuwyacegikmoqsuwy",
		str2:   "bdfhjlnprtvxzbdfhjlnprtvxzbdfhjlnprtvxzbdfhjlnprtvxz",
		expect: "abcdefghijklmnopqrstuvwxyzabcdefghijklmnopqrstuvwxyzabcdefghijklmnopqrstuvwxyzabcdefghijklmnopqrstuvwxyz",
	},
	{
		name:   "should return a string concatenating two DBCS strings",
		str1:   "パトカー",
		str2:   "タクシー",
		expect: "パタトクカシーー",
	},
}

func TestInsertChar(t *testing.T) {
	for _, testcase := range catStringTests {
		t.Log(testcase.name)
		if str := insertChar(testcase.str1, testcase.str2); !reflect.DeepEqual(str, testcase.expect) {
			t.Errorf("result => %#v\n expect => %#v", str, testcase.expect)
		}
	}
}

func BenchmarkInsertChar(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for _, testcase := range catStringTests {
			insertChar(testcase.str1, testcase.str2)
		}
	}
}

func TestCatCharAltenately(t *testing.T) {
	for _, testcase := range catStringTests {
		t.Log(testcase.name)
		if str := catCharAltenately(testcase.str1, testcase.str2); !reflect.DeepEqual(str, testcase.expect) {
			t.Errorf("result => %#v\n expect => %#v", str, testcase.expect)
		}
	}
}

func BenchmarkCatCharAltenately(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for _, testcase := range catStringTests {
			catCharAltenately(testcase.str1, testcase.str2)
		}
	}
}

func TestPileUpLeadChar(t *testing.T) {
	for _, testcase := range catStringTests {
		t.Log(testcase.name)
		if str := pileUpLeadChar(testcase.str1, testcase.str2); !reflect.DeepEqual(str, testcase.expect) {
			t.Errorf("result => %#v\n expect => %#v", str, testcase.expect)
		}
	}
}

func BenchmarkPileUpLeadChar(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for _, testcase := range catStringTests {
			pileUpLeadChar(testcase.str1, testcase.str2)
		}
	}
}
