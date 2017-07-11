package main

import (
	"fmt"
	//"strconv"
	"strings"
)

var k4 = "OBKRUOXOGHULBSOLIFBBWFLRVQQPRNGKSSOTWTQSJQSSEKZZWATJKLUDIAWINFBNYPVTTMZFPKWGDKZXTJCDIGKUHUAUEKCAR"

//var k4int = []int{14, 1, 10, 17, 20, 14, 23, 14, 6, 7, 20, 11, 1, 18, 14, 11, 8, 5, 1, 1, 22, 5, 11, 17, 21, 16, 16, 15, 17, 13, 6, 10, 18, 18, 14, 19, 22, 19, 16, 18, 9, 16, 18, 18, 4, 10, 25, 25, 22, 0, 19, 9, 10, 11, 20, 3, 8, 0, 22, 8, 13, 5, 1, 13, 24, 15, 21, 19, 19, 12, 25, 5, 15, 10, 22, 6, 3, 10, 25, 23, 19, 9, 2, 3, 8, 6, 10, 20, 7, 20, 0, 20, 4, 10, 2, 0, 17}

//var k4str = []string{"O", "B", "K", "R", "U", "O", "X", "O", "G", "H", "U", "L", "B", "S", "O", "L", "I", "F", "B", "B", "W", "F", "L", "R", "V", "Q", "Q", "P", "R", "N", "G", "K", "S", "S", "O", "T", "W", "T", "Q", "S", "J", "Q", "S", "S", "E", "K", "Z", "Z", "W", "A", "T", "J", "K", "L", "U", "D", "I", "A", "W", "I", "N", "F", "B", "N", "Y", "P", "V", "T", "T", "M", "Z", "F", "P", "K", "W", "G", "D", "K", "Z", "X", "T", "J", "C", "D", "I", "G", "K", "U", "H", "U", "A", "U", "E", "K", "C", "A", "R"}
var abc26 = []string{"A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z"}
var kry26 = []string{"K", "R", "Y", "P", "T", "O", "S", "A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "L", "M", "N", "Q", "U", "V", "W", "X", "Z"}

var bin32 = []string{"00000", "00001", "00010", "00011", "00100", "00101", "00110", "00111", "01000", "01001", "01010", "01011", "01100", "01101", "01110", "01111", "10000", "10001", "10010", "10011", "10100", "10101", "10110", "10111", "11000", "11001", "11010", "11011", "11100", "11101", "11110", "11111"}
var gry32 = []string{"00000", "00001", "00011", "00010", "00110", "00111", "00101", "00100", "01100", "01101", "01111", "01110", "01010", "01011", "01001", "01000", "11000", "11001", "11011", "11010", "11110", "11111", "11101", "11100", "10100", "10101", "10111", "10110", "10010", "10011", "10001", "10000"}

func pos(input []string, value string) int {
	for p, v := range input {
		if v == value {
			return p
		}
	}
	return -1
}

func not(bits string) string {
	slice := strings.Split(bits, "")
	s := ""
	for _, b := range slice {
		switch {
		case b == "0":
			s += "1"
		case b == "1":
			s += "0"
		default:
			panic("not a bit")
		}
	}
	return s
}

func tpose(input []string, a, b int) []string {
	output := make([]string, len(input))
	m := len(input)
	for x := range output {
		output[x] = input[(a*x+b)%m]
	}
	return output
}

func encodeK4(txt string, code, abc []string) []string {
	s := ""
	for _, c := range txt {
		s += code[pos(abc, string(c))] //code[c-'A']
	}
	return strings.Split(s, "")
}

func decodeK4(txt []string, code, abc []string) string {
	nbits := len(code[0])
	s, pt := "", ""
	for i, c := range txt {
		s += c
		if (i+1)%nbits == 0 {
			p := pos(code, s)
			if p > 25 {
				s = not(s)
			}
			p = pos(code, s)
			pt += abc[p]
			s = ""
		}
	}
	return pt
}

func rot1(r rune) rune {
	return 'A' + (r-'A'+1)%26
}

func main() {
	//	a := []string{"A", "B", "C", "D", "E", "F", "G", "H", "I", "J"}
	//	for i := 1; i < len(a); i++ {
	//		t := tpose(a, i, 0)
	//		fmt.Println(t)
	//	}

	code := bin32
	alph := abc26
	ciph := k4

	word := "BERLIN"
    cnt := 1

	for h := 0; h < 26; h++ { // CAESAR SHIFT K4
		enc := encodeK4(ciph, code, alph)

		for i := 1; i < len(enc); i++ { // TRANSPOSE THE BITS
			if i%5 == 0 || i%97 == 0 {
				continue
			}
			t := tpose(enc, i, 0)
			d := decodeK4(t, code, alph)

			for j := 0; j < 26; j++ { // CAESAR SHIFT NEW K4 DECODINGS
				d = strings.Map(rot1, d)
				if strings.Count(d, word) >= cnt {
					d := strings.Replace(d, word, " ["+word+"] ", -1)
					fmt.Println(d, h, j, i)
					//println(ct)
				}
			}
		}
		ciph = strings.Map(rot1, ciph)
	}
}
