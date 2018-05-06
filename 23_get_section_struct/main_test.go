package main

import (
	"os"
	"reflect"
	"testing"
)

var selectSectionTests = []struct {
	name    string
	file    string
	text    string
	keyword string
	expect  Sections
}{
	{
		name: "select section structure",
		file: "./test.json",
		text: `{"text": "イギリスではないですが、イギリスとします。\n==== [[ラフィーク・ハリーリー|ハリーリー]]元首相暗殺事件まで ====\n=== 政治潮流と政党 ===", "title":"イギリス"}
{"text": "イギリスです。\n|titlebar=#ddd\n== シリア軍撤退後 ==\n[[Category:英連邦王国| ]]\n[[Category:海洋国家]]", "title":"イギリス"}`,
		keyword: "イギリス",
		expect: Sections{
			Section{"[[ラフィーク・ハリーリー|ハリーリー]]元首相暗殺事件まで", 3},
			Section{"政治潮流と政党", 2},
			Section{"シリア軍撤退後", 1},
		},
	},
}

func TestSelectSection(t *testing.T) {
	for _, testcase := range selectSectionTests {
		t.Log(testcase.name)

		f, err := os.Create(testcase.file)
		if err != nil {
			t.Errorf("could not create a file: %s\n  %s\n", testcase.file, err)
		}
		f.Write([]byte(testcase.text))
		f.Close()

		articles, err := readJSON(testcase.file)
		if err != nil {
			t.Error(err)
		}

		if result := articles.find(testcase.keyword).getSection(); !reflect.DeepEqual(result, testcase.expect) {
			t.Errorf("result => %#v\n shpould contain => %#v\n", result, testcase.expect)
		}

		if err := os.Remove(testcase.file); err != nil {
			t.Errorf("could not delete a file: %s\n  %s\n", testcase.file, err)
		}
	}
}
