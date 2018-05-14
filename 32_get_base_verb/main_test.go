package main

import (
	"reflect"
	"strings"
	"testing"
)

var ParseMorphemesTests = []struct {
	name   string
	text   string
	field  string
	expect string
}{
	{
		name: "should return bare verbs",
		text: `一	名詞,数,*,*,*,*,一,イチ,イチ
EOS
EOS
　	記号,空白,*,*,*,*,　,　,　
吾輩	名詞,代名詞,一般,*,*,*,吾輩,ワガハイ,ワガハイ
は	助詞,係助詞,*,*,*,*,は,ハ,ワ
猫	名詞,一般,*,*,*,*,猫,ネコ,ネコ
で	助動詞,*,*,*,特殊・ダ,連用形,だ,デ,デ
ある	助動詞,*,*,*,五段・ラ行アル,基本形,ある,アル,アル
。	記号,句点,*,*,*,*,。,。,。
EOS
名前	名詞,一般,*,*,*,*,名前,ナマエ,ナマエ
は	助詞,係助詞,*,*,*,*,は,ハ,ワ
まだ	副詞,助詞類接続,*,*,*,*,まだ,マダ,マダ
無い	形容詞,自立,*,*,形容詞・アウオ段,基本形,無い,ナイ,ナイ
。	記号,句点,*,*,*,*,。,。,。
EOS
EOS
　	記号,空白,*,*,*,*,　,　,　
どこ	名詞,代名詞,一般,*,*,*,どこ,ドコ,ドコ
で	助詞,格助詞,一般,*,*,*,で,デ,デ
生れ	動詞,自立,*,*,一段,連用形,生れる,ウマレ,ウマレ
た	助動詞,*,*,*,特殊・タ,基本形,た,タ,タ
か	助詞,副助詞／並立助詞／終助詞,*,*,*,*,か,カ,カ
とんと	副詞,一般,*,*,*,*,とんと,トント,トント
見当	名詞,サ変接続,*,*,*,*,見当,ケントウ,ケントー
が	助詞,格助詞,一般,*,*,*,が,ガ,ガ
つか	動詞,自立,*,*,五段・カ行イ音便,未然形,つく,ツカ,ツカ
ぬ	助動詞,*,*,*,特殊・ヌ,基本形,ぬ,ヌ,ヌ
。	記号,句点,*,*,*,*,。,。,。
EOS
`,
		field:  "base",
		expect: "生れる\nつく",
	},
	{
		name: "should not return verbs",
		text: `EOS
EOS
EOS`,
		field:  "surface",
		expect: "",
	},
}

func TestParseMorphemesTests(t *testing.T) {
	for _, testcase := range ParseMorphemesTests {
		t.Log(testcase.name)

		r := strings.NewReader(testcase.text)
		morphs, err := newMorpheme(r)
		if err != nil {
			t.Error(err)
		}
		if result := morphs.filterByPos("動詞").stringify(testcase.field); !reflect.DeepEqual(result, testcase.expect) {
			t.Errorf("result => %#v\n shpould contain => %#v\n", result, testcase.expect)
		}
	}
}
