package main

import (
	"fmt"
)

func main() {
	nums := []int{5, 8, 1, 3, 7, 9, 2}
	fmt.Println(qsort(nums))
}

func qsort(nums []int) (srtd []int) {
	pivot := nums[0]
	low, high := partition(pivot, nums)
	if len(low) > 1 {
		low = qsort(low)
	}
	if len(high) > 1 {
		high = qsort(high)
	}

	srtd = append(low, pivot)
	srtd = append(srtd, high...)
	return
}

func partition(pivot int, nums []int) (low, high []int) {
	for _, n := range nums[1:] {
		if n < pivot {
			low = append(low, n)
		} else {
			high = append(high, n)
		}
	}
	return
}
