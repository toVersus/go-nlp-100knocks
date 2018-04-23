package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

func main() {
	text := "I couldn't believe that I could actually understand what I was reading : the phenomenal power of the human mind ."
	fmt.Println(typoglycemia(text))
}

// typoglycemia shuffles the elements of string extracted from the input string separated with a space,
// except for the first and last characters of the element.
func typoglycemia(s string) string {
	ss := strings.Fields(s)
	typoglycemia := make([]string, len(ss))
	seed := time.Now().UnixNano()
	for i, word := range ss {
		if len(word) <= 4 {
			typoglycemia[i] = word
		} else {
			b := []byte(word[1:])
			shuffle(len(b)-1, seed, func(i, j int) {
				b[i], b[j] = b[j], b[i]
			})
			typoglycemia[i] = string(word[0]) + string(b)
		}
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
