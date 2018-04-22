package main

import (
	"fmt"
)

func main() {
	var (
		x = 12
		y = "気温"
		z = 22.4
	)
	fmt.Println(newTemplate(x, y, z))
}

// newTemplate creates sentence with embeded arguments
func newTemplate(time interface{}, str interface{}, temp interface{}) string {
	return fmt.Sprintf("%d時の%sは%.1f", time, str, temp)
}
