/***************************************************************
dYAhR = 01101 = duudu
BKUXOHLBOIFBFLVQPNKSOWTSQSEZZAJKUIAIFBYVTMFPWDKXJCIKUUUECR
OROGUSLBWRQRGSTQJSKWTLDWNNPTZKGZTDGHAKA

BERLINCLOCK:    BRNLK  ELICOC
NYPVTTMZFPK:    NPTZK  YVTMFP
ELYOIECBAQK:    EYEBK  LOICAQ   (KRYPTOSABCD)
MUYKLGKORNA:    MYGOA  UKLKRN   (ABCDEFGHIJK)
EUYKLEKBRNK
MZULFRISMMT:    MURST  ZLFIMM   (EMUFPHZLRFA)
***************************************************************/
package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"math/rand"
	"regexp"
	"strings"
	"time"
)

const (
	kall  = "EMUFPHZLRFAXYUSDJKZLDKRNSHGNFIVJYQTQUXQBQVYUVLLTREVJYQTMKYRDMFDVFPJUDEEHZWETZYVGWHKKQETGFQJNCEGGWHKKDQMCPFQZDQMMIAGPFXHQRLGTIMVMZJANQLVKQEDAGDVFRPJUNGEUNAQZGZLECGYUXUEENJTBJLBQCRTBJDFHRRYIZETKZEMVDUFKSJHKFWHKUWQLSZFTIHHDDDUVHDWKBFUFPWNTDFIYCUQZEREEVLDKFEZMOQQJLTTUGSYQPFEUNLAVIDXFLGGTEZFKZBSFDQVGOGIPUFXHHDRKFFHQNTGPUAECNUVPDJMQCLQUMUNEDFQELZZVRRGKFFVOEEXBDMVPNFQXEZLGREDNQFMPNZGLFLPMRJQYALMGNUVPDXVKPDQUMEBEDMHDAFMJGZNUPLGEWJLLAETGENDYAHROHNLSRHEOCPTEOIBIDYSHNAIACHTNREYULDSLLSLLNOHSNOSMRWXMNETPRNGATIHNRARPESLNNELEBLPIIACAEWMTWNDITEENRAHCTENEUDRETNHAEOETFOLSEDTIWENHAEIOYTEYQHEENCTAYCREIFTBRSPAMHHEWENATAMATEGYEERLBTEEFOASFIOTUETUAEOTOARMAEERTNRTIBSEDDNIAAHTTMSTEWPIEROAGRIEWFEBAECTDDHILCEIHSITEGOEAOSDDRYDLORITRKLMLEHAGTDHARDPNEOHMGFMFEUHEECDMRIPFEIMEHNLSSTTRTVDOHWOBKRUOXOGHULBSOLIFBBWFLRVQQPRNGKSSOTWTQSJQSSEKZZWATJKLUDIAWINFBNYPVTTMZFPKWGDKZXTJCDIGKUHUAUEKCAR"
	k0d   = "BETWEENSUBTLESHADINGANDTHEABSENCEOFLIGHTLIESTHENUANCEOFIQLUSIONITWASTOTALLYINVISIBLEHOWSTHATPOSSIBLETHEYUSEDTHEEARTHSMAGNETICFIELDXTHEINFORMATIONWASGATHEREDANDTRANSMITTEDUNDERGRUUNDTOANUNKNOWNLOCATIONXDOESLANGLEYKNOWABOUTTHISTHEYSHOULDITSBURIEDOUTTHERESOMEWHEREXWHOKNOWSTHEEXACTLOCATIONONLYWWTHISWASHISLASTMESSAGEXTHIRTYEIGHTDEGREESFIFTYSEVENMINUTESSIXPOINTFIVESECONDSNORTHSEVENTYSEVENDEGREESEIGHTMINUTESFORTYFOURSECONDSWESTXLAYERTWOSLOWLYDESPARATLYSLOWLYTHEREMAINSOFPASSAGEDEBRISTHATENCUMBEREDTHELOWERPARTOFTHEDOORWAYWASREMOVEDWITHTREMBLINGHANDSIMADEATINYBREACHINTHEUPPERLEFTHANDCORNERANDTHENWIDENINGTHEHOLEALITTLEIINSERTEDTHECANDLEANDPEEREDINTHEHOTAIRESCAPINGFROMTHECHAMBERCAUSEDTHEFLAMETOFLICKERBUTPRESENTLYDETAILSOFTHEROOMWITHINEMERGEDFROMTHEMISTXCANYOUSEEANYTHINGQ"
	k439  = "OROGUSLBWRQRGSTQJSKWTLDWNNPTZKGZTDGHAKA"                    // 39
	k458  = "BKUXOHLBOIFBFLVQPNKSOWTSQSEZZAJKUIAIFBYVTMFPWDKXJCIKUUUECR" // 58
	k4ob  = "OROGUSLBWRQRGSTQJSKWTLDWNNPTZKGZTDGHAKABKUXOHLBOIFBFLVQPNKSOWTSQSEZZAJKUIAIFBYVTMFPWDKXJCIKUUUECR"
	k4bo  = "BKUXOHLBOIFBFLVQPNKSOWTSQSEZZAJKUIAIFBYVTMFPWDKXJCIKUUUECROROGUSLBWRQRGSTQJSKWTLDWNNPTZKGZTDGHAKA"
	kabc  = "KRYPTOSABCDEFGHIJLMNQUVWXZ"
	abc   = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	k4cnt = "45232442438413534466625214"
)

var k0 = kall[0:768]   // 768
var k1 = kall[0:63]    // 63
var k2 = kall[63:432]  // 369
var k3 = kall[432:768] // 336
var k4 = kall[768:865] // 97

var k1d = k0d[0:63]    // 63
var k2d = k0d[63:433]  // 370
var k3d = k0d[433:769] // 336

//var rePT = regexp.MustCompile(`B[AEIOUY]RL[AEIOUY]N[CQK]L[AEIOUY][CQK][CQK]`)
var rePT = regexp.MustCompile("B[A-Z]RL[A-Z]N[CQK]L[A-Z][CQK][CQK]")
var reKY = regexp.MustCompile("E[ULGWAK]YO[WNIZCY]E[CMX]B[ORWALD][QPC][LKS]")
var buf bytes.Buffer

/***************************************************************************/

func main() {
	//println(countchrs(k4))
	//pt:="BERLINCLOCK"
	//ct:="NYPVTTMZFPK"
	//for i := 0; i < len(k4); i++ {
	//print(matrix(string(k0[i]) + string(k3[i])))
	//}
	//println(matrix(k0pt))
	OTPSEARCH()
	//DICT()
}

/***************************************************************************/

func OTPSEARCH() {
	crib := []string{"SYMLWR", "LRCFPL", "MXFEVL", "VNMWWA", "PJYVLP", "UGFVVK", "ICXUEI", "HYSSRS", "RGTGNW","UCJFBF","QSQAFL","ELYOIE"}
	//n := 5
	//ngram(&crib, "BERLINCLOCK", n)
	//fmt.Println(crib, len(crib))
	key := k0

	for h := 6; h <= len(k4); h += 1 {
		println("SIZE =", h)
		for i := 0; i <= len(key)-h; i++ {
			sub := key[i : i+h]
			for j := 1; j < len(sub); j++ {
				for k := 0; k < 6; k++ {
					t := tpose(sub, j, k)
					//q := quag3(k4, kabc, t)
					//q := quag3(k4ob, kabc, t)
					//s := join(q[:39], q[39:])
					//q := quag3(k4bo, kabc, t)
					//s := join(q[58:], q[:58])
					//q1 := quag3(k439, kabc, t)
					//q2 := quag3(k458, kabc, t)
					//s := join(q1, q2)
					//w := t[63:74]
					//m := check64to75(w)
					//s1, s2 := split(t)
					//s := s1 + s2
					//if rePT.MatchString(q) {
					//fmt.Println(q, j, k, h, i+1, sub, t)
					//}
					c := check(t, crib)
					if c != "" {
						fmt.Println(c, j, k, h, i+1, sub)
					}
				}
			}
		}
	}
}

func DICT() {
	rand.Seed(time.Now().Unix())
	dict := getDict("dict_phoenix.txt", 0)
	println(len(dict))
	for i := 0; i < len(dict); i++ {
		r := rand.Intn(len(dict))
		alph := makeABC(dict[r])
		//println(alph)
		for j := 0; j < len(dict); j++ {
			//r1, r2 := rand.Intn(len(dict)), rand.Intn(len(dict))
			word := dict[j]
			q := quag3(k4, alph, word)
			//s := join(q[:39], q[39:])
			//q := quag3(k4ob, kabc, word)
			//s := join(q[:39], q[39:])
			//q := quag3(k4bo, kabc, word)
			//s := join(q[58:], q[:58])
			//w := q[63:74]
			//m := check64to75(w)
			if rePT.MatchString(q) {
				fmt.Println(q, dict[r], word)
			}
		}
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

func makeABC(word string) string {
	buf.Truncate(0)
	for _, c := range word {
		if !strings.Contains(buf.String(), string(c)) {
			buf.WriteString(string(c))
		}
	}
	for _, c := range abc {
		if !strings.Contains(buf.String(), string(c)) {
			buf.WriteString(string(c))
		}
	}
	return buf.String()
}

func join(s39, s58 string) string {
	buf.Truncate(0)
	b := "01101"
	l1, l2 := len(s39), len(s58)
	var m, n int
	for i := 0; i < l1+l2; i++ {
		if b[i%5] == '0' {
			buf.WriteString(string(s39[m]))
			m += 1
		} else {
			buf.WriteString(string(s58[n]))
			n += 1
		}
	}
	return buf.String()
}

func split(s string) (string, string) {
	var buf1, buf2 bytes.Buffer
	b := "01101"
	for i, chr := range s {
		if b[i%5] == '0' {
			buf1.WriteString(string(chr))
		} else {
			buf2.WriteString(string(chr))
		}
	}
	return buf1.String(), buf2.String()
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

func check64to75(s string) int {
	n := 0
	bc := "BERLINCLOCK"
	for i := range s {
		if s[i] == bc[i] {
			n += 1
		}
	}
	return n
}

func check(s string, words []string) string {
	out := ""
	for _, w := range words {
		if strings.Contains(s, w) {
			out = strings.Replace(s, w, "["+w+"]", -1)
			break
		}
	}
	return out
}

func countchrs(s string) string {
	buf.Truncate(0)
	for _, c := range abc {
		n := strings.Count(s, string(c))
		buf.WriteString(string(n + '0'))
	}
	return buf.String()
}

func matrix(ct string) string {
	buf.Truncate(0)
	for i := 0; i < len(ct)-1; i += 2 {
		m := strings.Index(abc, string(ct[i]))
		n := strings.Index(abc, string(ct[i+1]))
		buf.WriteString(string(kabc[(m+n)%26]))
	}
	return buf.String()
}

func unvowel(word string) string {
	re := regexp.MustCompile("[AEIOU]")
	return re.ReplaceAllString(word, "")
	//buf.Truncate(0)
	//for _, c := range word {
	//vwl := false
	//for _, v := range "AEIOU" {
	//if c == v {
	//vwl = true
	//break
	//}
	//}
	//if !vwl {
	//buf.WriteString(string(c))
	//}
	//}
	//return buf.String()
}

func getDict(fi string, lett int) []string {
	b, err := ioutil.ReadFile(fi)
	if err != nil {
		panic(err)
	}
	words := strings.Split(string(b), "\n")
	if lett > 0 {
		var neww []string
		for _, w := range words {
			if len(w) >= lett {
				neww = append(neww, w)
			}
		}
		return neww
	}
	return words
}
