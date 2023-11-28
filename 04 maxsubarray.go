package main

import (
	"fmt"
	"math"
)

func maxsubarray1() {
	nums := [10]int{3, 4, 5, -6, 8, 1, 3, -7, 5, 7}
	var (
		max           = math.MinInt32
		sum           int
		start, end, f int
	)
	for i, j := range nums {

		if sum > max {
			if f != 0 {
			}
		}
		if sum < 0 {

		}
	}
	fmt.Println("start index is %d, end index is %d\n", start, end, nums)
	return
}
func main() {

	maxsubarray1()

}
