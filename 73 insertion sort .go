package main

import "fmt"

func insertionSort(arr []int) {
	for i := 1; i < len(arr); i++ {
		key := arr[i]
		j := i - 1

		for j >= 0 && arr[j] > key {
			arr[j+1] = arr[j]
			j = j - 1
		}
		arr[j+1] = key
	}
}

func main() {
	arr := []int{5, 2, 4, 6, 1, 3}
	fmt.Println("Unsorted array:", arr)

	insertionSort(arr)
	fmt.Println("Sorted array:", arr)
}
