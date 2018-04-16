package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	fmt.Printf("%s\n", delCharInSequence("パタトクカシーー"))
}

// delCharInSequence deletes the specified elements of rune slice without creating new instance.
func delCharInSequence(str string) string {
	r := []rune(str)
	// Delete each element of slice.
	// https://github.com/golang/go/wiki/SliceTricks#delete
	for _, idx := range []int{1, 2, 3, 4} {
		r = append(r[:idx], r[idx+1:]...)
	}
	return string(r)
}

// catCharInSequence concatenates specified index of string converted from rune slice.
// This is basic and simple answer.
func catCharInSequence(str string) string {
	r := []rune(str)
	result := ""
	for _, i := range []int{0, 2, 4, 6} {
		result += string(r[i])
	}
	return result
}

// catSepCharInSequence extracts either odd-numbered or even-numbered elements from byte slice
func catSepCharInSequence(str string, isOddNumbered bool) string {
	r1, size1 := utf8.DecodeRuneInString(str)
	r2, size2 := utf8.DecodeRuneInString(str[size1:])

	if len(str) == 0 || (isOddNumbered == true && len(str) == size1) {
		return str
	} else if size2 == 0 || len(str) == size1 {
		return ""
	}

	r := r2
	if isOddNumbered {
		r = r1
	}
	return string(r) + catSepCharInSequence(str[(size1+size2):], isOddNumbered)
}
