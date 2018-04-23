package main

import (
	"fmt"
)

func main() {
	msg := "Cipher Testing!"
	fmt.Println("encrypted: ", forge(msg))
	fmt.Println("decrypted: ", forge(forge(msg)))
}

const MaxCharCode = 219

// forge converts the letter to the another one while the input text contains a lower-case letter.
func forge(s string) string {
	// Create new container for encrypted byte array
	// because string in golang cannot be rewritten or overwritten to anything
	b := make([]byte, len(s))
	for i := range s {
		b[i] = s[i]
		if s[i] >= 'a' && s[i] <= 'z' {
			// cannot do this
			// s[i] = MaxCharCode - s[i]
			b[i] = MaxCharCode - s[i]
		}
	}
	return string(b)
}
