package main

import (
	"reflect"
	"strings"
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

		r := strings.NewReader(testcase.text)
		result, err := newMorpheme(r)
		if err != nil {
			t.Error(err)
		}
		if !reflect.DeepEqual(result, testcase.expect) {
			t.Errorf("result => %#v\n should contain => %#v\n", result, testcase.expect)
		}
	}
}
