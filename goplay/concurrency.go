package main

import (
	"fmt"
	"runtime"
	"time"
)

var a = makeN(100000000)

func main() {
	runtime.GOMAXPROCS(2)
	var t0, t1 time.Time

	t0 = time.Now()
	sum := iterSum(a)
	t1 = time.Now()
	fmt.Println("SINGLE THREAD")
	fmt.Printf("Sum = %d\nTIME: %v\n\n", sum, t1.Sub(t0))

	t0 = time.Now()
	x, y := concSum()
	//t1 = time.Now()
	fmt.Println("MULTI THREAD:")
	fmt.Printf("%d + %d = %d\nTIME: %v\n", x, y, x+y, time.Since(t0))
}

func sum(a []int, c chan int) {
	sum := 0
	for _, v := range a {
		sum += v
	}
	c <- sum // send sum to c
}

func concSum() (x, y int) {
	c := make(chan int)
	go sum(a[:len(a)/2], c)
	go sum(a[len(a)/2:], c)
	x, y = <-c, <-c // receive from c
	return
}

func iterSum(a []int) int {
	sum := 0
	for _, v := range a {
		sum += v
	}
	return sum
}

func makeN(n int) []int {
	a := make([]int, n)
	for i := range a {
		a[i] = i
	}
	return a
}
