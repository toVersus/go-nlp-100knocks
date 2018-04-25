package main

import (
	"os"
	"reflect"
	"strings"
	"testing"
)

var convertTabToSpaceTests = []struct {
	name   string
	file   string
	text   string
	expect string
}{
	{
		name: "should return text after conveting tab to space",
		file: "text.txt",
		text: `Sed	ut	perspiciatis	unde	omnis iste	natus error sit voluptatem accusantium doloremque laudantium, totam rem aperiam, eaque ipsa quae ab illo inventore veritatis et quasi architecto beatae vitae dicta sunt explicabo.

Nemo enim ipsam voluptatem quia voluptas sit aspernatur	aut odit aut fugit, sed quia consequuntur magni dolores eos qui ratione voluptatem sequi nesciunt.
	
Neque porro	quisquam est, qui dolorem ipsum quia dolor	sit amet, consectetur, adipisci velit, sed quia non numquam eius modi tempora incidunt ut labore et dolore magnam aliquam quaerat voluptatem.

Ut enim ad	minima	veniam,	quis nostrum exercitationem	ullam corporis suscipit laboriosam, nisi ut aliquid ex`,

		expect: `Sed ut perspiciatis unde omnis iste natus error sit voluptatem accusantium doloremque laudantium, totam rem aperiam, eaque ipsa quae ab illo inventore veritatis et quasi architecto beatae vitae dicta sunt explicabo.

Nemo enim ipsam voluptatem quia voluptas sit aspernatur aut odit aut fugit, sed quia consequuntur magni dolores eos qui ratione voluptatem sequi nesciunt.
 
Neque porro quisquam est, qui dolorem ipsum quia dolor sit amet, consectetur, adipisci velit, sed quia non numquam eius modi tempora incidunt ut labore et dolore magnam aliquam quaerat voluptatem.

Ut enim ad minima veniam, quis nostrum exercitationem ullam corporis suscipit laboriosam, nisi ut aliquid ex`,
	},
}

func TestConvertTabToSpace(t *testing.T) {
	for _, testcase := range convertTabToSpaceTests {
		t.Log(testcase.name)

		f, err := os.Create(testcase.file)
		if err != nil {
			t.Errorf("could not create a file: %s\n  %s\n", testcase.file, err)
		}
		f.Write([]byte(testcase.text))

		if result := strings.Join(convertTabToSpace(testcase.file), "\n"); !reflect.DeepEqual(result, testcase.expect) {
			t.Errorf("result => %#v\n expect => %#v", result, testcase.expect)
		}
		f.Close()

		if err := os.Remove(testcase.file); err != nil {
			t.Errorf("could not delete a file: %s\n  %s\n", testcase.file, err)
		}
	}
}
