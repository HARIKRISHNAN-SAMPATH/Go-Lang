package main

import "fmt"

func RArray(nums []int, start, end int) {
	for start < end {
		nums[start], nums[end] = nums[end], nums[start]
		start++
		end--
	}
}
func main() {
	nums := []int{1, 2, 34, 4, 5, 6, 7, 7, 8, 3, 5}
	start := 0
	end := 10
	RArray(nums, start, end)
	fmt.Println("Reverse Array:", nums)
}
