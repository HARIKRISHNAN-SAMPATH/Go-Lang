package main

import (
	"fmt"
)

func moveZeroes1() {
	var nums [5]int = [5]int{0, 5, 7, 3, 2}
	var k int
	var i int
	for i = range nums {
		if nums[i] != 0 {
			if k != i {

				fmt.Println("moveZeroes", nums, k, i)
			}
			k++
		}
	}

}
func main() {

	moveZeroes1()

}
