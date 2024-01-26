package main

import "fmt"

func findSecondLargest(arr []int) int {
	largest := arr[0]
	secondLargest := arr[0]

	for i := 1; i < len(arr); i++ {
		if arr[i] > largest {
			secondLargest = largest
			largest = arr[i]
		} else if arr[i] > secondLargest && arr[i] != largest {
			secondLargest = arr[i]
		}
	}

	return secondLargest
}

func main() {
	arr := []int{5, 18, 9, 1, 7, 3, 16}
	secondLargest := findSecondLargest(arr)
	fmt.Println("Second largest element:", secondLargest)
}
