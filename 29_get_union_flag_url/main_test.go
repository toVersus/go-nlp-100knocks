package main

import (
	"os"
	"reflect"
	"testing"
)

var getUnionFlagURLTests = []struct {
	name    string
	file    string
	text    string
	keyword string
	expect  string
}{
	{
		name:    "should return formatted value removing emphatic syntax",
		file:    "./test.json",
		text:    `{"text": "イギリスです。{{基礎情報 国\n|日本語国名 = グレートブリテン及び北アイルランド連合王国\n|国旗画像 = Flag of the United Kingdom.svg\n|国歌 = [[女王陛下万歳|神よ女王陛下を守り給え]]\n|位置画像 = Location_UK_EU_Europe_001.svg\n}}\n", "title":"イギリス"}`,
		keyword: "イギリス",
		expect:  "https://upload.wikimedia.org/wikipedia/en/a/ae/Flag_of_the_United_Kingdom.svg",
	},
}

func TestGetUnionFlagURL(t *testing.T) {
	for _, testcase := range getUnionFlagURLTests {
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
		results, err := articles.find(testcase.keyword).getUnionFlagURL()
		if err != nil {
			t.Error(err)
		}

		for _, page := range results.PageSlice() {
			for _, imageInfo := range page.ImageInfo {
				if !reflect.DeepEqual(imageInfo.URL, testcase.expect) {
					t.Errorf("result => %#v\n should contain => %#v\n", imageInfo.URL, testcase.expect)
				}
			}
		}

		if err := os.Remove(testcase.file); err != nil {
			t.Errorf("could not delete a file: %s\n  %s\n", testcase.file, err)
		}
	}
}
