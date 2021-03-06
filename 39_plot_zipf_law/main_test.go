package main

import (
	"reflect"
	"strings"
	"testing"
)

var sortByAppearanceTests = []struct {
	name   string
	text   string
	chart  string
	expect string
}{
	{
		name: "should return the sorted surface from the full text",
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
		expect: `Key:南無阿弥陀仏 Count:2
Key:チー Count:1
Key:ン Count:1
Key:南無 Count:1
Key:猫 Count:1
Key:誉 Count:1
Key:信女 Count:1
Key:、 Count:1
Key:と Count:1
Key:御 Count:1
Key:師匠 Count:1
Key:さん Count:1
Key:の Count:1
Key:声 Count:1
Key:が Count:1
Key:する Count:1
Key:。 Count:1`,
		chart: "top10-commonly-used-words-test.png",
	},
	{
		name: "should return nothing from the text only containing \"EOS\"",
		text: `EOS
EOS
EOS`,
		expect: "",
	},
}

func TestSortByCounts(t *testing.T) {
	for _, testcase := range sortByAppearanceTests {
		t.Log(testcase.name)

		r := strings.NewReader(testcase.text)
		morphs, err := newMorpheme(r)
		if err != nil {
			t.Error(err)
		}

		sorted := morphs.sortByAppearance()

		if result := sorted.String(); !reflect.DeepEqual(result, testcase.expect) {
			t.Errorf("result => %#v\n shpould contain => %#v\n", result, testcase.expect)
		}
	}
}

var rankByCountsTests = []struct {
	name   string
	file   string
	text   string
	expect CountRanking
}{
	{
		name: "should return ranking of close context",
		file: "close-context-ranking.txt.mecab",
		text: `チー	名詞,固有名詞,人名,一般,*,*,チー,チー,チー
南無阿弥陀仏	名詞,一般,*,*,*,*,南無阿弥陀仏,ナムアミダブツ,ナムアミダブツ
南無阿弥陀仏	名詞,一般,*,*,*,*,南無阿弥陀仏,ナムアミダブツ,ナムアミダブツ
ン	名詞,非自立,一般,*,*,*,ン,ン,ン
南無	感動詞,*,*,*,*,*,南無,ナム,ナム
の	助詞,連体化,*,*,*,*,の,ノ,ノ
の	助詞,連体化,*,*,*,*,の,ノ,ノ
ン	名詞,非自立,一般,*,*,*,ン,ン,ン
猫	名詞,一般,*,*,*,*,猫,ネコ,ネコ
猫	名詞,一般,*,*,*,*,猫,ネコ,ネコ
チー	名詞,固有名詞,人名,一般,*,*,チー,チー,チー
南無阿弥陀仏	名詞,一般,*,*,*,*,南無阿弥陀仏,ナムアミダブツ,ナムアミダブツ
誉	名詞,固有名詞,地域,一般,*,*,誉,ホマレ,ホマレ
チー	名詞,固有名詞,人名,一般,*,*,チー,チー,チー
信女	名詞,一般,*,*,*,*,信女,シンニョ,シンニョ
チー	名詞,固有名詞,人名,一般,*,*,チー,チー,チー
南無阿弥陀仏	名詞,一般,*,*,*,*,南無阿弥陀仏,ナムアミダブツ,ナムアミダブツ
南無阿弥陀仏	名詞,一般,*,*,*,*,南無阿弥陀仏,ナムアミダブツ,ナムアミダブツ
EOS
`,
		expect: CountRanking{
			CountRank{5, 1},
			CountRank{4, 2},
			CountRank{2, 3},
			CountRank{2, 3},
			CountRank{2, 3},
			CountRank{1, 6},
			CountRank{1, 6},
			CountRank{1, 6},
		},
	},
	{
		name: "should return all-alike ranking",
		file: "all-alike-ranking.txt.mecab",
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
		expect: CountRanking{
			CountRank{2, 1},
			CountRank{1, 2},
			CountRank{1, 2},
			CountRank{1, 2},
			CountRank{1, 2},
			CountRank{1, 2},
			CountRank{1, 2},
			CountRank{1, 2},
			CountRank{1, 2},
			CountRank{1, 2},
			CountRank{1, 2},
			CountRank{1, 2},
			CountRank{1, 2},
			CountRank{1, 2},
			CountRank{1, 2},
			CountRank{1, 2},
			CountRank{1, 2},
		},
	},
	{
		name: "should return nothing from the text only containing \"EOS\"",
		file: "fail-text.txt.mecab",
		text: `EOS
EOS
EOS`,
		expect: CountRanking(nil),
	},
}

func TestRankByCount(t *testing.T) {
	for _, testcase := range rankByCountsTests {
		t.Log(testcase.name)

		r := strings.NewReader(testcase.text)
		morphs, err := newMorpheme(r)
		if err != nil {
			t.Error(err)
		}

		if results := morphs.sortByAppearance().rankByCounts(); !reflect.DeepEqual(results, testcase.expect) {
			t.Errorf("result => %#v\n should contain => %#v\n", results, testcase.expect)
		}
	}
}
