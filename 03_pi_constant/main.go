package main

import (
	"fmt"
	"strings"
)

func main() {
	str := "Now I need a drink, alcoholic of course, after the heavy lectures involving quantum mechanics."
	fmt.Println(countWordLenByCounter(str))
}

// countWordLenByCounter implements self counter and counts when the rune matches to the pattern.
// This is an incorrect answer because the input sentence doesn't split into the words.
func countWordLenByCounter(str string) []int {
	if len(str) == 0 {
		return []int{0}
	}
	str = strings.Replace(str, ",", "", -1)

	count := []int{}
	counter := 0
	for _, s := range str {
		if (s != ' ') && (s != '.') {
			counter++
			continue
		}
		count = append(count, counter)
		counter = 0
	}
	return count
}

// countWordLenByRecursiveFunc counts the length of each element splitted into the words.
// It creates new instance when the recursive function is called and destroy old instance...a hell of allocation
func countWordLenByRecursiveFunc(ss []string) []int {
	if len(ss) == 0 {
		return []int{}
	}
	return append([]int{len(ss[0])}, countWordLenByRecursiveFunc(ss[1:])...)
}

// countWordLen counts the length of each element splitted into the words.
func countWordLen(str string) []int {
	ss := strings.Split(str, " ")
	count := make([]int, len(ss))
	for i := 0; i < len(ss); i++ {
		count[i] = len(ss[i])
	}
	return count
}

// countWordLenByAppend counts the length of each element splitted into the words.
// The append should not be used when a length of slice is fixed...
func countWordLenByAppend(str string) []int {
	count := []int{}
	for _, s := range strings.Split(str, " ") {
		count = append(count, len(s))
	}
	return count
}

// countWordLenCallByReference counts the length of each element splitted into the words.
func countWordLenCallByReference(count []int, ss []string) []int {
	if len(ss) == 0 {
		return count
	}
	return countWordLenCallByReference(append(count, len(ss[0])), ss[1:])
}
