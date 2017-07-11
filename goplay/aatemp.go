package main

import (
	"fmt"
)

func main() {
	var b [10]int
	//var a = new([10]int)[:]
	a := b[:9]
	a[8] = 9
	fmt.Println("A:", a, len(a), cap(a))
	fmt.Println("B:", b, len(b), cap(b))

	a = append(a, 10)
	a[7] = 8
	b[0] = 1
	fmt.Println("A:", a, len(a), cap(a))
	fmt.Println("B:", b, len(b), cap(b))

	a = append(a, 11)
	fmt.Println("A:", a, len(a), cap(a))
	fmt.Println("B:", b, len(b), cap(b))

	a[6] = 7
	fmt.Println("A:", a, len(a), cap(a))
	fmt.Println("B:", b, len(b), cap(b))

	var m = map[int]int{2: 20}
	//m = make(map[int]int)
	m[1] = 10
	fmt.Println(m)
}
