package main

import (
	"reflect"
	"strings"
	"testing"
)

var listDependedPhrasesTests = []struct {
	name   string
	file   string
	text   string
	expect string
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
		expect: "猫\tで\tある\n吾輩\tは\t猫\tで\tある\n名前\tは\t無い\nまだ\t無い",
	},
	{
		name: "should return nothing from the text only containing \"EOS\"",
		file: "fail-text.txt.mecab",
		text: `EOS
EOS
EOS`,
		expect: "",
	},
	{
		name: "should return the sorted surface from the full text",
		file: "full-test.txt.mecab",
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
		expect: "そこ\tを\t我慢\tし\tて\n我慢\tし\tて\t這っ\tて\t行く\tと\n無理やり\tに\t這っ\tて\t行く\tと\n這っ\tて\t行く\tと\t出\tた\nようやく\tの\t事\tで\n事\tで\t出\tた\n何となく\t人間\t臭い\n人間\t臭い\t所\tへ\n所\tへ\t出\tた",
	},
}

func TestExtractDependedPhrases(t *testing.T) {
	for _, testcase := range listDependedPhrasesTests {
		t.Log(testcase.name)

		r := strings.NewReader(testcase.text)
		if result := newChunkPassage(r).listDependedPhrases(); !reflect.DeepEqual(result, testcase.expect) {
			t.Errorf("result => %#v\n should contain => %#v\n", result, testcase.expect)
		}
	}
}

var result string

func BenchmarkListDependedPhrases(b *testing.B) {
	var s string
	for i := 0; i < b.N; i++ {
		for _, testcase := range listDependedPhrasesTests {
			r := strings.NewReader(testcase.text)
			s = newChunkPassage(r).listDependedPhrases()
		}
	}
	result = s
}
