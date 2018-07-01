package main

import (
	"reflect"
	"strings"
	"testing"
)

var findKeyTests = []struct {
	name    string
	text    string
	keyword string
	expect  Articles
}{
	{
		name: "should return one body of text filtered by the keyword (\"イギリス\") matched in title",
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

		r := strings.NewReader(testcase.text)
		filtered, err := readJSON(r)
		if err != nil {
			t.Error(err)
		}
		if result := filtered.find(testcase.keyword); !reflect.DeepEqual(result, testcase.expect) {
			t.Errorf("result => %#v\n should contain => %#v\n", result, testcase.expect)
		}
	}
}

var result Articles

func BenchmarkFind(b *testing.B) {
	var filtered Articles

	for i := 0; i < b.N; i++ {
		b.StopTimer()
		testcase := findKeyTests[0]
		r := strings.NewReader(testcase.text)
		b.StartTimer()

		a, _ := readJSON(r)
		filtered = a.find(testcase.keyword)
	}
	result = filtered
}
