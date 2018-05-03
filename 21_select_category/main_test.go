package main

import (
	"os"
	"reflect"
	"testing"
)

var selectCategoryTests = []struct {
	name    string
	file    string
	text    string
	keyword string
	expect  []string
}{
	{
		name: "should return 5 categolized lines",
		file: "simple-test.txt",
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
		file:    "empty-test.txt",
		text:    `{"text": "インドです。\nイギリスに関係します。\n[[Category:Category:]]\n[[Category: [[Sub Category]]]]\n[[Category:インド|*]]\n[[Category:イギリス連邦]][[foobar]]\n[[Category:連邦]]]]", "title":"インド"}`,
		keyword: "イギリス",
		expect:  nil,
	},
}

func TestSelectCategory(t *testing.T) {
	for _, testcase := range selectCategoryTests {
		t.Log(testcase.name)

		f, err := os.Create(testcase.file)
		if err != nil {
			t.Errorf("could not create a file: %s\n  %s\n", testcase.file, err)
		}
		f.WriteString(testcase.text)
		f.Close()

		artists, err := readJSON(testcase.file)
		if err != nil {
			t.Error(err)
		}

		if results := artists.find(testcase.keyword).selectCategory(); !reflect.DeepEqual(results, testcase.expect) {
			t.Errorf("result => %#v\n should contain => %#v\n", results, testcase.expect)
		}

		if err := os.Remove(testcase.file); err != nil {
			t.Errorf("could not delete a file: %s\n  %s\n", testcase.file, err)
		}
	}
}

func BenchmarkSelectCategory(b *testing.B) {
	for _, testcase := range selectCategoryTests {
		f, err := os.Create(testcase.file)
		if err != nil {
			b.Errorf("could not create a file: %s\n  %s\n", testcase.file, err)
		}
		f.WriteString(testcase.text)
		f.Close()
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for _, testcase := range selectCategoryTests {
			filtered, _ := readJSON(testcase.file)
			filtered.find(testcase.keyword).selectCategory()
		}
	}
	b.StopTimer()

	for _, testcase := range selectCategoryTests {
		if err := os.Remove(testcase.file); err != nil {
			b.Errorf("could not delete a file: %s\n  %s\n", testcase.file, err)
		}
	}
}
