package main

import (
	"bufio"
	"os"
	"reflect"
	"testing"
)

var pasteTests = []struct {
	name   string
	file1  string
	file2  string
	dest   string
	text1  string
	text2  string
	expect string
}{
	{
		name:  "should output the content of input files side-by-side",
		file1: "col-test1.txt",
		file2: "col-test2.txt",
		dest:  "test.txt",
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
5	Lake_Michigan
`,
	},
}

func TestPasteByChannel(t *testing.T) {
	for _, testcase := range pasteTests {
		t.Log(testcase.name)

		file1, err := os.Create(testcase.file1)
		if err != nil {
			t.Errorf("could not create a file: %s\n  %s", testcase.file1, err)
		}
		file1.WriteString(testcase.text1)
		file1.Close()

		file2, err := os.Create(testcase.file2)
		if err != nil {
			t.Errorf("could not create a file: %s\n  %s", testcase.file2, err)
		}
		file2.WriteString(testcase.text2)
		file2.Close()

		pasteByChannel(testcase.file1, testcase.file2, testcase.dest)
		dest, err := os.Open(testcase.dest)
		if err != nil {
			t.Errorf("could not create a file: %s\n  %s", testcase.dest, err)
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

		if err := os.Remove(testcase.file1); err != nil {
			t.Errorf("could not delete a file: %s\n  %s\n", testcase.file1, err)
		}
		if err := os.Remove(testcase.file2); err != nil {
			t.Errorf("could not delete a file: %s\n  %s\n", testcase.file2, err)
		}
		if err := os.Remove(testcase.dest); err != nil {
			t.Errorf("could not delete a file: %s\n  %s\n", testcase.dest, err)
		}
	}
}
