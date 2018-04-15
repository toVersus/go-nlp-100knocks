package main

import (
	"reflect"
	"testing"
)

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

func TestCatSepCharAlternately(t *testing.T) {
	for _, testcase := range catSepCharAlternatelyTests {
		t.Log(testcase.name)
		if result := catSepCharAlternately(testcase.str, testcase.isOdd); !reflect.DeepEqual(result, testcase.expect) {
			t.Errorf("result => %s\n expect => %s\n", result, testcase.expect)
		}
	}
}
