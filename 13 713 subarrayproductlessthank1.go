package main

import (
	"fmt"
)

func numSubArrayProductLessThanK(nums []int, k int) int {
	if k <= 1 {
		return 0
	}

	var (
		n    = len(nums)
		l, r int
		res  int
		prod = 1
	)

	for l < n {
		if r < n && prod*nums[r] < k {
			prod *= nums[r]
			r++
		} else if l == r {
			l++
			r++
		} else {
			res += r - l
			prod /= nums[l]
			l++
		}
	}
	return res
}

func main() {

	nums := []int{10, 5, 3, 6}
	k := 100
	result := numSubArrayProductLessThanK(nums, k)
	fmt.Println("Result:", result)
}
