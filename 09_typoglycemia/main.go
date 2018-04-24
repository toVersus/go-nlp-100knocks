package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

const immutableWordCount = 4

func main() {
	text := "I couldn't believe that I could actually understand what I was reading : the phenomenal power of the human mind ."

	fmt.Println(typoglycemia(text))
}

// typoglycemia shuffles the elements of string except the first and last characters of the element.
func typoglycemia(s string) string {
	ss := strings.Fields(s)
	typoglycemia := make([]string, len(ss))
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i, word := range ss {
		if len(word) <= immutableWordCount {
			typoglycemia[i] = word
			continue
		}
		// exclude first character
		b := []byte(word[1:])
		// exclude last character by specifying the length of b
		r.Shuffle(len(b)-1, func(i, j int) {
			b[i], b[j] = b[j], b[i]
		})
		typoglycemia[i] = string(word[0]) + string(b)
	}
	return strings.Join(typoglycemia, " ")
}

// typoglycemiaBySelfImplShuffle shuffles the elements of string except the first and last characters of the element
// by using self implementation of shuffle.
func typoglycemiaBySelfImplShuffle(s string) string {
	ss := strings.Fields(s)
	typoglycemia := make([]string, len(ss))
	seed := time.Now().UnixNano()
	for i, word := range ss {
		if len(word) <= immutableWordCount {
			typoglycemia[i] = word
			continue
		}
		// exclude first character
		b := []byte(word[1:])
		// exclude last character by specifying the length of b
		shuffle(len(b)-1, seed, func(i, j int) {
			b[i], b[j] = b[j], b[i]
		})
		typoglycemia[i] = string(word[0]) + string(b)
	}
	return strings.Join(typoglycemia, " ")
}

// shuffle implements the modern Fisher and Yates' shuffle algorithm.
// https://en.wikipedia.org/wiki/Fisher%E2%80%93Yates_shuffle#The_modern_algorithm
// n is the number of elements and j is a random integer such that 0 <= j < i.
// swap swaps i-th element and j-th element of array (or slice).
func shuffle(n int, seed int64, swap func(i, j int)) {
	rand.Seed(seed)
	for i := n - 1; i > 0; i-- {
		j := int(rand.Intn(i + 1))
		swap(i, j)
	}
}
