package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Printf("%#v\n", biGramChar("I am an NLPer"))
}

// biGramChar concatenates adjecent pairs of byte by just using plus operator.
func biGramChar(s string) []string {
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
