package main

import (
	"reflect"
	"strings"
	"testing"
)

var sortByAppearanceTests = []struct {
	name   string
	file   string
	text   string
	expect CountSorters
}{
	{
		name: "should return the surface verb stably sorted by appearance",
		file: "full-test.txt.mecab",
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
		expect: CountSorters{
			CountSorter{"南無阿弥陀仏", 2},
			CountSorter{"チー", 1},
			CountSorter{"ン", 1},
			CountSorter{"南無", 1},
			CountSorter{"猫", 1},
			CountSorter{"誉", 1},
			CountSorter{"信女", 1},
			CountSorter{"、", 1},
			CountSorter{"と", 1},
			CountSorter{"御", 1},
			CountSorter{"師匠", 1},
			CountSorter{"さん", 1},
			CountSorter{"の", 1},
			CountSorter{"声", 1},
			CountSorter{"が", 1},
			CountSorter{"する", 1},
			CountSorter{"。", 1},
		},
	},
	{
		name: "should return nothing from the text only containing \"EOS\"",
		file: "fail-text.txt.mecab",
		text: `EOS
EOS
EOS`,
		expect: CountSorters(nil),
	},
}

func TestSortByAppearanceTests(t *testing.T) {
	for _, testcase := range sortByAppearanceTests {
		t.Log(testcase.name)

		r := strings.NewReader(testcase.text)
		morphs, err := newMorpheme(r)
		if err != nil {
			t.Error(err)
		}

		if result := morphs.sortByAppearance(); !reflect.DeepEqual(result, testcase.expect) {
			t.Errorf("result => %#v\n shpould contain => %#v\n", result, testcase.expect)
		}
	}
}
