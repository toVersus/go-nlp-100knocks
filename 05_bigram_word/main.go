package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Printf("%#v\n", bigramWordByConcat("I am an NLPer"))
}

// bigramWordByCopy concatenates adjecent pairs of word
// by reusing the input string and suppressing the allocation.
func bigramWordByCopy(str string) []string {
	ss := strings.Split(str, " ")

	// Handle unexpected string input:
	// ss = ""       => []string{""}
	// ss = "foobar" => []string{"foobar"}
	if len(ss) <= 1 {
		return ss
	}

	bigram := make([]string, len(ss)-1)
	for i := 0; i < len(ss)-1; i++ {
		// Counts the length of primary and secondary words and whitespace
		// and copy it to the element of slice.
		bigram[i] = str[:(len(ss[i])+1)+len(ss[i+1])]
		// Drop the primary word and whitespace.
		str = str[len(ss[i])+1:]
	}
	return bigram
}

// bigramWordByConcat concatenates adjecent pairs of word
// just by using plus operator.
func bigramWordByConcat(str string) []string {
	ss := strings.Split(str, " ")

	// Handle unexpected string input:
	// ss = ""       => []string{""}
	// ss = "foobar" => []string{"foobar"}
	if len(ss) <= 1 {
		return ss
	}

	bigram := make([]string, len(ss)-1)
	for i := 0; i < len(ss)-1; i++ {
		bigram[i] = ss[i] + " " + ss[i+1]
	}
	return bigram
}

// bigramWordByJoin concatenates adjecent pairs of string
// by using strings.Join function.
func bigramWordByJoin(str string) []string {
	ss := strings.Split(str, " ")

	// Handle unexpected string input:
	// ss = ""       => []string{""}
	// ss = "foobar" => []string{"foobar"}
	if len(ss) <= 1 {
		return ss
	}

	bigram := make([]string, len(ss)-1)
	for i := 0; i < len(ss)-1; i++ {
		bigram[i] = strings.Join([]string{ss[i], ss[i+1]}, " ")
	}
	return bigram
}

// bigramWordByAppend concatenates adjecent pairs of string
// by using variable-length array and append.
// Append creates new instance at each time when called, so it is high cost...
func bigramWordByAppend(str string) []string {
	ss := strings.Split(str, " ")

	// Handle unexpected string input:
	// ss = ""       => []string{""}
	// ss = "foobar" => []string{"foobar"}
	if len(ss) <= 1 {
		return ss
	}

	var tmp string
	var bigram []string
	for i, s := range ss {
		if i != 0 {
			bigram = append(bigram, strings.Join([]string{tmp, s}, " "))
		}
		tmp = s
	}
	return bigram
}
