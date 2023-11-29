package main

import (
	"fmt"
)

func twosum2() {
	var numbers [3]int = [3]int{1, 2, 3}
	var target int
	n := len(numbers)
	var (
		l = 0
		r = n - 1
	)
	for l < r {
		if numbers[l]+numbers[r] == target {
			fmt.Println("twosum2", l+1, r+1)
		} else if numbers[l]+numbers[r] < target {
			l++
		} else {
			r--
		}
	}
	fmt.Println("the input has no solution.", l+1, r+1)
}
func main() {

	twosum2()

}
