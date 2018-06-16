package main

import (
	"reflect"
	"strings"
	"testing"
)

var tailTests = []struct {
	name    string
	text    string
	lineNum int
	expect  string
}{
	{
		name: "should write down the last 3 lines into file",
		text: `1	Caspian_Sea	436,000	78,200
2	Lake_Superior	82,100	12,100
3	Lake_Victoria	68,870	2,760
4	Lake_Huron	59,600	3,540
5	Lake_Michigan	57,800	4,920
`,
		lineNum: 3,
		expect: `3	Lake_Victoria	68,870	2,760
4	Lake_Huron	59,600	3,540
5	Lake_Michigan	57,800	4,920
`,
	},
	{
		name: "should write down the all lines into file",
		text: `1	Caspian_Sea	436,000	78,200
2	Lake_Superior	82,100	12,100
3	Lake_Victoria	68,870	2,760
4	Lake_Huron	59,600	3,540
5	Lake_Michigan	57,800	4,920
`,
		lineNum: 40,
		expect: `1	Caspian_Sea	436,000	78,200
2	Lake_Superior	82,100	12,100
3	Lake_Victoria	68,870	2,760
4	Lake_Huron	59,600	3,540
5	Lake_Michigan	57,800	4,920
`,
	},
}

func TestTail(t *testing.T) {
	for _, testcase := range tailTests {
		t.Log(testcase.name)

		r := strings.NewReader(testcase.text)
		result, err := tail(r, testcase.lineNum)
		if err != nil {
			t.Error(err)
		}
		if !reflect.DeepEqual(result, testcase.expect) {
			t.Errorf("result => %#v\n expect => %#v\n", result, testcase.expect)
		}
	}
}
