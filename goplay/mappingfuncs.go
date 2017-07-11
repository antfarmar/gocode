package main

import (
	"fmt"
	"strings"
)

func main() {
	ss := []string{"A", "B", "C", "D"}
	mapf(strings.ToUpper, mapf(strings.ToLower, ss))
	//fmt.Println(ss)
}

func mapf(f func(string) string, ss []string) []string {
	for i, e := range ss {
		ss[i] = f(e)
	}
	fmt.Println(ss)
	return ss
}
