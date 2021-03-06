package main

import (
	"reflect"
	"testing"
)

var encryptionTests = []struct {
	name   string
	str    string
	expect string
}{
	{
		name:   "should encrypt the text by replacing only the lower-case letters",
		str:    "aTbTcT",
		expect: "zTyTxT",
	},
	{
		name:   "should NOT replace anything while input text contains just a number",
		str:    "1234",
		expect: "1234",
	},
	{
		name:   "should NOT be encrypted due to empty input text",
		str:    "",
		expect: "",
	},
}

func TestEncrypt(t *testing.T) {
	for _, testcase := range encryptionTests {
		t.Log(testcase.name)
		if result := forge(testcase.str); !reflect.DeepEqual(result, testcase.expect) {
			t.Errorf("result => %s\n expect => %s\n", result, testcase.expect)
		}
	}
}

func BenchmarkForge(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for _, testcase := range encryptionTests {
			forge(testcase.str)
		}
	}
}

var decryptionTests = []struct {
	name   string
	str    string
	expect string
}{
	{
		name:   "should decrypt the decrypted text by replacing only the lower-case letters",
		str:    "aTbTcT",
		expect: "aTbTcT",
	},
	{
		name:   "should NOT replace anything while input text contains just a number",
		str:    "1234",
		expect: "1234",
	},
	{
		name:   "should NOT encrypt/decrypt due to the empty input text",
		str:    "",
		expect: "",
	},
}

func TestDecrypt(t *testing.T) {
	for _, testcase := range decryptionTests {
		t.Log(testcase.name)
		if result := forge(forge(testcase.str)); !reflect.DeepEqual(result, testcase.expect) {
			t.Errorf("result => %s\n expect => %s\n", result, testcase.expect)
		}
	}
}
