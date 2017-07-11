package main

import (
	"fmt"
	"math"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

//var nn = []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31}
var k4 = []int{14, 1, 10, 17, 20, 14, 23, 14, 6, 7, 20, 11, 1, 18, 14, 11, 8, 5, 1, 1, 22, 5, 11, 17, 21, 16, 16, 15, 17, 13, 6, 10, 18, 18, 14, 19, 22, 19, 16, 18, 9, 16, 18, 18, 4, 10, 25, 25, 22, 0, 19, 9, 10, 11, 20, 3, 8, 0, 22, 8, 13, 5, 1, 13, 24, 15, 21, 19, 19, 12, 25, 5, 15, 10, 22, 6, 3, 10, 25, 23, 19, 9, 2, 3, 8, 6, 10, 20, 7, 20, 0, 20, 4, 10, 2, 0, 17}
var newk4 []int
var nbits = 5
var nbs = strconv.Itoa(nbits)
var mod = int(math.Exp2(float64(nbits)))

func pos(input []int, value int) int {
	for p, v := range input {
		if v == value {
			return p
		}
	}
	return -1
}

func Add(input []int, add int) {
	for i, n := range input {
		input[i] = (n + add) % mod
	}
}

func Tpose(input []string, i int, off int) []string {
	output := make([]string, len(input))
	m := len(input)
	for j := range output {
		output[j] = input[(i*j+off)%m]
	}
	return output
}

func BitsToStr(input []string) []string {
	output := make([]string, len(k4))
	s := ""
	for i, c := range input {
		s += c
		if (i+1)%nbits == 0 {
			n, _ := strconv.ParseInt(s, 2, 0)
			j := pos(newk4, int(n))
			if j == -1 {
				output[i/nbits] = "-" //fmt.Sprintf("%c", (n%26)+65)
			} else {
				output[i/nbits] = fmt.Sprintf("%c", k4[j]+65) //fmt.Sprintf("%c", n+65)
			}
			s = ""
		}
	}
	return output
}

func IntToBin(input []int) []string {
	bin := make([]string, len(input)*nbits)
	for i, n := range input {
		s := fmt.Sprintf("%0"+nbs+"b", n)
		for j, c := range s {
			bin[i*nbits+j] = string(c)
		}
	}
	return bin
}

func ToBinStr(input []int) []string {
	bs := make([]string, len(input))
	for i, n := range input {
		bs[i] = fmt.Sprintf("%0"+nbs+"b", n)
	}
	return bs
}

func MapK4() []int {
	output := make([]int, len(k4))
	p := rand.Perm(mod)
	for i, n := range k4 {
		output[i] = p[n]
	}
	return output
}

func main() {
	rand.Seed(time.Now().Unix())
	for {
		// Map K4 to other numbers (to get new bit mappings)
		newk4 = MapK4()
		//fmt.Println(newk4)

		// Try all binary encodings of a letter
		for a := 0; a < mod; a++ {
			bits := IntToBin(newk4)
			Add(newk4, 1)

			// Try all transpositions of bits
			for i := 0; i < len(bits); i++ {
				if i%97 == 0 {
					continue
				}
				// Try shifting the bits (for border cases)
				for j := 0; j < nbits; j++ {
					t := Tpose(bits, i, j)
					d := BitsToStr(t)
					// Try all transpositions of letter decodings from bits
					for k := 1; k < len(k4); k++ {
						s := strings.Join(Tpose(d, k, 0), "")
						if strings.Contains(s, "BERLI") {
							fmt.Print(s, " ", a, i, j, k)
							fmt.Println(newk4)
							//break
						}
					}
				}
			}
		}
	}
}
