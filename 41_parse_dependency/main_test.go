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
	expect *ChunkPassage
}{
	{
		name: "should return the parsed chunk of passages from the result of morpheme/dependency analysis",
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
		expect: &ChunkPassage{
			Passage{
				Chunk{
					[]*Morph{
						&Morph{surface: "一", base: "一", pos: "名詞", pos1: "数"},
					},
					-1,
					nil,
				},
			},
			Passage{
				Chunk{
					[]*Morph{
						&Morph{surface: "　", base: "　", pos: "記号", pos1: "空白"},
					},
					2,
					nil,
				},
				Chunk{
					[]*Morph{
						&Morph{surface: "吾輩", base: "吾輩", pos: "名詞", pos1: "代名詞"},
						&Morph{surface: "は", base: "は", pos: "助詞", pos1: "係助詞"},
					},
					2,
					nil,
				},
				Chunk{
					[]*Morph{
						&Morph{surface: "猫", base: "猫", pos: "名詞", pos1: "一般"},
						&Morph{surface: "で", base: "だ", pos: "助動詞", pos1: "*"},
						&Morph{surface: "ある", base: "ある", pos: "助動詞", pos1: "*"},
						&Morph{surface: "。", base: "。", pos: "記号", pos1: "句点"},
					},
					-1,
					[]int{0, 1},
				},
			},
			Passage{
				Chunk{
					[]*Morph{
						&Morph{surface: "名前", base: "名前", pos: "名詞", pos1: "一般"},
						&Morph{surface: "は", base: "は", pos: "助詞", pos1: "係助詞"},
					},
					2,
					nil,
				},
				Chunk{
					[]*Morph{
						&Morph{surface: "まだ", base: "まだ", pos: "副詞", pos1: "助詞類接続"},
					},
					2,
					nil,
				},
				Chunk{
					[]*Morph{
						&Morph{surface: "無い", base: "無い", pos: "形容詞", pos1: "自立"},
						&Morph{surface: "。", base: "。", pos: "記号", pos1: "句点"},
					},
					-1,
					[]int{0, 1},
				},
			},
		},
	},
	{
		name: "should return the fully parsed single passage from the result of morpheme/dependency analysis",
		file: "full-one-passage-test.txt.mecab",
		text: `
EOS
* 0 1D 0/1 1.816431
そこ	名詞,代名詞,一般,*,*,*,そこ,ソコ,ソコ
を	助詞,格助詞,一般,*,*,*,を,ヲ,ヲ
* 1 3D 1/2 0.538467
我慢	名詞,サ変接続,*,*,*,*,我慢,ガマン,ガマン
し	動詞,自立,*,*,サ変・スル,連用形,する,シ,シ
て	助詞,接続助詞,*,*,*,*,て,テ,テ
* 2 3D 0/1 2.021870
無理やり	名詞,一般,*,*,*,*,無理やり,ムリヤリ,ムリヤリ
に	助詞,格助詞,一般,*,*,*,に,ニ,ニ
* 3 9D 0/3 -1.346861
這っ	動詞,自立,*,*,五段・ワ行促音便,連用タ接続,這う,ハッ,ハッ
て	助詞,接続助詞,*,*,*,*,て,テ,テ
行く	動詞,非自立,*,*,五段・カ行促音便,基本形,行く,イク,イク
と	助詞,接続助詞,*,*,*,*,と,ト,ト
* 4 5D 0/1 1.516662
ようやく	副詞,一般,*,*,*,*,ようやく,ヨウヤク,ヨーヤク
の	助詞,連体化,*,*,*,*,の,ノ,ノ
* 5 9D 0/1 -1.346861
事	名詞,非自立,一般,*,*,*,事,コト,コト
で	助詞,格助詞,一般,*,*,*,で,デ,デ
* 6 7D 0/0 0.252264
何となく	副詞,一般,*,*,*,*,何となく,ナントナク,ナントナク
* 7 8D 1/1 1.736467
人間	名詞,一般,*,*,*,*,人間,ニンゲン,ニンゲン
臭い	形容詞,自立,*,*,形容詞・アウオ段,基本形,臭い,クサイ,クサイ
* 8 9D 0/1 -1.346861
所	名詞,非自立,副詞可能,*,*,*,所,トコロ,トコロ
へ	助詞,格助詞,一般,*,*,*,へ,ヘ,エ
* 9 -1D 0/1 0.000000
出	動詞,自立,*,*,一段,連用形,出る,デ,デ
た	助動詞,*,*,*,特殊・タ,基本形,た,タ,タ
。	記号,句点,*,*,*,*,。,。,。
EOS

`,
		expect: &ChunkPassage{
			Passage{
				Chunk{
					[]*Morph{
						&Morph{surface: "そこ", base: "そこ", pos: "名詞", pos1: "代名詞"},
						&Morph{surface: "を", base: "を", pos: "助詞", pos1: "格助詞"},
					},
					1,
					[]int(nil),
				},
				Chunk{
					[]*Morph{
						&Morph{surface: "我慢", base: "我慢", pos: "名詞", pos1: "サ変接続"},
						&Morph{surface: "し", base: "する", pos: "動詞", pos1: "自立"},
						&Morph{surface: "て", base: "て", pos: "助詞", pos1: "接続助詞"},
					},
					3,
					[]int{0},
				},
				Chunk{
					[]*Morph{
						&Morph{surface: "無理やり", base: "無理やり", pos: "名詞", pos1: "一般"},
						&Morph{surface: "に", base: "に", pos: "助詞", pos1: "格助詞"},
					},
					3,
					[]int(nil),
				},
				Chunk{
					[]*Morph{
						&Morph{surface: "這っ", base: "這う", pos: "動詞", pos1: "自立"},
						&Morph{surface: "て", base: "て", pos: "助詞", pos1: "接続助詞"},
						&Morph{surface: "行く", base: "行く", pos: "動詞", pos1: "非自立"},
						&Morph{surface: "と", base: "と", pos: "助詞", pos1: "接続助詞"},
					},
					9,
					[]int{1, 2},
				},
				Chunk{
					[]*Morph{
						&Morph{surface: "ようやく", base: "ようやく", pos: "副詞", pos1: "一般"},
						&Morph{surface: "の", base: "の", pos: "助詞", pos1: "連体化"},
					},
					5,
					[]int(nil),
				},
				Chunk{
					[]*Morph{
						&Morph{surface: "事", base: "事", pos: "名詞", pos1: "非自立"},
						&Morph{surface: "で", base: "で", pos: "助詞", pos1: "格助詞"},
					},
					9,
					[]int{4},
				},
				Chunk{
					[]*Morph{
						&Morph{surface: "何となく", base: "何となく", pos: "副詞", pos1: "一般"},
					},
					7,
					[]int(nil),
				},
				Chunk{
					[]*Morph{
						&Morph{surface: "人間", base: "人間", pos: "名詞", pos1: "一般"},
						&Morph{surface: "臭い", base: "臭い", pos: "形容詞", pos1: "自立"},
					},
					8,
					[]int{6},
				},
				Chunk{
					[]*Morph{
						&Morph{surface: "所", base: "所", pos: "名詞", pos1: "非自立"},
						&Morph{surface: "へ", base: "へ", pos: "助詞", pos1: "格助詞"},
					},
					9,
					[]int{7},
				},
				Chunk{
					[]*Morph{
						&Morph{surface: "出", base: "出る", pos: "動詞", pos1: "自立"},
						&Morph{surface: "た", base: "た", pos: "助動詞", pos1: "*"},
						&Morph{surface: "。", base: "。", pos: "記号", pos1: "句点"},
					},
					-1,
					[]int{3, 5, 8},
				},
			},
		},
	},
}

func TestParseMorph(t *testing.T) {
	for _, testcase := range parseMorphTests {
		t.Log(testcase.name)

		r := strings.NewReader(testcase.text)
		if result := newChunkPassage(r); !reflect.DeepEqual(result, testcase.expect) {
			t.Errorf("result => %#v\n should contain => %#v\n", result, testcase.expect)
		}
	}
}
