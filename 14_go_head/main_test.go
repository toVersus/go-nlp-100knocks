package main

import (
	"bufio"
	"os"
	"reflect"
	"testing"
)

var headTests = []struct {
	name   string
	src    string
	text   string
	n      int
	dest   string
	expect string
}{
	{
		name: "should print the first 3 lines of input file",
		src:  "./test.txt",
		text: `1	Caspian_Sea	436,000	78,200
2	Lake_Superior	82,100	12,100
3	Lake_Victoria	68,870	2,760
4	Lake_Huron	59,600	3,540
5	Lake_Michigan	57,800	4,920`,
		n:    3,
		dest: "./result.txt",
		expect: `1	Caspian_Sea	436,000	78,200
2	Lake_Superior	82,100	12,100
3	Lake_Victoria	68,870	2,760
`,
	},
	{
		name:   "should not print anything due to the empty file",
		src:    "./empty-test.txt",
		text:   "",
		n:      3,
		dest:   "./result.txt",
		expect: "",
	},
}

func TestHead(t *testing.T) {
	for _, testcase := range headTests {
		t.Log(testcase.name)

		f, err := os.Create(testcase.src)
		if err != nil {
			t.Errorf("could not create a new file: %s\n  %s", testcase.src, err)
		}
		f.WriteString(testcase.text)
		f.Close()

		dest, err := os.Create(testcase.dest)
		if err != nil {
			t.Errorf("could not create a new file: %s\n  %s", testcase.dest, err)
		}
		if err := head(testcase.src, testcase.n, *dest); err != nil {
			t.Error(err)
		}
		dest.Close()

		dest, err = os.Open(testcase.dest)
		if err != nil {
			t.Errorf("could not open a file: %s\n  %s", testcase.dest, err)
		}

		sc := bufio.NewScanner(dest)
		result := ""
		for sc.Scan() {
			result += sc.Text() + "\n"
		}
		dest.Close()

		if !reflect.DeepEqual(result, testcase.expect) {
			t.Errorf("result => %#v\n expect => %#v\n", result, testcase.expect)
		}

		if err := os.Remove(testcase.src); err != nil {
			t.Errorf("could not delete a file: %s\n  %s\n", testcase.src, err)
		}
		if err := os.Remove(testcase.dest); err != nil {
			t.Errorf("could not delete a file: %s\n  %s\n", testcase.dest, err)
		}
	}
}
