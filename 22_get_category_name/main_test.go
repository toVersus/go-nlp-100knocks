package main

import (
	"reflect"
	"strings"
	"testing"
)

var getCategoryNameTest = []struct {
	name    string
	text    string
	keyword string
	expect  []string
}{
	{
		name: "should return 3 category names",
		text: `{"text": "インドです。\nイギリスに関係します。\n[[Category:インド|*]]\n[[Category:イギリス連邦]]", "title":"インド"}
{"text": "イギリスです。\n[[Category:イギリス|イギリス]]\n[[Category:英連邦王国|*]]\n[[Category:海洋国家]]", "title":"イギリス"}`,
		keyword: "イギリス",
		expect: []string{
			"イギリス",
			"英連邦王国",
			"海洋国家",
		},
	},
}

func TestGetCategoryName(t *testing.T) {
	for _, testcase := range getCategoryNameTest {
		t.Log(testcase.name)

		r := strings.NewReader(testcase.text)
		filtered, err := readJSON(r)
		if err != nil {
			t.Error(err)
		}
		for i, result := range filtered.find(testcase.keyword).getCategoryName() {
			if !reflect.DeepEqual(result, testcase.expect[i]) {
				t.Errorf("result => %#v\n should contain => %#v\n", result, testcase.expect[i])
			}
		}
	}
}

var result []string

func BenchmarkGetCategoryName(b *testing.B) {
	var categories []string
	for i := 0; i < b.N; i++ {
		for _, testcase := range getCategoryNameTest {
			b.StopTimer()
			r := strings.NewReader(testcase.text)
			filtered, _ := readJSON(r)
			b.StartTimer()

			categories = filtered.find(testcase.keyword).getCategoryName()
		}
	}
	result = categories
}
