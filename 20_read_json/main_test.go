package main

import (
	"os"
	"reflect"
	"testing"
)

var findKeyTests = []struct {
	name    string
	file    string
	text    string
	keyword string
	expect  Articles
}{
	{
		name: "should return one body of text filtered by the keyword (\"イギリス\") matched in title",
		file: "simple-test.json",
		text: `{"text": "グレートブリテン及び北アイルランド連合王国（英: United Kingdom of Great Britain and Northern Ireland）。", "title":"イギリス"}
{"text": "アメリカ合衆国（英語: United States of America）、通称アメリカ、米国は、50の州および連邦区から成る連邦共和国である[6][7]。", "title":"アメリカ"}
{"text": "12345", "title": "number"}`,
		keyword: "イギリス",
		expect: Articles{
			Article{Text: "グレートブリテン及び北アイルランド連合王国（英: United Kingdom of Great Britain and Northern Ireland）。", Title: "イギリス"},
		},
	},
	{
		name: "should return nothing filtered by the keyword (\"ドイツ\")",
		file: "empty-test.json",
		text: `{"text": "グレートブリテン及び北アイルランド連合王国（英: United Kingdom of Great Britain and Northern Ireland）。", "title":"イギリス"}
{"text": "アメリカ合衆国（英語: United States of America）、通称アメリカ、米国は、50の州および連邦区から成る連邦共和国である[6][7]。", "title":"アメリカ"}
{"text": "北大西洋に位置するイギリスの島", "title": "グレートブリテン島"}`,
		keyword: "ドイツ",
		expect:  nil,
	},
}

func TestFind(t *testing.T) {
	for _, testcase := range findKeyTests {
		t.Log(testcase.name)

		f, err := os.Create(testcase.file)
		if err != nil {
			t.Errorf("could not create a file: %s\n  %s\n", testcase.file, err)
		}
		f.WriteString(testcase.text)
		f.Close()

		filtered, err := readJSON(testcase.file)
		if err != nil {
			t.Error(err)
		}
		if result := filtered.find(testcase.keyword); !reflect.DeepEqual(result, testcase.expect) {
			t.Errorf("result => %#v\n should contain => %#v\n", result, testcase.expect)
		}

		if err := os.Remove(testcase.file); err != nil {
			t.Errorf("could not delete a file: %s\n  %s\n", testcase.file, err)
		}
	}
}

func BenchmarkFind(b *testing.B) {
	for _, testcase := range findKeyTests {
		f, err := os.Create(testcase.file)
		if err != nil {
			b.Errorf("could not create a file: %s\n  %s\n", testcase.file, err)
		}
		f.WriteString(testcase.text)
		f.Close()
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for _, testcase := range findKeyTests {
			filtered, _ := readJSON(testcase.file)
			filtered.find(testcase.keyword)
		}
	}
	b.StopTimer()

	for _, testcase := range findKeyTests {
		if err := os.Remove(testcase.file); err != nil {
			b.Errorf("could not delete a file: %s\n  %s\n", testcase.file, err)
		}
	}
}
