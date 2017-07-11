package main

import (
	"bytes"
	"fmt"
	"strings"
	//"time"
	//"math/rand"
)

const (
	kall = "EMUFPHZLRFAXYUSDJKZLDKRNSHGNFIVJYQTQUXQBQVYUVLLTREVJYQTMKYRDMFDVFPJUDEEHZWETZYVGWHKKQETGFQJNCEGGWHKKDQMCPFQZDQMMIAGPFXHQRLGTIMVMZJANQLVKQEDAGDVFRPJUNGEUNAQZGZLECGYUXUEENJTBJLBQCRTBJDFHRRYIZETKZEMVDUFKSJHKFWHKUWQLSZFTIHHDDDUVHDWKBFUFPWNTDFIYCUQZEREEVLDKFEZMOQQJLTTUGSYQPFEUNLAVIDXFLGGTEZFKZBSFDQVGOGIPUFXHHDRKFFHQNTGPUAECNUVPDJMQCLQUMUNEDFQELZZVRRGKFFVOEEXBDMVPNFQXEZLGREDNQFMPNZGLFLPMRJQYALMGNUVPDXVKPDQUMEBEDMHDAFMJGZNUPLGEWJLLAETGENDYAHROHNLSRHEOCPTEOIBIDYSHNAIACHTNREYULDSLLSLLNOHSNOSMRWXMNETPRNGATIHNRARPESLNNELEBLPIIACAEWMTWNDITEENRAHCTENEUDRETNHAEOETFOLSEDTIWENHAEIOYTEYQHEENCTAYCREIFTBRSPAMHHEWENATAMATEGYEERLBTEEFOASFIOTUETUAEOTOARMAEERTNRTIBSEDDNIAAHTTMSTEWPIEROAGRIEWFEBAECTDDHILCEIHSITEGOEAOSDDRYDLORITRKLMLEHAGTDHARDPNEOHMGFMFEUHEECDMRIPFEIMEHNLSSTTRTVDOHWOBKRUOXOGHULBSOLIFBBWFLRVQQPRNGKSSOTWTQSJQSSEKZZWATJKLUDIAWINFBNYPVTTMZFPKWGDKZXTJCDIGKUHUAUEKCAR"
	kabc = "KRYPTOSABCDEFGHIJLMNQUVWXZ"
	abc  = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
)

//var k1 = "EMUFPHZLRFAXYUSDJKZLDKRNSHGNFIVJYQTQUXQBQVYUVLLTREVJYQTMKYRDMFD"
//var k4 = "OBKRUOXOGHULBSOLIFBBWFLRVQQPRNGKSSOTWTQSJQSSEKZZWATJKLUDIAWINFBNYPVTTMZFPKWGDKZXTJCDIGKUHUAUEKCAR"
var k0 = kall[0:768]   // 768
var k1 = kall[0:63]    // 63
var k2 = kall[63:432]  // 369
var k3 = kall[432:768] // 336
var k4 = kall[768:865] // 97
var buf bytes.Buffer

/***************************************************************************/

func main() {
	OTPSEARCH()
}

/***************************************************************************/

func OTPSEARCH() {
	var crib []string
	key := kall
	n := 5
	//ngram(&crib, "MUYKLGKORNA", n)
	//ngram(&crib, "ELYOIECBAQK", n)

	for g := 0; g < 1; g++ {
		ngram(&crib, rotn("EYEBK", g), n)
	}
	fmt.Println(crib, len(crib))

	for h := 1; h <= len(k4); h += 1 {
		println("SIZE=", h)
		for i := 0; i <= len(key)-h; i++ {
			sub := key[i : i+h]
			for j := 1; j < len(sub); j++ {
				t := tpose(sub, j, 0)
				w := check(t, crib)
				if w != t {
					fmt.Println(w, j, h, i+1, sub)
					fmt.Println(quag3(k4, kabc, t), "<--QUAG\n")
				}
			}
		}
	}
}

func ngram(slice *[]string, word string, n int) {
	for i := 0; i <= len(word)-n; i++ {
		*slice = append(*slice, word[i:i+n])
	}
}

func rotn(s string, n int) string {
	buf.Truncate(0)
	for _, chr := range s {
		buf.WriteString(string('A' + (chr-'A'+rune(n))%26))
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

func tpose(input string, a, b int) string {
	buf.Truncate(0)
	m := len(input)
	for x := range input {
		buf.WriteString(string(input[(a*x+b)%m]))
	}
	return buf.String()
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
