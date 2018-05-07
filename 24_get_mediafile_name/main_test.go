package main

import (
	"os"
	"reflect"
	"testing"
)

var mediaFileTests = []struct {
	name    string
	file    string
	text    string
	keyword string
	expect  []string
}{
	{
		name: "should return only one media file in the article titled \"UK\"",
		file: "./full-test.json",
		text: `{"text": "イギリスではないです。\n[[file: Liberdade sao paulo.jpg|thumb|[[サンパウロ]]の[[日本人街]]「[[リベルダージ]]」。]]\n[[ファイル:Feiraliberdadesaopaulo.jpg|thumb|left|250px|地域で有名な見本市。]]", "title":"イギリス"}
{"text": "イギリスです。\n[[ファイル:Egypt Topography.png|thumb|200px]]", "title":"イギリス"}`,
		keyword: "イギリス",
		expect: []string{
			" Liberdade sao paulo.jpg",
			"Feiraliberdadesaopaulo.jpg",
			"Egypt Topography.png",
		},
	},
	{
		name:    "should return no media file",
		file:    "./empty-test.json",
		text:    `{"text": "{{otheruses|主に現代のエジプト・アラブ共和国|古代|古代エジプト}}\n{{基礎情報 国\n|略名 =エジプト\n|日本語国名 =エジプト・アラブ共和国\n|公式国名 ='''{{lang|ar|جمهورية مصر العربية}}'''\n|国旗画像 =Flag of Egypt.svg\n|国章画像 =[[ファイル:Coat_of_arms_of_Egypt.svg|100px|エジプトの国章]]\n|国章リンク =（[[エジプトの国章|国章]]）\n|標語 =なし\n|位置画像 =Egypt (orthographic projection).svg\n|公用語 =[[アラビア語]]\n|首都 =[[カイロ]]\n|最大都市 =カイロ\n|元首等肩書 =[[近代エジプトの国家元首の一覧|大統領]]\n|元首等氏名 =[[アブドルファッターフ・アッ＝シーシー]]\n|首相等肩書 =[[エジプトの首相|首相]]\n|首相等氏名 =[[イブラヒーム・メフレブ]]\n|面積順位 =29\n|面積大きさ =1 E12\n|面積値 =1,001,450\n|水面積率 =0.6%\n|人口統計年 =2011\n|人口順位 =\n|人口大きさ =1 E7\n|人口値 =81,120,000\n|人口密度値 =76\n|GDP統計年元 =2008\n|GDP値元 =8,965億<ref name=\"economy\">IMF Data and Statistics 2009年4月27日閲覧（[http://www.imf.org/external/pubs/ft/weo/2009/01/weodata/weorept.aspx?pr.x=77&pr.y=19&sy=2008&ey=2008&scsm=1&ssd=1&sort=country&ds=.&br=1&c=469&s=NGDP%2CNGDPD%2CPPPGDP%2CPPPPC&grp=0&a=]）</ref>", "title":"エジプト"}`,
		keyword: "イギリス",
		expect:  nil,
	},
}

func TestExtractMediaFile(t *testing.T) {
	for _, testcase := range mediaFileTests {
		t.Log(testcase.name)

		f, err := os.Create(testcase.file)
		if err != nil {
			t.Errorf("could not create a file: %s\n  %s\n", testcase.file, err)
		}
		f.WriteString(testcase.text)
		f.Close()

		artists, err := readJSON(testcase.file)
		if err != nil {
			t.Error(err)
		}

		results := artists.find(testcase.keyword).getMediaFileName()
		if !reflect.DeepEqual(results, testcase.expect) {
			t.Errorf("result => %#v\n shpould contain => %#v\n", results, testcase.expect)
		}

		if err := os.Remove(testcase.file); err != nil {
			t.Errorf("could not delete a file: %s\n  %s\n", testcase.file, err)
		}
	}
}

func BenchmarkExtractMediaFile(b *testing.B) {
	for _, testcase := range mediaFileTests {
		f, err := os.Create(testcase.file)
		if err != nil {
			b.Errorf("could not create a file: %s\n  %s\n", testcase.file, err)
		}
		f.WriteString(testcase.text)
		f.Close()
	}

	for i := 0; i < b.N; i++ {
		for _, testcase := range mediaFileTests {
			articles, err := readJSON(testcase.file)
			if err != nil {
				b.Error(err)
			}
			articles.find(testcase.keyword).getMediaFileName()
		}
	}

	for _, testcase := range mediaFileTests {
		if err := os.Remove(testcase.file); err != nil {
			b.Errorf("could not delete a file: %s\n  %s\n", testcase.file, err)
		}
	}
}
