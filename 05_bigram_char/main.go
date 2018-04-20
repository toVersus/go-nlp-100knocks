package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Printf("%#v\n", biGramChar("I am an NLPer"))
}

// biGramCharByConcat concatenates adjecent pair of characters after converting rune.
func biGramCharByConcat(s string) []string {
	r := []rune(strings.Replace(s, " ", "", -1))
	// Handle unexpected string input.
	// ss = ""       => []string{""}
	// ss = "foobar" => []string{"foobar"}
	if len(r) <= 1 {
		return []string{string(r)}
	}

	list := make([]string, len(r)-1)
	for i := 0; i < len(r)-1; i++ {
		list[i] = string(r[i : i+2])
	}
	return list
}

// biGramChar concatenates adjecent pair of characters copying from the string.
func biGramChar(str string) []string {
	s := strings.Replace(str, " ", "", -1)

	if len(str) <= 1 {
		return []string{str}
	}

	bigram := make([]string, len(s)-1)
	for i := 0; i < len(s)-1; i++ {
		bigram[i] = s[i : i+2]
	}
	return bigram
}
