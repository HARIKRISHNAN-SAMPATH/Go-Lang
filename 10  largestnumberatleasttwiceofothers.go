package main

import (
	"fmt"
	"math"
)

func largestnumberatleasttwiceofothers() {
	var nums [5]int = [5]int{5, 4, 3, 2, 1}
	var maxnum int = math.MinInt32
	var res int
	for _, j := range nums {
		if j > maxnum {

		}
	}
	for _, j := range nums {
		if maxnum < 2*j && j != maxnum {
			fmt.Println(-1)
		}
	}
	fmt.Println("largestnumberatleasttwiceofothers", res)
}
func main() {

	largestnumberatleasttwiceofothers()

}
