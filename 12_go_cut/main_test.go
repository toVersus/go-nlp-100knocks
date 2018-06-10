package main

import (
	"reflect"
	"strings"
	"testing"
)

var cutTests = []struct {
	name      string
	text      string
	columnNum int
	expect    string
}{
	{
		name: "should extract second column of text",
		text: `1	Caspian_Sea	436,000	78,200
2	Lake_Superior	82,100	12,100
3	Lake_Victoria	68,870	2,760
4	Lake_Huron	59,600	3,540
5	Lake_Michigan	57,800	4,920`,
		columnNum: 2,
		expect: `Caspian_Sea
Lake_Superior
Lake_Victoria
Lake_Huron
Lake_Michigan`,
	},
	{
		name: "should extract out of column of text",
		text: `1	Caspian_Sea	436,000	78,200
2	Lake_Superior	82,100	12,100
3	Lake_Victoria	68,870	2,760
4	Lake_Huron	59,600	3,540
5	Lake_Michigan	57,800	4,920`,
		columnNum: 10,
		expect:    "",
	},
}

func TestCut(t *testing.T) {
	for _, testcase := range cutTests {
		t.Log(testcase.name)

		r := strings.NewReader(testcase.text)
		if result := cut(r, testcase.columnNum); !reflect.DeepEqual(result, testcase.expect) {
			t.Errorf("result => %#v\n expect => %#v", result, testcase.expect)
		}
	}
}
