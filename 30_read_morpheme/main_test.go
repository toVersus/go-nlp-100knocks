package main

import (
	"os"
	"reflect"
	"testing"
)

var ParseMorphemesTests = []struct {
	name   string
	file   string
	text   string
	expect morphemes
}{
	{
		name: "should parse the whitespace from *.mecab file",
		file: "whitespace-test.txt.mecab",
		text: `　	記号,空白,*,*,*,*,　,　,　`,
		expect: morphemes{
			morpheme{"surface": "　", "base": "　", "pos": "記号", "pos1": "空白"},
		},
	},
	{
		name: "should parse the full set of morphemes from *.mecab file",
		file: "fullset-test.txt.mecab",
		text: `吾輩	名詞,代名詞,一般,*,*,*,吾輩,ワガハイ,ワガハイ
は	助詞,係助詞,*,*,*,*,は,ハ,ワ
猫	名詞,一般,*,*,*,*,猫,ネコ,ネコ
で	助動詞,*,*,*,特殊・ダ,連用形,だ,デ,デ
ある	助動詞,*,*,*,五段・ラ行アル,基本形,ある,アル,アル
。	記号,句点,*,*,*,*,。,。,。
EOS
`,
		expect: morphemes{
			morpheme{"surface": "吾輩", "base": "吾輩", "pos": "名詞", "pos1": "代名詞"},
			morpheme{"surface": "は", "base": "は", "pos": "助詞", "pos1": "係助詞"},
			morpheme{"surface": "猫", "base": "猫", "pos": "名詞", "pos1": "一般"},
			morpheme{"surface": "で", "base": "だ", "pos": "助動詞", "pos1": "*"},
			morpheme{"surface": "ある", "base": "ある", "pos": "助動詞", "pos1": "*"},
			morpheme{"surface": "。", "base": "。", "pos": "記号", "pos1": "句点"},
		},
	},
	{
		name: "should return the empty morphemes from *mecab file",
		file: "fail-text.txt.mecab",
		text: `EOS
EOS
EOS`,
		expect: morphemes{},
	},
}

func TestParseMorphemesTests(t *testing.T) {
	for _, testcase := range ParseMorphemesTests {
		t.Log(testcase.name)

		f, err := os.Create(testcase.file)
		if err != nil {
			t.Errorf("could not crearte a new file: %s\n  %s\n", testcase.file, err)
		}
		f.Write([]byte(testcase.text))
		f.Close()

		results, err := newMorpheme(testcase.file)
		if err != nil {
			t.Error(err)
		}
		if !reflect.DeepEqual(results, testcase.expect) {
			t.Errorf("result => %#v\n shpould contain => %#v\n", results, testcase.expect)
		}

		if err := os.Remove(testcase.file); err != nil {
			t.Errorf("could not delete a file: %s\n  %s\n", testcase.file, err)
		}
	}
}
