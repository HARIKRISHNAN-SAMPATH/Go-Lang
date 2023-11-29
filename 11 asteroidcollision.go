package main

import (
	"fmt"
)

func asteroidcollision() {
	var astroids [5]int = [5]int{5, 4, 3, 2, 1}
	stack := make([]int, 0)
	for _, astroid := range astroids {
		flag := true
		for len(stack) > 0 && astroid < 0 && stack[len(stack)-1] > 0 {
			if stack[len(stack)-1] == -astroid {
				fmt.Println("asteroidcollision1", stack)

			} else if stack[len(stack)-1] < -astroid {
				fmt.Println("asteroidcollision2", stack)
				continue
			}

			break
		}
		if flag {
			fmt.Println("asteroidcollision3", stack)

		}
	}
	fmt.Println("asteroidcollision4", stack)
	return
}

func main() {

	asteroidcollision()

}
