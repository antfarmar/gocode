// typeassertion
package main

import (
	"fmt"
)

func main() {
	// E is an empty interface variable
	var E interface{}

	var i int = 5
	var s string = "Hello world"

	// These are legal statements
	E = i
	fmt.Println(E.(int) + E.(int)) // type assertion needed for operations
	E = s
	fmt.Println(E.(string) + E.(string)) // type assertion
}
