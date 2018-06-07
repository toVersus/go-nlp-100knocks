package main

import (
	"reflect"
	"strings"
	"testing"
)

var countLineTests = []struct {
	name   string
	text   string
	expect int
}{
	{
		name: "should return the number of lines from full text",
		text: `Sed ut perspiciatis unde omnis iste natus error sit voluptatem accusantium doloremque laudantium, totam rem aperiam, eaque ipsa quae ab illo inventore veritatis et quasi architecto beatae vitae dicta sunt explicabo.

		Nemo enim ipsam voluptatem quia voluptas sit aspernatur aut odit aut fugit, sed quia consequuntur magni dolores eos qui ratione voluptatem sequi nesciunt.
		
		Neque porro quisquam est, qui dolorem ipsum quia dolor sit amet, consectetur, adipisci velit, sed quia non numquam eius modi tempora incidunt ut labore et dolore magnam aliquam quaerat voluptatem.
		
		Ut enim ad minima veniam, quis nostrum exercitationem ullam corporis suscipit laboriosam, nisi ut aliquid ex`,
		expect: 7,
	},
	{
		name:   "should return 0 due to the empty file",
		text:   ``,
		expect: 0,
	},
}

func TestCountLineByScanner(t *testing.T) {
	for _, testcase := range countLineTests {
		t.Log(testcase.name)

		r := strings.NewReader(testcase.text)
		result, err := countLineByScanner(r)
		if err != nil {
			t.Error(err)
		}
		if !reflect.DeepEqual(result, testcase.expect) {
			t.Errorf("result => %#v\n expect => %#v", result, testcase.expect)
		}
	}
}

var result int

func BenchmarkCountLineByScanner(b *testing.B) {
	var n int
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		r := strings.NewReader(countLineTests[0].text)
		n, _ = countLineByScanner(r)
	}
	result = n
}

func TestCountLineByReadLine(t *testing.T) {
	for _, testcase := range countLineTests {
		t.Log(testcase.name)

		r := strings.NewReader(testcase.text)
		result, err := countLineByReadLine(r)
		if err != nil {
			t.Error(err)
		}
		if !reflect.DeepEqual(result, testcase.expect) {
			t.Errorf("result => %#v\n expect => %#v", result, testcase.expect)
		}
	}
}

func BenchmarkCountLineByReadLine(b *testing.B) {
	var n int
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		r := strings.NewReader(countLineTests[0].text)
		n, _ = countLineByReadLine(r)
	}
	result = n
}
