package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	fmt.Printf("%#v\n", insertChar("ace", "bdf"))
}

func insertChar(str1, str2 string) string {
	r1, r2 := []rune(str1), []rune(str2)
	for i := 0; i < len(r1); i++ {
		// Insert rune into odd-numbered index
		// https://github.com/golang/go/wiki/SliceTricks#insert
		r2 = append(r2, 0)
		copy(r2[2*i+1:], r2[2*i:])
		r2[2*i] = r1[i]
	}
	return string(r2)
}

// catCharAltenately concatenates head of each string alternately
// ("ace", "bdf") -> ("abcdef")
func catCharAltenately(str1, str2 string) string {
	l1, l2 := len(str1), len(str2)
	if l1 == 0 && l2 == 0 {
		return ""
	} else if l1 == 0 {
		return str2
	} else if l2 == 0 {
		return str1
	}

	r1, size1 := utf8.DecodeRuneInString(str1)
	r2, size2 := utf8.DecodeRuneInString(str2)
	return string(r1) + string(r2) + catCharAltenately(str1[size1:], str2[size2:])
}

// pileUpLeadChar concatenates leading character of each string alternately
// ("ace", "bdf") -> ("abcdef")
func pileUpLeadChar(str1, str2 string) string {
	r1, r2 := []rune(str1), []rune(str2)
	r := make([]rune, len(r1)+len(r2))
	for i := 0; i < minRuneLen(r1, r2); i++ {
		r[2*i] = r1[i]
		r[2*i+1] = r2[i]
	}
	return string(r)
}

func minRuneLen(r1, r2 []rune) int {
	l1, l2 := len(r1), len(r2)
	if l1 < l2 {
		return l1
	}
	return l2
}
