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
		name: "should parse the full set of morphemes from *.mecab file",
		file: "fullset-test.txt.mecab",
		text: `　	記号,空白,*,*,*,*,　,　,　
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
		expect: morphemes{
			&morpheme{"surface": "生れ", "base": "生れる", "pos": "動詞", "pos1": "自立"},
			&morpheme{"surface": "つか", "base": "つく", "pos": "動詞", "pos1": "自立"},
		},
	},
	{
		name: "should return the empty morphemes from *mecab file",
		file: "fail-text.txt.mecab",
		text: `EOS
EOS
EOS`,
		expect: nil,
	},
}

func TestParseMorphemesTests(t *testing.T) {
	for _, testcase := range ParseMorphemesTests {
		t.Log(testcase.name)

		f, err := os.Create(testcase.file)
		if err != nil {
			t.Errorf("could not crearte a new file: %s\n  %s\n", testcase.file, err)
		}
		f.WriteString(testcase.text)
		f.Close()

		morphs, err := newMorpheme(testcase.file)
		if err != nil {
			t.Error(err)
		}
		if result := morphs.filterByPos("動詞"); !reflect.DeepEqual(result, testcase.expect) {
			t.Errorf("result => %#v\n shpould contain => %#v\n", result, testcase.expect)
		}

		if err := os.Remove(testcase.file); err != nil {
			t.Errorf("could not delete a file: %s\n  %s\n", testcase.file, err)
		}
	}
}
