package main

import (
	"fmt"
	str "strings"
)

func main() {
	abc := "TAGC"
	n := 2
	p := prd(str.Split(abc, ""), n) // [T A G C]
	fmt.Println(str.Join(p, "\n"))
	//prdrec(abc, "", 0, n)
}

func prd(abc []string, n int) (nary []string) {
	nary = abc
	for ; n > 1; n-- {
		temp := []string{}
		for _, s := range abc {
			for _, t := range nary {
				temp = append(temp, s+t)
			}
		}
		nary = temp
	}
	return
}

func prdrec(abc, s string, i, n int) {
	if i == n {
		fmt.Printf("%s\n", s)
	} else {
		for _, c := range abc {
			prdrec(abc, s+string(c), i+1, n)
		}
	}
}
