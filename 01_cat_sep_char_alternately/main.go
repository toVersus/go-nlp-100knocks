package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	fmt.Printf("%s\n", string(catSepCharAlternately("日本はstressed-society", false)))
}

// catSepCharAlternately extracts either odd-numbered or even-numbered elements from byte slice
func catSepCharAlternately(str string, isOddNumbered bool) string {
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
	return string(r) + catSepCharAlternately(str[(size1+size2):], isOddNumbered)
}
