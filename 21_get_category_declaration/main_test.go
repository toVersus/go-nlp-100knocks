package main

import (
	"reflect"
	"strings"
	"testing"
)

var selectCategoryTests = []struct {
	name    string
	text    string
	keyword string
	expect  []string
}{
	{
		name: "should return 3 categolized lines",
		text: `{"text": "インドです。\nイギリスに関係します。\n[[Category:インド|*]]\n[[Category:イギリス連邦]]", "title":"インド"}
{"text": "イギリスです。\n[[Category:イギリス|*]]\n[[Category:英連邦王国|*]]\n[[Category:海洋国家]]", "title":"イギリス"}`,
		keyword: "イギリス",
		expect: []string{
			"[[Category:イギリス|*]]",
			"[[Category:英連邦王国|*]]",
			"[[Category:海洋国家]]",
		},
	},
	{
		name:    "should return empty categories",
		text:    `{"text": "インドです。\nイギリスに関係します。\n[[Category:Category:]]\n[[Category: [[Sub Category]]]]\n[[Category:インド|*]]\n[[Category:イギリス連邦]][[foobar]]\n[[Category:連邦]]]]", "title":"インド"}`,
		keyword: "イギリス",
		expect:  nil,
	},
}

func TestSelectCategory(t *testing.T) {
	for _, testcase := range selectCategoryTests {
		t.Log(testcase.name)

		r := strings.NewReader(testcase.text)
		filtered, err := readJSON(r)
		if err != nil {
			t.Error(err)
		}
		if results := filtered.find(testcase.keyword).getCategory(); !reflect.DeepEqual(results, testcase.expect) {
			t.Errorf("result => %#v\n should contain => %#v\n", results, testcase.expect)
		}
	}
}

var result []string

func BenchmarkSelectCategory(b *testing.B) {
	var categories []string
	for i := 0; i < b.N; i++ {
		for _, testcase := range selectCategoryTests {
			b.StopTimer()
			r := strings.NewReader(testcase.text)
			filtered, _ := readJSON(r)
			b.StartTimer()
			filtered.find(testcase.keyword).getCategory()
		}
	}
	result = categories
}
