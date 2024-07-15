package main

import "fmt"

func main() {
	arr := []int{1, 2, 3, 4, 5}
	fmt.Println("Original Array:", arr)
	reversed := reverseArray(arr)
	fmt.Println("Reversed Array:", reversed)
}

func reverseArray(arr []int) []int {
	tempArray := make([]int, len(arr)) // Initialize with the same length
	for i := range arr {
		tempArray[i] = arr[len(arr)-1-i]
	}
	return tempArray
}
