package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

const (
	k3e = "ENDYAHROHNLSRHEOCPTEOIBIDYSHNAIACHTNREYULDSLLSLLNOHSNOSMRWXMNETPRNGATIHNRARPESLNNELEBLPIIACAEWMTWNDITEENRAHCTENEUDRETNHAEOETFOLSEDTIWENHAEIOYTEYQHEENCTAYCREIFTBRSPAMHHEWENATAMATEGYEERLBTEEFOASFIOTUETUAEOTOARMAEERTNRTIBSEDDNIAAHTTMSTEWPIEROAGRIEWFEBAECTDDHILCEIHSITEGOEAOSDDRYDLORITRKLMLEHAGTDHARDPNEOHMGFMFEUHEECDMRIPFEIMEHNLSSTTRTVDOHW"
	k3d = "SLOWLYDESPARATLYSLOWLYTHEREMAINSOFPASSAGEDEBRISTHATENCUMBEREDTHELOWERPARTOFTHEDOORWAYWASREMOVEDWITHTREMBLINGHANDSIMADEATINYBREACHINTHEUPPERLEFTHANDCORNERANDTHENWIDENINGTHEHOLEALITTLEIINSERTEDTHECANDLEANDPEEREDINTHEHOTAIRESCAPINGFROMTHECHAMBERCAUSEDTHEFLAMETOFLICKERBUTPRESENTLYDETAILSOFTHEROOMWITHINEMERGEDFROMTHEMISTXCANYOUSEEANYTHINGQ"
	k4e = "OBKRUOXOGHULBSOLIFBBWFLRVQQPRNGKSSOTWTQSJQSSEKZZWATJKLUDIAWINFBNYPVTTMZFPKWGDKZXTJCDIGKUHUAUEKCAR"
)

var k4s = []int{14, 1, 10, 17, 20, 14, 23, 14, 6, 7, 20, 11, 1, 18, 14, 11, 8, 5, 1, 1, 22, 5, 11, 17, 21, 16, 16, 15, 17, 13, 6, 10, 18, 18, 14, 19, 22, 19, 16, 18, 9, 16, 18, 18, 4, 10, 25, 25, 22, 0, 19, 9, 10, 11, 20, 3, 8, 0, 22, 8, 13, 5, 1, 13, 24, 15, 21, 19, 19, 12, 25, 5, 15, 10, 22, 6, 3, 10, 25, 23, 19, 9, 2, 3, 8, 6, 10, 20, 7, 20, 0, 20, 4, 10, 2, 0, 17}

var key = "BERLIN"

func permk3() {
	//rand.Seed(time.Now().Unix())
	for {
		se, sd := "", ""
		a := rand.Perm(len(k3e)) //336
		a = a[:97]
		for i := range a {
			se += string(k3e[a[i]])
			sd += string(k3d[a[i]])
		}

		if strings.Contains(se, key) {
			x := strings.Index(se, key)
			print(x, " ")
			println("se:", se)
			fmt.Println(a, "\n", a[x:x+len(key)], "\n")
		}

		if strings.Contains(sd, key) {
			x := strings.Index(sd, key)
			print(x, " ")
			println("sd:", sd)
			fmt.Println(a, "\n", a[x:x+len(key)], "\n")
		}
	}
}

///////////////////////////////////////////////////////
func permk4() {
	for {
		r := rand.Perm(len(k4e))
		se := ""

		for i := range r {
			se += string(k4e[r[i]])
		}

		if strings.Contains(se, key) {
			x := strings.Index(se, key)
			print(x, " ")
			println("se:", se)
			fmt.Println(r, "\n", r[x:x+len(key)], "\n")
		}
	}
}

//////////////////////////////////////////
func shuffle(s string, n int) []int {
	r := rand.Perm(len(s))
	t := make([]int, len(s))
	for i := range r {
		t[i] = int(k4e[r[i]]-65) + n
	}
	return t
}

///////////////////////////////////////////

func main() {
	rand.Seed(time.Now().Unix())
	permk4()
}
