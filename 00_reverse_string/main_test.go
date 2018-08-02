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
		name:   "should return string in reverse order",
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

var result string

func BenchmarkReverseStringRecursively(b *testing.B) {
	var s string
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for _, testcase := range reverseTests {
			b := []byte(testcase.str)
			s = string(reverseStringRecursively(b))
		}
	}
	result = s
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
	var s string
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for _, testcase := range reverseTests {
			s = reverseBySelfImplSwapper(testcase.str)
		}
	}
	result = s
}

func BenchmarkReverseByDeferStringBuilder(b *testing.B) {
	var s string
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for _, testcase := range reverseTests {
			buf := *reverseByDeferStringBuilder(testcase.str)
			s = buf.String()
		}
	}
	result = s
}
