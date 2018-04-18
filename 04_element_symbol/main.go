package main

import (
	"fmt"
	"strings"
)

func main() {
	str := "Hi He Lied Because Boron Could Not Oxidize Fluorine. New Nations Might Also Sign Peace Security Clause. Arthur King Can."
	a := []int{1, 5, 6, 7, 8, 9, 15, 16, 19}
	fmt.Printf("%#v\n", getElemSymbol(str, a))
}

// getElemSymbol extracts both first and secondary characters from the split word in case of specified index.
// In other cases, it extracts only first character from the split word.
func getElemSymbol(s string, index []int) map[int]string {
	m := make(map[int]string)
	for i, word := range strings.Split(s, " ") {
		if isInArray(index, i+1) {
			m[i+1] = word[:1]
		} else {
			m[i+1] = word[:2]
		}
	}
	return m
}

// isInArray check whether the specified index exists in the slice or not.
func isInArray(a []int, n int) bool {
	for _, j := range a {
		if j == n {
			return true
		}
	}
	return false
}

// https://gist.github.com/c-yan/1b945b04ce246ee53b7229949225eb7c
func getElemSymbolBySet(s string, index []int) map[string]int {
	firstOnly := make(map[int]struct{})
	for _, i := range index {
		firstOnly[i] = struct{}{}
	}
	result := make(map[string]int)
	for i, w := range strings.Split(s, " ") {
		var n int
		if _, ok := firstOnly[i+1]; ok {
			n = 1
		} else {
			n = 2
		}
		result[w[:n]] = i + 1
	}
	return result
}
