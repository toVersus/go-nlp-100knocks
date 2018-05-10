package main

import (
	"os"
	"reflect"
	"testing"
)

var getTemplateTests = []struct {
	name    string
	file    string
	text    string
	keyword string
	expect  map[string]string
}{
	{
		name:    "should return fields and values of templates as dictionary object by removing markup syntax",
		file:    "./full-test.json",
		text:    `{"text": "{{redirect|UK}}\n{{基礎情報 国\n|略名 = イギリス\n|日本語国名 = グレートブリテン及び北アイルランド連合王国\n|公式国名 = {{lang|en|United Kingdom of Great Britain and Northern Ireland}}<ref>英語以外での正式国名:<br/>\n*{{lang|gd|An Rìoghachd Aonaichte na Breatainn Mhòr agus Eirinn mu Thuath}}（[[スコットランド・ゲール語]]）<br/>\n*{{lang|cy|Teyrnas Gyfunol Prydain Fawr a Gogledd Iwerddon}}（[[ウェールズ語]]）<br/>\n*{{lang|ga|Ríocht Aontaithe na Breataine Móire agus Tuaisceart na hÉireann}}（[[アイルランド語]]）<br/>\n*{{lang|kw|An Rywvaneth Unys a Vreten Veur hag Iwerdhon Glédh}}（[[コーンウォール語]]）<br/>\n*{{lang|sco|Unitit Kinrick o Great Breetain an Northren Ireland}}（[[スコットランド語]]）<br/>\n**{{lang|sco|Claught Kängrick o Docht Brätain an Norlin Airlann}}、{{lang|sco|Unitet Kängdom o Great Brittain an Norlin Airlann}}（アルスター・スコットランド語）</ref>\n|国旗画像 = Flag of the United Kingdom.svg\n|国章画像 = [[ファイル:Royal Coat of Arms of the United Kingdom.svg|85px|イギリスの国章]]\n|国章リンク = （[[イギリスの国章|国章]]）\n|標語 = {{lang|fr|Dieu et mon droit}}<br/>（[[フランス語]]:神と私の権利）\n|国歌 = [[女王陛下万歳|神よ女王陛下を守り給え]]\n|位置画像 = Location_UK_EU_Europe_001.svg\n|公用語 = [[英語]]（事実上）\n|首都 = [[ロンドン]]\n|最大都市 = ロンドン\n|元首等肩書 = [[イギリスの君主|女王]]\n|元首等氏名 = [[エリザベス2世]]\n|首相等肩書 = [[イギリスの首相|首相]]\n|首相等氏名 = [[デーヴィッド・キャメロン]]\n|面積順位 = 76\n|面積大きさ = 1 E11\n|面積値 = 244,820\n|水面積率 = 1.3%\n|人口統計年 = 2011\n|人口順位 = 22\n|人口大きさ = 1 E7\n|人口値 = 63,181,775<ref>[http://esa.un.org/unpd/wpp/Excel-Data/population.htm United Nations Department of Economic and Social Affairs>Population Division>Data>Population>Total Population]</ref>\n|人口密度値 = 246\n|GDP統計年元 = 2012\n|GDP値元 = 1兆5478億<ref name=\"imf-statistics-gdp\">[http://www.imf.org/external/pubs/ft/weo/2012/02/weodata/weorept.aspx?pr.x=70&pr.y=13&sy=2010&ey=2012&scsm=1&ssd=1&sort=country&ds=.&br=1&c=112&s=NGDP%2CNGDPD%2CPPPGDP%2CPPPPC&grp=0&a= IMF>Data and Statistics>World Economic Outlook Databases>By Countrise>United Kingdom]</ref>\n|GDP統計年MER = 2012\n|GDP順位MER = 5\n|GDP値MER = 2兆4337億<ref name=\"imf-statistics-gdp\" />\n|GDP統計年 = 2012\n|GDP順位 = 6\n|GDP値 = 2兆3162億<ref name=\"imf-statistics-gdp\" />\n|GDP/人 = 36,727<ref name=\"imf-statistics-gdp\" />\n|建国形態 = 建国\n|確立形態1 = [[イングランド王国]]／[[スコットランド王国]]<br />（両国とも[[連合法 (1707年)|1707年連合法]]まで）\n|確立年月日1 = [[927年]]／[[843年]]\n|確立形態2 = [[グレートブリテン王国]]建国<br />（[[連合法 (1707年)|1707年連合法]]）\n|確立年月日2 = [[1707年]]\n|確立形態3 = [[グレートブリテン及びアイルランド連合王国]]建国<br />（[[連合法 (1800年)|1800年連合法]]）\n|確立年月日3 = [[1801年]]\n|確立形態4 = 現在の国号「'''グレートブリテン及び北アイルランド連合王国'''」に変更\n|確立年月日4 = [[1927年]]\n|通貨 = [[スターリング・ポンド|UKポンド]] (&pound;)\n|通貨コード = GBP\n|時間帯 = ±0\n|夏時間 = +1\n|ISO 3166-1 = GB / GBR\n|ccTLD = [[.uk]] / [[.gb]]<ref>使用は.ukに比べ圧倒的少数。</ref>\n|国際電話番号 = 44\n|注記 = <references />\n}}\n[[Category:君主国]]\n[[Category:島国|くれいとふりてん]]\n[[Category:1801年に設立された州・地域]]", "title": "イギリス"}`,
		keyword: "イギリス",
		expect: map[string]string{
			"略名":         "イギリス",
			"日本語国名":      "グレートブリテン及び北アイルランド連合王国",
			"公式国名":       "{{lang|en|United Kingdom of Great Britain and Northern Ireland}}<ref>英語以外での正式国名:<br/>",
			"国旗画像":       "Flag of the United Kingdom.svg",
			"国章画像":       "[[ファイル:Royal Coat of Arms of the United Kingdom.svg|85px|イギリスの国章]]",
			"国章リンク":      "国章",
			"標語":         "フランス語",
			"国歌":         "神よ女王陛下を守り給え",
			"位置画像":       "Location_UK_EU_Europe_001.svg",
			"公用語":        "英語",
			"首都":         "ロンドン",
			"最大都市":       "ロンドン",
			"元首等肩書":      "女王",
			"元首等氏名":      "エリザベス2世",
			"首相等肩書":      "首相",
			"首相等氏名":      "デーヴィッド・キャメロン",
			"面積順位":       "76",
			"面積大きさ":      "1 E11",
			"面積値":        "244,820",
			"水面積率":       "1.3%",
			"人口統計年":      "2011",
			"人口順位":       "22",
			"人口大きさ":      "1 E7",
			"人口値":        "63,181,775<ref>[http://esa.un.org/unpd/wpp/Excel-Data/population.htm United Nations Department of Economic and Social Affairs>Population Division>Data>Population>Total Population]</ref>",
			"人口密度値":      "246",
			"GDP統計年元":    "2012",
			"GDP値元":      "1兆5478億<ref name=\"imf-statistics-gdp\">[http://www.imf.org/external/pubs/ft/weo/2012/02/weodata/weorept.aspx?pr.x=70&pr.y=13&sy=2010&ey=2012&scsm=1&ssd=1&sort=country&ds=.&br=1&c=112&s=NGDP%2CNGDPD%2CPPPGDP%2CPPPPC&grp=0&a= IMF>Data and Statistics>World Economic Outlook Databases>By Countrise>United Kingdom]</ref>",
			"GDP統計年MER":  "2012",
			"GDP順位MER":   "5",
			"GDP値MER":    "2兆4337億<ref name=\"imf-statistics-gdp\" />",
			"GDP統計年":     "2012",
			"GDP順位":      "6",
			"GDP値":       "2兆3162億<ref name=\"imf-statistics-gdp\" />",
			"GDP/人":      "36,727<ref name=\"imf-statistics-gdp\" />",
			"建国形態":       "建国",
			"確立形態1":      "1707年連合法",
			"確立年月日1":     "927年／843年",
			"確立形態2":      "1707年連合法",
			"確立年月日2":     "1707年",
			"確立形態3":      "1800年連合法",
			"確立年月日3":     "1801年",
			"確立形態4":      "現在の国号「グレートブリテン及び北アイルランド連合王国」に変更",
			"確立年月日4":     "1927年",
			"通貨":         "UKポンド",
			"通貨コード":      "GBP",
			"時間帯":        "±0",
			"夏時間":        "+1",
			"ISO 3166-1": "GB / GBR",
			"ccTLD":      ".uk / .gb",
			"国際電話番号":     "44",
			"注記":         "<references />",
		},
	},
	{
		name: "should not return variables defined in Article with similar format",
		file: "./test.json",
		text: `{"text": "イギリスではないです。\n{{cite news |title=ニュースサイト免許制度に反対デモ、シンガポール |newspaper=[[TBSテレビ|TBS]]|date=2013-6-9 |url=http://news.tbs.co.jp/newseye/tbs_newseye5353121.html | accessdate=2013-6-9\n}}</ref>。\n", "title":"エジプト"}
`,
		keyword: "イギリス",
		expect:  map[string]string{},
	},
}

func TestGetTemplate(t *testing.T) {
	for _, testcase := range getTemplateTests {
		t.Log(testcase.name)

		f, err := os.Create(testcase.file)
		if err != nil {
			t.Errorf("could not create a file: %s\n  %s\n", testcase.file, err)
		}
		f.WriteString(testcase.text)
		f.Close()

		articles, err := readJSON(testcase.file)
		if err != nil {
			t.Error(err)
		}

		result := articles.find(testcase.keyword).getTemplate()
		if !reflect.DeepEqual(result, testcase.expect) {
			t.Errorf("result => %#v\n should contain => %#v\n", result, testcase.expect)
		}

		if err := os.Remove(testcase.file); err != nil {
			t.Errorf("could not delete a file: %s\n  %s\n", testcase.file, err)
		}
	}
}
