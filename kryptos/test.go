package main

import (
	"fmt"
	"strings"
)

const (
	ka = "KRYPTOSABCDEFGHIJLMNQUVWXZ"
	pp = "OIEHARXXXXXXXXXBRAVELY"                                                                            //"PALIMPSEST"
	k1 = "OBKRUOXOGHULBSOLIFBBWFLRVQQPRNGKSSOTWTQSJQSSEKZZWATJKLUDIAWINFBNYPVTTMZFPKWGDKZXTJCDIGKUHUAUEKCAR" //"EMUFPHZLRFAXYUSDJKZLDKRNSHGNFIVJYQTQUXQBQVYUVLLTREVJYQTMKYRDMFD"
)

var n = 0

func main() {
	fmt.Println(strings.Map(qua3, k1))
}

func qua3(r rune) rune {
	ppi := strings.IndexRune(ka, rune(pp[n]))
	cti := strings.IndexRune(ka, r)
	abs := cti - ppi

	if abs < 0 {
		abs = 26 + abs
	}

	n += 1
	if n == len(pp) {
		n = 0
	}

	return rune(ka[abs])
}
