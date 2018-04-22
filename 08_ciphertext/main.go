package main

import (
	"fmt"
)

func main() {
	msg := "Cipher Testing!"
	fmt.Printf("encrypted: %#v\n", forge(msg))
	fmt.Printf("decrypted: %#v\n", forge(forge(msg)))
}

// forge encrypts the text according to the following rune:
// If it contains lower-case letters, convert that letter to the another one
func forge(s string) string {
	// Create new container for encrypted byte array
	// because string in golang cannot be rewritten or overwritten to anything
	b := make([]byte, len(s))
	for i := range s {
		if s[i] >= 'a' && s[i] <= 'z' {
			b[i] = 219 - s[i]
		} else {
			b[i] = s[i]
		}
	}
	return string(b)
}
