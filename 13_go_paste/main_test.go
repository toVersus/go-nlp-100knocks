package main

import (
	"reflect"
	"strings"
	"testing"
)

var pasteTests = []struct {
	name   string
	text1  string
	text2  string
	expect string
}{
	{
		name: "should output the content of input files side-by-side",
		text1: `1
2
3
4
5`,
		text2: `Caspian_Sea
Lake_Superior
Lake_Victoria
Lake_Huron
Lake_Michigan`,
		expect: `1	Caspian_Sea
2	Lake_Superior
3	Lake_Victoria
4	Lake_Huron
5	Lake_Michigan`,
	},
}

func TestPaste(t *testing.T) {
	for _, testcase := range pasteTests {
		t.Log(testcase.name)

		r1 := strings.NewReader(testcase.text1)
		r2 := strings.NewReader(testcase.text2)
		result, err := paste(r1, r2)
		if err != nil {
			t.Error(err)
		}
		if !reflect.DeepEqual(result, testcase.expect) {
			t.Errorf("result => %#v\n expect => %#v\n", result, testcase.expect)
		}
	}
}

func TestPasteByChannel(t *testing.T) {
	for _, testcase := range pasteTests {
		t.Log(testcase.name)

		r1 := strings.NewReader(testcase.text1)
		r2 := strings.NewReader(testcase.text2)
		result := pasteByChannel(r1, r2)
		if !reflect.DeepEqual(result, testcase.expect) {
			t.Errorf("result => %#v\n expect => %#v\n", result, testcase.expect)
		}
	}
}
