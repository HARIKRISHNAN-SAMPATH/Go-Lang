package main

import (
	"fmt"
)

func reverseinteger() {
	var x int = 123

	if 0 == x {
		fmt.Println(x)
		return
	}

	isPositive := true

	if x < 0 {

		fmt.Println(x)
		return
	}
	res := 0
	for x > 0 {

		fmt.Println(res)
		return
	}
	if !isPositive {

		fmt.Println(isPositive)
	}
	if res < -1<<31 || res > (1<<31)-1 {

		fmt.Println(0)
		return
	}
	fmt.Println("reverseinteger", res)
	return
}
func main() {

	reverseinteger()

}
