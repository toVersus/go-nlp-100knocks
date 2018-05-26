package main

import (
	"reflect"
	"strings"
	"testing"
)

var digraphTests = []struct {
	name           string
	text           string
	sentenceNum    int
	nodesExpect    []string
	srcEdgesExpect []string
	dstEdgesExpect []string
}{
	{
		name: "should return the sorted surface from the full text",
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
		sentenceNum: 2,
		nodesExpect: []string{
			"吾輩は",
			"猫である",
		},
		srcEdgesExpect: []string{
			"吾輩は",
		},
		dstEdgesExpect: []string{
			"猫である",
		},
	},
	{
		name: "should return the sorted surface from the full text",
		text: `EOS
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
		sentenceNum: 1,
		nodesExpect: []string{
			"そこを",
			"我慢して",
			"無理やりに",
			"這って行くと",
			"ようやくの",
			"事で",
			"何となく",
			"人間臭い",
			"所へ",
			"出た",
		},
		srcEdgesExpect: []string{
			"そこを",
			"我慢して",
			"無理やりに",
			"這って行くと",
			"ようやくの",
			"事で",
			"何となく",
			"人間臭い",
			"所へ",
		},
		dstEdgesExpect: []string{
			"我慢して",
			"這って行くと",
			"這って行くと",
			"出た",
			"事で",
			"出た",
			"人間臭い",
			"所へ",
			"出た",
		},
	},
}

func TestExtractDependedPhrases(t *testing.T) {
	for _, testcase := range digraphTests {
		t.Log(testcase.name)

		r := strings.NewReader(testcase.text)
		chunkPassage := newChunkPassage(r)

		var nodesResults, srcEdgesResults, dstEdgesResults []string
		for _, chunk := range (*chunkPassage)[testcase.sentenceNum-1] {
			var srcWords, dstWords []string
			var srcStr, dstStr string
			for _, src := range chunk.morphems {
				if src.pos == "記号" {
					continue
				}
				srcWords = append(srcWords, src.surface)
			}

			srcStr = strings.Join(srcWords, "")
			if srcStr != "" {
				nodesResults = append(nodesResults, srcStr)
			}

			if chunk.dst != -1 {
				for _, dst := range (*chunkPassage)[testcase.sentenceNum-1][chunk.dst].morphems {
					if dst.pos == "記号" {
						continue
					}
					dstWords = append(dstWords, dst.surface)
				}
				dstStr = strings.Join(dstWords, "")

				if srcStr != "" && dstStr != "" {
					srcEdgesResults = append(srcEdgesResults, srcStr)
					dstEdgesResults = append(dstEdgesResults, dstStr)
				}
			}
		}

		if !reflect.DeepEqual(nodesResults, testcase.nodesExpect) {
			t.Errorf("node result => %#v\n  should contain => %#v\n", nodesResults, testcase.nodesExpect)
		}

		if !reflect.DeepEqual(srcEdgesResults, testcase.srcEdgesExpect) {
			t.Errorf("src edge result => %#v\n  should contain => %#v\n", srcEdgesResults, testcase.srcEdgesExpect)
		}

		if !reflect.DeepEqual(dstEdgesResults, testcase.dstEdgesExpect) {
			t.Errorf("dst edge result => %#v\n  should contain => %#v\n", dstEdgesResults, testcase.dstEdgesExpect)
		}
	}
}
