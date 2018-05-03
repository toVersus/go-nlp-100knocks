package main

import (
	"os"
	"reflect"
	"testing"
)

var getCategoryNameTest = []struct {
	name    string
	file    string
	text    string
	keyword string
	expect  []string
}{
	{
		name: "should return 5 category names",
		file: "test.txt",
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

		f, err := os.Create(testcase.file)
		if err != nil {
			t.Errorf("could not create a file: %s\n  %s\n", testcase.file, err)
		}
		f.Write([]byte(testcase.text))
		f.Close()

		articles, err := readJSON(testcase.file)
		if err != nil {
			t.Error(err)
		}

		for i, result := range articles.find(testcase.keyword).getCategoryName() {
			if !reflect.DeepEqual(result, testcase.expect[i]) {
				t.Errorf("result => %#v\n should contain => %#v\n", result, testcase.expect[i])
			}
		}

		if err := os.Remove(testcase.file); err != nil {
			t.Errorf("could not delete a file: %s\n  %s\n", testcase.file, err)
		}
	}
}

func BenchmarkGetCategoryName(b *testing.B) {
	for _, testcase := range getCategoryNameTest {
		f, err := os.Create(testcase.file)
		if err != nil {
			b.Errorf("could not create a file: %s\n  %s\n", testcase.file, err)
		}
		f.Write([]byte(testcase.text))
		f.Close()
	}

	for i := 0; i < b.N; i++ {
		for _, testcase := range getCategoryNameTest {
			articles, err := readJSON(testcase.file)
			if err != nil {
				b.Error(err)
			}

			articles.find(testcase.keyword).getCategoryName()

		}
	}

	for _, testcase := range getCategoryNameTest {
		if err := os.Remove(testcase.file); err != nil {
			b.Errorf("could not delete a file: %s\n  %s\n", testcase.file, err)
		}
	}
}
