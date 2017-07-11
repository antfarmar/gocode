// slicedelete
package main

import "fmt"

// append is a variadic function ...
func delete(i int, slice []int) []int {
	slice = append(slice[:i], slice[i+1:]...)
	return slice
}

func main() {
	slice := []int{1, 2, 3, 4, 5, 6, 7}
	fmt.Println("In the beginning...")
	fmt.Println("slice = ", slice)
	slice = append(slice, 8, 9, 10)
	fmt.Println("A variadic append...")
	fmt.Println("slice = ", slice)
	fmt.Println("Let's delete the first element")
	slice = delete(0, slice)
	fmt.Println("slice = ", slice)
	fmt.Println("Let's delete the last element")
	slice = delete(len(slice)-1, slice)
	fmt.Println("slice = ", slice)
	fmt.Println("Let's delete the 3rd element")
	slice = delete(2, slice)
	fmt.Println("slice = ", slice)
}
