package main

import "fmt"

func main() {
	slice := []int{1, 2, 3, 4, 5}
	fmt.Println("slice = ", slice)
	reverseRec(slice)
	fmt.Println("Invert(slice) = ", slice)
}

func swap(a []uint8, i, j int) {
	a[j], a[i] = a[i], a[j]
}

func reverse(a []uint8) {
	for i, j := 0, len(a)-1; i < j; i, j = i+1, j-1 {
		a[i], a[j] = a[j], a[i]
	}
}

func reverseRec(slice []int) {
	length := len(slice)
	if length > 1 { // Only a slice of 2 or more elements can be inverted
		slice[0], slice[length-1] = slice[length-1], slice[0] // Swap first and last ones
		reverseRec(slice[1 : length-1])                       // Invert the slice in between
	}
}
