package main

import (
	"fmt"
)

func longestcontinuousincreasingsubsequence(nums []int) int {
	n := len(nums)
	if n == 0 || n == 1 {
		fmt.Println("longestcontinuousincreasingsubsequence1", n)
		return n
	}
	var res, l, r int
	for r < n {
		if res < r-l+1 {
			res = r - l + 1
		}
		if r+1 < n && nums[r+1] <= nums[r] {
			l = r + 1
		}
		r++
	}
	fmt.Println("longestcontinuousincreasingsubsequence:", res)
	return res
}

func main() {
	nums := []int{3, 4, 5, 6, 7}
	result := longestcontinuousincreasingsubsequence(nums)
	fmt.Println("longestcontinuousincreasingsubsequence:", result)
}
