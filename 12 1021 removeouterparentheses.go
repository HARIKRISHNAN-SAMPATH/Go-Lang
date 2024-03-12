package main

import "fmt"

func removeouterparentheses(s string) string {

	if len(s) == 0 {
		return s
	}
	var stacklength, left int
	var ret string
	for i := 0; i < len(s); i++ {
		if stacklength == 0 {
			left = i

		}
		if s[i] == '(' {
			stacklength++
		} else {
			stacklength--
		}
		if stacklength == 0 {
			ret += s[left+1 : i]

		}
	}

	return ret
}
func main() {
	s := "(()())(())"
	res := removeouterparentheses(s)

	fmt.Println("removeouterparentheses:", res)

}
