package main

import (
	"bufio"
	"fmt"
	"os"
	"reflect"
	"strings"
)

func main() {
	buf := *reverseByDeferStringBuilder("stressed")
	fmt.Printf("%+v\n", buf.String())
}

// reverseStringRecursively rearranges the elements of byte slice in reverse order streamingly
func reverseStringRecursively(b []byte) []byte {
	if len(b) == 0 {
		return b
	}
	return append(reverseStringRecursively(b[1:]), b[0])
}

// reverseBySwapper rearranges the elements of slice from tail to head by using swapper
func reverseBySwapper(slice interface{}) {
	t := reflect.ValueOf(slice)
	swap := reflect.Swapper(slice)
	l := t.Len()
	for i := 0; i < l/2; i++ {
		swap(i, l-i-1)
	}
}

// reverseBySelfImplSwapper reverses input string just by using for-loop self implementation
func reverseBySelfImplSwapper(s string) string {
	t := strings.Split(s, "")
	l := len(t)
	for i := 0; i < l/2; i++ {
		t[i], t[l-i-1] = t[l-i-1], t[i]
	}
	return strings.Join(t, "")
}

// reverseByDefer outputs reversed string by using defer's Last in First Out order characteristics
func reverseByDefer(str string) {
	writer := bufio.NewWriter(os.Stdout)
	for _, s := range str {
		defer writer.Flush()
		defer fmt.Fprint(writer, string(s))
	}
}

// reverseByDeferStringBuilder returns the reversed string within string builder
// by using defer's Last in First Out order characteristics
func reverseByDeferStringBuilder(str string) *strings.Builder {
	var buf strings.Builder
	for _, s := range str {
		defer fmt.Fprint(&buf, string(s))
	}
	return &buf
}
