package main

import "fmt"

func removeouterparentheses() {
	var s string
	if len(s) == 0 {
		fmt.Println(s)
	}
	var stacklength, left int
	var ret string
	for i := 0; i < len(s); i++ {
		if stacklength == 0 {
			fmt.Println("removeouterparentheses", left)

		}
		if s[i] == '(' {
			stacklength++
		} else {
			stacklength--
		}
		if stacklength == 0 {

		}
	}
	fmt.Println("removeouterparentheses", ret)
}
func main() {

	removeouterparentheses()

}
