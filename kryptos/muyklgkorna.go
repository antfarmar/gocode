package main

import (
	"bytes"
	"fmt"
	"math/rand"
	"strings"
	"time"
)

var k4m = "WGDKZXTJCDIGKUHUAUEKCAROBKRUOXOGHULBSOLIFBBWFLRVQQPRNGKSSOTWTQSJQSSEKZZWATJKLUDIAWINFBNYPVTTMZFPK"
var otp = "AAABBBBBCCDDDEEFFFFGGGHHIIIIJJJKKKKKKLLLNNOOOOPPPQQQQRRRSSSSSSTTTTTTUUUUUVVWWWWWXXZZZZMUYKLGKORNA"

//var k4s = "AAAABBBBBCCDDDEEFFFFGGGGHHIIIIJJJKKKKKKKKLLLLMNNNOOOOOPPPQQQQRRRRSSSSSSTTTTTTUUUUUUVVWWWWWXXYZZZZ"
var kabc = "KRYPTOSABCDEFGHIJLMNQUVWXZ"
var abc = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
var buf bytes.Buffer

/***************************************************************************/

func main() {
	//sl := strings.Split(otp, "")
	//sort.Sort(sort.StringSlice(sl))
	rand.Seed(time.Now().Unix())
	crib := []string{"KRYPTOS"}

	for {
		p := shuffle(otp[:86])
		pt := quag3(k4m, abc, p)
		w := check(pt, crib)
		if w != pt {
			fmt.Println(w, p)
		}
	}
}

/***************************************************************************/

func shuffle(s string) string {
	buf.Truncate(0)
	r := rand.Perm(len(s))
	for _, n := range r {
		buf.WriteString(string(s[n]))
	}
	buf.WriteString(otp[86:])
	return buf.String()
}

func ngram(slice *[]string, word string, n int) {
	for i := 0; i <= len(word)-n; i++ {
		*slice = append(*slice, word[i:i+n])
	}
}

func tpose(input string, a, b int) string {
	buf.Truncate(0)
	m := len(input)
	for x := range input {
		buf.WriteString(string(input[(a*x+b)%m]))
	}
	return buf.String()
}

func check(s string, words []string) string {
	for _, w := range words {
		if strings.Contains(s, w) {
			s = strings.Replace(s, w, "["+w+"]", -1)
		}
	}
	return s
}

func quag3(ciph, alph, key string) string {
	buf.Truncate(0)
	for i, c := range ciph {
		ltr := string(key[i%len(key)])
		n := strings.Index(alph, ltr)
		m := strings.Index(alph, string(c))
		diff := m - n
		if diff < 0 {
			diff = diff + 26
		}
		buf.WriteString(string(alph[diff]))
	}
	return buf.String()
}
