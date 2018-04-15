package main

import (
	"reflect"
	"testing"
)

var reverseTests = []struct {
	name   string
	str    string
	expect string
}{
	{
		name:   "should return byte array in reverse order",
		str:    "stressed",
		expect: "desserts",
	},
}

func TestReverseStringRecursively(t *testing.T) {
	for _, testcase := range reverseTests {
		t.Log(testcase.name)
		b := []byte(testcase.str)
		if result := reverseStringRecursively(b); !reflect.DeepEqual(string(result), testcase.expect) {
			t.Errorf("result => %s\n expect => %s\n", result, testcase.expect)
		}
	}
}

func TestReverseBySwapper(t *testing.T) {
	for _, testcase := range reverseTests {
		t.Log(testcase.name)
		result := []byte(testcase.str)
		if reverseBySwapper(result); !reflect.DeepEqual(string(result), testcase.expect) {
			t.Errorf("result => %s\n expect => %s\n", string(result), testcase.expect)
		}
	}
}

func TestReverseBySelfImplSwapper(t *testing.T) {
	for _, testcase := range reverseTests {
		t.Log(testcase.name)
		if result := reverseBySelfImplSwapper(testcase.str); !reflect.DeepEqual(result, testcase.expect) {
			t.Errorf("result => %s\n expect => %s\n", result, testcase.expect)
		}
	}
}

func BenchmarkReverseStringRecursively(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for _, testcase := range reverseTests {
			b := []byte(testcase.str)
			reverseStringRecursively(b)
		}
	}
}

func BenchmarkReverseBySwapper(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for _, testcase := range reverseTests {
			b := []byte(testcase.str)
			reverseBySwapper(b)
		}
	}
}

func BenchmarkReverseBySelfImplSwapper(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for _, testcase := range reverseTests {
			reverseBySelfImplSwapper(testcase.str)
		}
	}
}

func BenchmarkReverseByDeferStringBuilder(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for _, testcase := range reverseTests {
			buf := *reverseByDeferStringBuilder(testcase.str)
			buf.String()
		}
	}
}
