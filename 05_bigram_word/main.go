package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Printf("%#v\n", bigramWordByConcat("I am an NLPer"))
}

// bigramWordByConcat splits a string into strings and concatenates
// adjecent pairs of string just using plus operator.
func bigramWordByConcat(str string) []string {
	ss := strings.Split(str, " ")

	// Handle unexpected string input
	// ss = ""       => []string{""}
	// ss = "foobar" => []string{"foobar"}
	if len(ss) <= 1 {
		return ss
	}

	list := make([]string, len(ss)-1)
	for i := 0; i < len(ss)-1; i++ {
		list[i] = ss[i] + " " + ss[i+1]
	}
	return list
}

// bigramWordByJoin splits a string into strings and concatenates
// adjecent pairs of string using strings.Join function.
func bigramWordByJoin(str string) []string {
	ss := strings.Split(str, " ")

	// Handle unexpected string input
	// ss = ""       => []string{""}
	// ss = "foobar" => []string{"foobar"}
	if len(ss) <= 1 {
		return ss
	}

	list := make([]string, len(ss)-1)
	for i := 0; i < len(ss)-1; i++ {
		list[i] = strings.Join([]string{ss[i], ss[i+1]}, " ")
	}
	return list
}

// bigramWordByAppend splits a string into strings and concatenates
// adjecent pairs of string using variable-length array and append.
// Append creates new instance at each time when called, so it is high cost...
func bigramWordByAppend(str string) []string {
	ss := strings.Split(str, " ")

	// Handle unexpected string input
	// ss = ""       => []string{""}
	// ss = "foobar" => []string{"foobar"}
	if len(ss) <= 1 {
		return ss
	}

	var tmp string
	var list []string
	for i, s := range ss {
		if i != 0 {
			list = append(list, strings.Join([]string{tmp, s}, " "))
		}
		tmp = s
	}
	return list
}
