// aatemp
package main

import (
	"fmt"
)

type Iter chan interface{}

func NewIter(els ...interface{}) Iter {
	c := make(Iter)
	go func() {
		for _, el := range els {
			c <- el
		}
		close(c)
	}()
	return c
}

// Count from i to infinity
func Count(i int) Iter {
	c := make(Iter)
	go func() {
		for {
			c <- i
			i++
		}
	}()
	return c
}

func CountStep(start, step, end int) Iter {
	c := make(Iter)
	go func() {
		for start <= end {
			c <- start
			start += step
		}
		close(c)
	}()
	return c
}

// Cycle through an iterator infinitely (requires memory)
func Cycle(it Iter) Iter {
	c, a := make(Iter), make([]interface{}, 0, 1)
	go func() {
		for el := range it {
			a = append(a, el)
			c <- el
		}
		for {
			for _, el := range a {
				c <- el
			}
		}
	}()
	return c
}

func main() {
	iter := NewIter("first", 22222, 3.141)
	cntIter := Count(10)
	cntStep := CountStep(10, 5, 100)
	cycIter := Cycle(NewIter("even", "odd"))
	for i := 0; i <= 10; i++ {
		fmt.Println(<-cntIter, <-cntStep, <-iter, <-cycIter)
	}
}
