package main

import (
	"reflect"
	"strings"
	"testing"
)

var getNounPhrasesTests = []struct {
	name   string
	text   string
	expect string
}{
	{
		name: "should get the noun phrase from the full text",
		text: `掌	名詞,一般,*,*,*,*,掌,テノヒラ,テノヒラ
の	助詞,連体化,*,*,*,*,の,ノ,ノ
上	名詞,非自立,副詞可能,*,*,*,上,ウエ,ウエ
で	助詞,格助詞,一般,*,*,*,で,デ,デ
少し	副詞,助詞類接続,*,*,*,*,少し,スコシ,スコシ
落ちつい	動詞,自立,*,*,五段・カ行イ音便,連用タ接続,落ちつく,オチツイ,オチツイ
て	助詞,接続助詞,*,*,*,*,て,テ,テ
書生	名詞,一般,*,*,*,*,書生,ショセイ,ショセイ
の	助詞,連体化,*,*,*,*,の,ノ,ノ
顔	名詞,一般,*,*,*,*,顔,カオ,カオ
を	助詞,格助詞,一般,*,*,*,を,ヲ,ヲ
見	動詞,自立,*,*,一段,連用形,見る,ミ,ミ
た	助動詞,*,*,*,特殊・タ,基本形,た,タ,タ
の	名詞,非自立,一般,*,*,*,の,ノ,ノ
が	助詞,格助詞,一般,*,*,*,が,ガ,ガ
いわゆる	連体詞,*,*,*,*,*,いわゆる,イワユル,イワユル
人間	名詞,一般,*,*,*,*,人間,ニンゲン,ニンゲン
という	助詞,格助詞,連語,*,*,*,という,トイウ,トユウ
もの	名詞,非自立,一般,*,*,*,もの,モノ,モノ
の	助詞,格助詞,一般,*,*,*,の,ノ,ノ
見	動詞,自立,*,*,一段,連用形,見る,ミ,ミ
始	名詞,固有名詞,人名,名,*,*,始,ハジメ,ハジメ
で	助動詞,*,*,*,特殊・ダ,連用形,だ,デ,デ
あろ	助動詞,*,*,*,五段・ラ行アル,未然ウ接続,ある,アロ,アロ
う	助動詞,*,*,*,不変化型,基本形,う,ウ,ウ
。	記号,句点,*,*,*,*,。,。,。
EOS
`,
		expect: "掌の上\n書生の顔",
	},
	{
		name: "should get the empty noun phrases because the file starts with the character of \"の\"",
		text: `の	助詞,連体化,*,*,*,*,の,ノ,ノ
上	名詞,非自立,副詞可能,*,*,*,上,ウエ,ウエ
で	助詞,格助詞,一般,*,*,*,で,デ,デ`,
		expect: "",
	},
	{
		name: "should get the empty noun phrases due to the input file containing only \"EOS\" words",
		text: `EOS
EOS
EOS`,
		expect: "",
	},
}

func TestGetNounPhrases(t *testing.T) {
	for _, testcase := range getNounPhrasesTests {
		t.Log(testcase.name)

		r := strings.NewReader(testcase.text)
		morphs, err := newMorpheme(r)
		if err != nil {
			t.Error(err)
		}
		if result := morphs.getNounPhrases(); !reflect.DeepEqual(result, testcase.expect) {
			t.Errorf("result => %#v\n shpould contain => %#v\n", result, testcase.expect)
		}
	}
}
