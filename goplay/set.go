// set
package main

import (
	"fmt"
	"strings"
)

// type Set map[string]int
type Set map[string]int

// Create a new Set.
func New(e ...string) (Set, int) {
	s := make(Set)
	n := s.add(e...)
	return s, n
}

// Returns how many new elements were added
func (s Set) add(e ...string) (newCnt int) {
	for _, v := range e {
		if _, in := s[v]; !in {
			newCnt++
			s[v]++
		} else {
			s[v]++
		}
	}
	return
}

func main() {
	set, n := New("s", "t", "a", "r", "t")
	fmt.Printf("%v (%v new)\n", set, n)

	n = set.add([]string{}...)
	fmt.Printf("%v (%v new)\n", set, n)

	n = set.add("")
	fmt.Printf("%v (%v new)\n", set, n)

	n = set.add([]string{"a", "b", "c"}...) // flatten []string
	fmt.Printf("%v (%v new)\n", set, n)

	word := "abracadabra"
	list := strings.Split(word, "")
	n = set.add(list...)
	fmt.Printf("%v (%v new)\n", set, n)
}
