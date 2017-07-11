// sortedmap
package main

import (
	"fmt"
	"sort"
)

func main() {
	n := 26
	m := make(map[int]string, n)
	keys := make([]int, 0, 26)

	// Populate map
	for i := 0; i < n; i++ {
		m[i] = string(i + 'A')
	}
	fmt.Printf("%v\n", m)

	// Parallel int slice of map keys
	for k, v := range m {
		keys = append(keys, k)
		fmt.Printf("%v ", v)
	}

	// Sort slice to print map in order
	sort.Ints(keys)
	fmt.Println("\n\nSORTED:")
	for _, k := range keys {
		fmt.Printf("%v ", m[k])
	}
	fmt.Println()
}
