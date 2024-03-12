package main

import (
	"fmt"
)

func subarrayproductlessthank2(nums []int, k int) int {

	if k <= 1 {
		fmt.Println(0)
		return 0
	}
	var (
		prod = 1
		res  = 0
		left = 0
	)

	for right, val := range nums {
		fmt.Println("bala", prod)
		prod *= val
		fmt.Println(prod)
		for prod >= k {
			prod /= nums[left]
			fmt.Println("raj:", left)
			left++
			fmt.Println(left, prod)
		}
		res += right - left + 1
		fmt.Println("hari:", res)
	}
	fmt.Println("subarrayproductlessthank2:", res)
	return res
}

func main() {
	nums := []int{4, 5, 1, 2}
	k := 70
	result := subarrayproductlessthank2(nums, k)
	fmt.Println("Number of subarrays with product less than:", k, result)
}
