package main

import (
	"bytes"
	"fmt"
)

// Strings are read-only, so concatenation creates new strings.
// Using []bytes instead takes O(n) time/memory.
func main() {
	var buffer bytes.Buffer

	for i := 0; i < 100; i++ {
		buffer.WriteString("a")
	}

	fmt.Println(buffer.String())
}
