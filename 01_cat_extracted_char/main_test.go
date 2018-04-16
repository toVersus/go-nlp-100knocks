package main

import (
	"reflect"
	"testing"
)

var catCharTests = []struct {
	name   string
	str    string
	expect string
}{
	{
		name:   "should concatenate extracted char from DBCS string",
		str:    "パタトクカシーー",
		expect: "パトカー",
	},
}

func TestDelCharInSequence(t *testing.T) {
	for _, testcase := range catCharTests {
		t.Log(testcase.name)

		if result := delCharInSequence(testcase.str); !reflect.DeepEqual(result, testcase.expect) {
			t.Errorf("result => %s\n expect => %s\n", result, testcase.expect)
		}
	}
}

func BenchmarkDelCharInSequence(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for _, testcase := range catCharTests {
			delCharInSequence(testcase.str)
		}
	}
}

func TestCatCharInSequence(t *testing.T) {
	for _, testcase := range catCharTests {
		t.Log(testcase.name)

		if result := catCharInSequence(testcase.str); !reflect.DeepEqual(result, testcase.expect) {
			t.Errorf("result => %s\n expect => %s\n", result, testcase.expect)
		}
	}
}

func BenchmarkCatCharInSequence(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for _, testcase := range catCharTests {
			catCharInSequence(testcase.str)
		}
	}
}

var catSepCharAlternatelyTests = []struct {
	name   string
	str    string
	isOdd  bool
	expect string
}{
	{
		name:   "should return odd-numbered elements from odd-elements string input",
		str:    "oddstring",
		isOdd:  true,
		expect: "odtig",
	},
	{
		name:   "should return odd-numbered elements from even-elements string input",
		str:    "evenstring",
		isOdd:  true,
		expect: "eesrn",
	},
	{
		name:   "should return even-numbered elements from odd-elements string input",
		str:    "oddstring",
		isOdd:  false,
		expect: "dsrn",
	},
	{
		name:   "should return even-numbered elements from even-elements string input",
		str:    "evenstring",
		isOdd:  false,
		expect: "vntig",
	},
	{
		name:   "should return odd-numbered elements from single string input",
		str:    "o",
		isOdd:  true,
		expect: "o",
	},
	{
		name:   "should return nothing from single string input",
		str:    "e",
		isOdd:  false,
		expect: "",
	},
	{
		name:   "should return odd-numbered elements from DBCS string input",
		str:    "日本語だよー",
		isOdd:  true,
		expect: "日語よ",
	},
	{
		name:   "should return even-numbered elements from DBCS string input",
		str:    "日本語だよー",
		isOdd:  false,
		expect: "本だー",
	},
}

func TestCatSepCharInSequence(t *testing.T) {
	for _, testcase := range catSepCharAlternatelyTests {
		t.Log(testcase.name)
		if result := catSepCharInSequence(testcase.str, testcase.isOdd); !reflect.DeepEqual(result, testcase.expect) {
			t.Errorf("result => %s\n expect => %s\n", result, testcase.expect)
		}
	}
}
