package main

import (
	"fmt"
)

func moveNegativesToLeft(arr []int) {
	left, right := 0, len(arr)-1

	for left <= right {
		// Increment left pointer until a positive element is found
		for left <= right && arr[left] < 0 {
			left++
		}
		// Decrement right pointer until a negative element is found
		for left <= right && arr[right] >= 0 {
			right--
		}
		// Swap elements at left and right pointers
		if left < right {
			arr[left], arr[right] = arr[right], arr[left]
			left++
			right--
		}
	}
}

func main() {
	// Example array
	arr := []int{1, -2, 3, -4, 5, -6, -7, 8, 9}
	fmt.Println("Original array:", arr)
	moveNegativesToLeft(arr)
	fmt.Println("Array after moving negatives to left:", arr)
}
