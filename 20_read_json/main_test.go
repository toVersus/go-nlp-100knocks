package main

import (
	"os"
	"reflect"
	"testing"
)

var readJSONTests = []struct {
	name    string
	file    string
	text    string
	keyword string
	expect  Articles
}{
	{
		name: "should return one body of text filtered by the keyword (\"イギリス\") matched in title",
		file: "./test.json",
		text: `{"text": "グレートブリテン及び北アイルランド連合王国（英: United Kingdom of Great Britain and Northern Ireland）。", "title":"イギリス"}
{"text": "アメリカ合衆国（英語: United States of America）、通称アメリカ、米国は、50の州および連邦区から成る連邦共和国である[6][7]。", "title":"アメリカ"}
{"text": "12345", "title": "number"}`,
		keyword: "イギリス",
		expect: Articles{
			Article{Text: "グレートブリテン及び北アイルランド連合王国（英: United Kingdom of Great Britain and Northern Ireland）。", Title: "イギリス"},
		},
	},
	{
		name: "should return two body of texts filtered by the keyword (\"イギリス\") matched in title and body respectively",
		file: "./test.json",
		text: `{"text": "グレートブリテン及び北アイルランド連合王国（英: United Kingdom of Great Britain and Northern Ireland）。", "title":"イギリス"}
{"text": "アメリカ合衆国（英語: United States of America）、通称アメリカ、米国は、50の州および連邦区から成る連邦共和国である[6][7]。", "title":"アメリカ"}
{"text": "北大西洋に位置するイギリスの島", "title": "グレートブリテン島"}`,
		keyword: "イギリス",
		expect: Articles{
			Article{Text: "グレートブリテン及び北アイルランド連合王国（英: United Kingdom of Great Britain and Northern Ireland）。", Title: "イギリス"},
		},
	},
	{
		name: "should return nothing filtered by the keyword (\"ドイツ\")",
		file: "./test.json",
		text: `{"text": "グレートブリテン及び北アイルランド連合王国（英: United Kingdom of Great Britain and Northern Ireland）。", "title":"イギリス"}
{"text": "アメリカ合衆国（英語: United States of America）、通称アメリカ、米国は、50の州および連邦区から成る連邦共和国である[6][7]。", "title":"アメリカ"}
{"text": "北大西洋に位置するイギリスの島", "title": "グレートブリテン島"}`,
		keyword: "ドイツ",
		expect:  nil,
	},
}

func TestFilterJSON(t *testing.T) {
	for _, testcase := range readJSONTests {
		t.Log(testcase.name)

		f, err := os.Create(testcase.file)
		if err != nil {
			t.Errorf("could not create a file: %s\n  %s\n", testcase.file, err)
		}
		f.WriteString(testcase.text)
		f.Close()

		if result := readJSON(testcase.file).find(testcase.keyword); !reflect.DeepEqual(result, testcase.expect) {
			t.Errorf("result => %#v\n should contain => %#v\n", result, testcase.expect)
		}

		if err := os.Remove(testcase.file); err != nil {
			t.Errorf("could not delete a file: %s\n  %s\n", testcase.file, err)
		}
	}
}
