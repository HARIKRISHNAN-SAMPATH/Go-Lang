package main

import "fmt"

func nextdifferentcharacterindex() {
	var nums [2]int = [2]int{1, 2}
	var p int = 3
	for ; p < len(nums); p++ {
		if nums[p] != nums[p-1] {
			break
		}
	}
	fmt.Println("nextdifferentcharacterindex", p)
}
func main() {

	nextdifferentcharacterindex()

}
