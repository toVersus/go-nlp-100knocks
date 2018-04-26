package main

import (
	"bufio"
	"os"
	"reflect"
	"testing"
)

var cutTests = []struct {
	name      string
	text      string
	columnNum int
	srcPath   string
	destPath  string
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
		srcPath:   "./test_src.txt",
		destPath:  "./test_dest.txt",
		expect: `Caspian_Sea
Lake_Superior
Lake_Victoria
Lake_Huron
Lake_Michigan
`,
	},
	{
		name: "should extract out of column of text",
		text: `1	Caspian_Sea	436,000	78,200
2	Lake_Superior	82,100	12,100
3	Lake_Victoria	68,870	2,760
4	Lake_Huron	59,600	3,540
5	Lake_Michigan	57,800	4,920`,
		columnNum: 10,
		srcPath:   "./test_src.txt",
		destPath:  "./test_dest.txt",
		expect:    "",
	},
}

func TestCut(t *testing.T) {
	for _, testcase := range cutTests {
		t.Log(testcase.name)

		src, err := os.Create(testcase.srcPath)
		if err != nil {
			t.Errorf("could not create a file: %s\n  %s", testcase.srcPath, err)
		}
		src.Write([]byte(testcase.text))
		src.Close()

		cut(testcase.srcPath, testcase.destPath, testcase.columnNum)

		dest, err := os.Open(testcase.destPath)
		if err != nil {
			t.Errorf("could not create a file: %s\n  %s", testcase.destPath, err)
		}

		// Use bufio scan loop to get text from dest file
		sc := bufio.NewScanner(dest)
		var result string
		for sc.Scan() {
			result += sc.Text() + "\n"
		}
		if !reflect.DeepEqual(result, testcase.expect) {
			t.Errorf("result => %#v\n expect => %#v\n", result, testcase.expect)
		}
		dest.Close()

		if err := os.Remove(testcase.srcPath); err != nil {
			t.Errorf("could not delete a file: %s\n  %s\n", testcase.srcPath, err)
		}
		if err := os.Remove(testcase.destPath); err != nil {
			t.Errorf("could not delete a file: %s\n  %s\n", testcase.destPath, err)
		}
	}
}
