package main

import (
	"reflect"
	"strings"
	"testing"
)

var headTests = []struct {
	name    string
	text    string
	lineNum int
	expect  string
}{
	{
		name: "should print the first 3 lines of input file",
		text: `1	Caspian_Sea	436,000	78,200
2	Lake_Superior	82,100	12,100
3	Lake_Victoria	68,870	2,760
4	Lake_Huron	59,600	3,540
5	Lake_Michigan	57,800	4,920`,
		lineNum: 3,
		expect: `1	Caspian_Sea	436,000	78,200
2	Lake_Superior	82,100	12,100
3	Lake_Victoria	68,870	2,760`,
	},
	{
		name:    "should not print anything due to the empty file",
		text:    "",
		lineNum: 3,
		expect:  "",
	},
}

func TestHead(t *testing.T) {
	for _, testcase := range headTests {
		t.Log(testcase.name)

		r := strings.NewReader(testcase.text)
		result := head(r, testcase.lineNum)
		if !reflect.DeepEqual(result, testcase.expect) {
			t.Errorf("result => %#v\n expect => %#v\n", result, testcase.expect)
		}
	}
}
