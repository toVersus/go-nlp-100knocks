package main

import (
	"reflect"
	"strings"
	"testing"
)

var parseMorphTests = []struct {
	name   string
	file   string
	text   string
	expect Passage
}{
	{
		name: "should return the sorted surface from the full text",
		file: "full-test.txt.mecab",
		text: `* 0 -1D 0/0 0.000000
一	名詞,数,*,*,*,*,一,イチ,イチ
EOS
EOS
* 0 2D 0/0 -0.764522
　	記号,空白,*,*,*,*,　,　,　
* 1 2D 0/1 -0.764522
吾輩	名詞,代名詞,一般,*,*,*,吾輩,ワガハイ,ワガハイ
は	助詞,係助詞,*,*,*,*,は,ハ,ワ
* 2 -1D 0/2 0.000000
猫	名詞,一般,*,*,*,*,猫,ネコ,ネコ
で	助動詞,*,*,*,特殊・ダ,連用形,だ,デ,デ
ある	助動詞,*,*,*,五段・ラ行アル,基本形,ある,アル,アル
。	記号,句点,*,*,*,*,。,。,。
EOS
* 0 2D 0/1 -1.911675
名前	名詞,一般,*,*,*,*,名前,ナマエ,ナマエ
は	助詞,係助詞,*,*,*,*,は,ハ,ワ
* 1 2D 0/0 -1.911675
まだ	副詞,助詞類接続,*,*,*,*,まだ,マダ,マダ
* 2 -1D 0/0 0.000000
無い	形容詞,自立,*,*,形容詞・アウオ段,基本形,無い,ナイ,ナイ
。	記号,句点,*,*,*,*,。,。,。
EOS
EOS
`,
		expect: Passage{
			Chunk{
				morphems: []Morph{
					Morph{surface: "一", base: "一", pos: "名詞", pos1: "数"},
				},
			},
			Chunk{
				morphems: []Morph{
					Morph{surface: "　", base: "　", pos: "記号", pos1: "空白"},
					Morph{surface: "吾輩", base: "吾輩", pos: "名詞", pos1: "代名詞"},
					Morph{surface: "は", base: "は", pos: "助詞", pos1: "係助詞"},
					Morph{surface: "猫", base: "猫", pos: "名詞", pos1: "一般"},
					Morph{surface: "で", base: "だ", pos: "助動詞", pos1: "*"},
					Morph{surface: "ある", base: "ある", pos: "助動詞", pos1: "*"},
					Morph{surface: "。", base: "。", pos: "記号", pos1: "句点"},
				},
			},
			Chunk{
				morphems: []Morph{
					Morph{surface: "名前", base: "名前", pos: "名詞", pos1: "一般"},
					Morph{surface: "は", base: "は", pos: "助詞", pos1: "係助詞"},
					Morph{surface: "まだ", base: "まだ", pos: "副詞", pos1: "助詞類接続"},
					Morph{surface: "無い", base: "無い", pos: "形容詞", pos1: "自立"},
					Morph{surface: "。", base: "。", pos: "記号", pos1: "句点"},
				},
			},
		},
	},
	{
		name: "should return nothing from the text only containing \"EOS\"",
		file: "fail-text.txt.mecab",
		text: `EOS
EOS
EOS`,
		expect: Passage{},
	},
}

func TestParseMorph(t *testing.T) {
	for _, testcase := range parseMorphTests {
		t.Log(testcase.name)

		r := strings.NewReader(testcase.text)
		if result := newMorpheme(r); !reflect.DeepEqual(result, testcase.expect) {
			t.Errorf("result => %#v\n should contain => %#v\n", result, testcase.expect)
		}
	}
}
