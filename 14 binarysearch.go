package main

import (
	"fmt"
)

func search(nums []int, target int) int {
	var (
		l int
		r = len(nums) - 1
	)

	for l <= r {
		mid := l + (r-l)/2
		if nums[mid] == target {
			return mid
		}
		if nums[mid] > target {
			r = mid - 1
		} else {
			l = mid + 1
		}
	}
	fmt.Println(-1)
	return -1
}

func main() {

	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	target := 7
	result := search(nums, target)
	fmt.Println("binary search:", result)

}
