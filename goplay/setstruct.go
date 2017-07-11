// set
package main

import (
	"fmt"
	"strings"
)

// type Set map[string]int
type Set struct {
	m map[string]int
}

func NewSet() (s Set) {
	s = Set{make(map[string]int)}
	return
}

// Returns how many times e was "added"
func (s *Set) add(e string) int {
	s.m[e] += 1
	return s.m[e]
}

// Returns how many new elements were added
func (s *Set) addChars(r string) (newCnt int) {
	for _, c := range r {
		e := string(c)
		if _, in := s.m[e]; !in {
			newCnt++
			s.m[e]++
		} else {
			s.m[e]++
		}
	}
	return
}

// Returns how many new elements were added
func (s *Set) addSlice(es []string) (newCnt int) {
	for _, e := range es {
		if _, in := s.m[e]; !in {
			newCnt++
		}
		s.m[e]++
	}
	return
}

// Stringer "String() string" method
func (s Set) String() string {
	t := fmt.Sprintf("%v", s.m)
	t = fmt.Sprintf("%q #%v", t[3:], len(s.m))
	return t
}

func main() {
	set := NewSet()

	n := set.add("") // legal
	n = set.add("a")
	n = set.add("a")
	fmt.Printf("%v cnt(a)=%v\n", set, n)

	n = set.addSlice([]string{"a", "b", "c", "d"})
	fmt.Printf("%v (%v new)\n", set, n)

	word := "abracadabra"
	n = set.addChars(word)
	fmt.Printf("%v (%v new)\n", set, n)

	n = set.addChars("")
	fmt.Printf("%v (%v new)\n", set, n)

	n = set.addSlice([]string{})
	fmt.Printf("%v (%v new)\n", set, n)

	list := strings.Split(word, "")
	n = set.addSlice(list)
	fmt.Printf("%v (%v new)\n", set, n)
}
