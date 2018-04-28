package main

import (
	"bufio"
	"os"
	"reflect"
	"testing"
)

var tailTests = []struct {
	name   string
	text   string
	n      int
	file   string
	dest   string
	expect string
}{
	{
		name: "should write down the last 3 lines into file",
		text: `1	Caspian_Sea	436,000	78,200
2	Lake_Superior	82,100	12,100
3	Lake_Victoria	68,870	2,760
4	Lake_Huron	59,600	3,540
5	Lake_Michigan	57,800	4,920
`,
		n:    3,
		file: "./test.txt",
		dest: "./result.txt",
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
		n:    40,
		file: "./test.txt",
		dest: "./result.txt",
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

		f, err := os.Create(testcase.file)
		if err != nil {
			t.Errorf("could not create a new file: %s\n  %s\n", testcase.file, err)
		}
		f.Write([]byte(testcase.text))
		f.Close()

		dest, err := os.Create(testcase.dest)
		if err != nil {
			t.Errorf("could not create a new file: %s\n  %s\n", testcase.dest, err)
		}
		if err := tail(testcase.file, testcase.n, *dest); err != nil {
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
