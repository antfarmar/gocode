package main

import (
	"fmt"
	"strconv"
)

func fizzbuzz_expression(n int) (fizzbuzz string) {
	switch {
	case n%15 == 0:
		fizzbuzz = "FizzBuzz"
	case n%5 == 0:
		fizzbuzz = "Buzz"
	case n%3 == 0:
		fizzbuzz = "Fizz"
	default:
		fizzbuzz = strconv.Itoa(n)
	}
	return
}

func fizzbuzz_value(n int) (fizzbuzz string) {
	switch n {
	case 3, 6, 9, 12, 18:
		fizzbuzz = "Fizz"
	case 5, 10, 20:
		fizzbuzz = "Buzz"
	case 15:
		fizzbuzz = "FizzBuzz"
	default:
		fizzbuzz = strconv.Itoa(n)
	}
	return
}

func main() {
	for i := 1; i <= 20; i++ {
		v := fizzbuzz_value(i)
		e := fizzbuzz_expression(i)
		fmt.Printf("%2d: value => %s, exp => %s\n", i, v, e)
	}
}
