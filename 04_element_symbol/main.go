package main

import (
	"fmt"
	"strings"
)

func main() {
	text := "Hi He Lied Because Boron Could Not Oxidize Fluorine. New Nations Might Also Sign Peace Security Clause. Arthur King Can."
	idx := []int{1, 5, 6, 7, 8, 9, 15, 16, 19}
	fmt.Printf("%#v\n", getElemSymbolWithExtraWork(text, idx))
}

// getElemSymbol extracts both first and secondary characters from the split word in case of specified index.
// In other cases, it extracts only first character from the split word.
func getElemSymbol(str string, idxs []int) map[int]string {
	words := strings.Fields(str)
	symbol := make(map[int]string, len(words))
	for i, word := range words {
		symbol[i+1] = word[:2]
		if isInArray(idxs, i+1) {
			symbol[i+1] = word[:1]
		}
	}
	return symbol
}

// isInArray check whether the specified index exists in the slice or not.
func isInArray(idxs []int, n int) bool {
	for _, idx := range idxs {
		if idx == n {
			return true
		}
	}
	return false
}

// getElemSymbolWithExtraWork extracts both first and secondary characters from the split word.
// Then, it deletes second character from extracted characters in case of specified index.
func getElemSymbolWithExtraWork(str string, idxs []int) map[int]string {
	words := strings.Fields(str)
	symbol := make(map[int]string, len(words))
	for i, word := range words {
		symbol[i+1] = word[:2]
	}

	for _, idx := range idxs {
		symbol[idx] = symbol[idx][:1]
	}
	return symbol
}

// Just refactor the following code. Self implementation of Set type.
// https://gist.github.com/c-yan/1b945b04ce246ee53b7229949225eb7c
func getElemSymbolBySet(str string, index []int) map[int]string {
	firstOnly := make(map[int]struct{})
	for _, i := range index {
		firstOnly[i] = struct{}{}
	}

	words := strings.Fields(str)
	result := make(map[int]string, len(words))
	for i, w := range words {
		n := 1
		if _, ok := firstOnly[i+1]; !ok {
			n = 2
		}
		result[i+1] = w[:n]
	}
	return result
}
