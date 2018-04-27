package main

import (
	"bufio"
	"os"
	"reflect"
	"testing"
)

var headTests = []struct {
	name   string
	file   string
	text   string
	n      int
	dest   string
	expect string
}{
	{
		name: "should print the first 3 lines of input file",
		file: "./test.txt",
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
}

func TestHead(t *testing.T) {
	for _, testcase := range headTests {
		t.Log(testcase.name)

		f, err := os.Create(testcase.file)
		if err != nil {
			t.Errorf("could not create a new file: %s\n", err)
		}
		f.Write([]byte(testcase.text))
		f.Close()

		dest, _ := os.Create(testcase.dest)
		if err := head(testcase.file, testcase.n, *dest); err != nil {
			t.Error(err)
		}
		dest.Close()

		dest, _ = os.Open(testcase.dest)
		sc := bufio.NewScanner(dest)
		var result string
		for sc.Scan() {
			result += sc.Text() + "\n"
		}
		dest.Close()

		if !reflect.DeepEqual(result, testcase.expect) {
			t.Errorf("result => %#v\n expect => %#v\n", result, testcase.expect)
		}

		if err := os.Remove(testcase.file); err != nil {
			t.Errorf("could not delete a file: %s\n  %s\n", testcase.file, err)
		}
		if err := os.Remove(testcase.dest); err != nil {
			t.Errorf("could not delete a file: %s\n  %s\n", testcase.dest, err)
		}
	}
}
