package main

import (
	"reflect"
	"strings"
	"testing"
)

var getLinkingNounPhrasesTests = []struct {
	name   string
	text   string
	expect string
}{
	{
		name: "should select the linking noun phrases from the full text including symbol such as \"、\" and \"。\"",
		text: `チー	名詞,固有名詞,人名,一般,*,*,チー,チー,チー
ン	名詞,非自立,一般,*,*,*,ン,ン,ン
南無	感動詞,*,*,*,*,*,南無,ナム,ナム
猫	名詞,一般,*,*,*,*,猫,ネコ,ネコ
誉	名詞,固有名詞,地域,一般,*,*,誉,ホマレ,ホマレ
信女	名詞,一般,*,*,*,*,信女,シンニョ,シンニョ
、	記号,読点,*,*,*,*,、,、,、
南無阿弥陀仏	名詞,一般,*,*,*,*,南無阿弥陀仏,ナムアミダブツ,ナムアミダブツ
南無阿弥陀仏	名詞,一般,*,*,*,*,南無阿弥陀仏,ナムアミダブツ,ナムアミダブツ
と	助詞,並立助詞,*,*,*,*,と,ト,ト
御	接頭詞,名詞接続,*,*,*,*,御,ゴ,ゴ
師匠	名詞,一般,*,*,*,*,師匠,シショウ,シショー
さん	名詞,接尾,人名,*,*,*,さん,サン,サン
の	助詞,連体化,*,*,*,*,の,ノ,ノ
声	名詞,一般,*,*,*,*,声,コエ,コエ
が	助詞,格助詞,一般,*,*,*,が,ガ,ガ
する	動詞,自立,*,*,サ変・スル,基本形,する,スル,スル
。	記号,句点,*,*,*,*,。,。,。
EOS
`,
		expect: "チーン\n猫誉信女\n南無阿弥陀仏南無阿弥陀仏\n師匠さん",
	},
	{
		name: "should select the linking noun phrases from the full text including lower case letter",
		text: `呼ん	動詞,自立,*,*,五段・バ行,連用タ接続,呼ぶ,ヨン,ヨン
で	助詞,接続助詞,*,*,*,*,で,デ,デ
御前	名詞,一般,*,*,*,*,御前,ゴゼン,ゴゼン
は	助詞,係助詞,*,*,*,*,は,ハ,ワ
女	名詞,一般,*,*,*,*,女,オンナ,オンナ
だ	助動詞,*,*,*,特殊・ダ,基本形,だ,ダ,ダ
けれども	助詞,接続助詞,*,*,*,*,けれども,ケレドモ,ケレドモ
many	名詞,一般,*,*,*,*,*
a	名詞,一般,*,*,*,*,*
slip	名詞,一般,*,*,*,*,*
'	名詞,サ変接続,*,*,*,*,*
twixt	名詞,一般,*,*,*,*,*
the	名詞,一般,*,*,*,*,*
cup	名詞,一般,*,*,*,*,*
and	名詞,一般,*,*,*,*,*
the	名詞,一般,*,*,*,*,*
lip	名詞,一般,*,*,*,*,*
EOS`,
		expect: "manyaslip'twixtthecupandthelip",
	},
	{
		name: "should select the linking noun phrases from the full text",
		text: `「	記号,括弧開,*,*,*,*,「,「,「
なるほど	感動詞,*,*,*,*,*,なるほど,ナルホド,ナルホド
面白い	形容詞,自立,*,*,形容詞・アウオ段,基本形,面白い,オモシロイ,オモシロイ
解釈	名詞,サ変接続,*,*,*,*,解釈,カイシャク,カイシャク
だ	助動詞,*,*,*,特殊・ダ,基本形,だ,ダ,ダ
」	記号,括弧閉,*,*,*,*,」,」,」
と	助詞,格助詞,引用,*,*,*,と,ト,ト
独	名詞,固有名詞,地域,国,*,*,独,ドク,ドク
仙	名詞,固有名詞,人名,名,*,*,仙,セン,セン
君	名詞,接尾,人名,*,*,*,君,クン,クン
が	助詞,格助詞,一般,*,*,*,が,ガ,ガ
云い	動詞,自立,*,*,五段・ワ行促音便,連用形,云う,イイ,イイ
出し	動詞,非自立,*,*,五段・サ行,連用形,出す,ダシ,ダシ
た	助動詞,*,*,*,特殊・タ,基本形,た,タ,タ
。	記号,句点,*,*,*,*,。,。,。
EOS`,
		expect: "独仙君",
	},
}

func TestGetLinkingNounPhrases(t *testing.T) {
	for _, testcase := range getLinkingNounPhrasesTests {
		t.Log(testcase.name)

		r := strings.NewReader(testcase.text)
		morphs, err := newMorpheme(r)
		if err != nil {
			t.Error(err)
		}
		if result := morphs.getLinkingNounPhrases(); !reflect.DeepEqual(result, testcase.expect) {
			t.Errorf("result => %#v\n shpould contain => %#v\n", result, testcase.expect)
		}
	}
}
