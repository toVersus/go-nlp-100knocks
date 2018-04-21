package main

import (
	"fmt"
	"strings"
)

func main() {
	str1 := "paraparaparadise"
	str2 := "paragraph"
	x := NewBigram(str1)
	y := NewBigram(str2)
	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(x.Union(y))
	fmt.Println(x.Intersect(y))
	fmt.Println(x.Difference(y))
	fmt.Println(x.Contains("se"))
	fmt.Println(y.Contains("se"))
}

// Bigram represents Set type for Golang.
type Bigram map[string]struct{}

// Set implements fundamental set operators.
type Set interface {
	Add(s string) Set
	Contains(s string) bool
	Union(other Bigram) Set
	Intersect(other Bigram) Set
	Difference(other Bigram) Set
}

// NewBigram concatenates adjecent pairs of character and returns pairs removing duplicated ones.
func NewBigram(s string) Bigram {
	var tmp string
	pairs := make(map[string]struct{})
	for _, char := range strings.Split(s, "") {
		if tmp != "" {
			pairs[tmp+char] = struct{}{}
		}
		tmp = char
	}
	return pairs
}

// Add adds element into the set.
func (set Bigram) Add(s string) Set {
	set[s] = struct{}{}
	return set
}

// Contains returns true in case that the element of set matches to string.
func (set Bigram) Contains(s string) bool {
	if _, ok := set[s]; !ok {
		return false
	}
	return true
}

// Union returns the set of all elements in the collection.
func (set Bigram) Union(other Bigram) Set {
	union := make(Bigram)

	for i := range set {
		union.Add(i)
	}

	for j := range other {
		union.Add(j)
	}

	return union
}

// Intersect returns the set that contains all elements of A as well as B.
func (set Bigram) Intersect(other Bigram) Set {
	intersect := make(Bigram)

	for i := range set {
		if other.Contains(i) {
			intersect.Add(i)
		}
	}

	return intersect
}

// Difference returns the set of differences of each pairs.
func (set Bigram) Difference(other Bigram) Set {
	diff := make(Bigram)

	for i := range set {
		if !other.Contains(i) {
			diff.Add(i)
		}
	}

	return diff
}
