package main

import "fmt"

func twosum11() {
	nums := [6]int{2, 3, 5, 1, 6, 7}
	target := 12
	for i, j := range nums {
		for k := i + 1; k < len(nums); k++ {
			if nums[k]+j == target {
				fmt.Println("twosum1", nums[i], nums[k])
			}
		}
	}
	fmt.Println("twosum1", nums)
}
func main() {

	twosum11()

}
