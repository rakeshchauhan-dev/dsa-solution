package main

import (
	"fmt"
)

func sortColors(arr []int) {
	low, mid, high := 0, 0, len(arr)-1

	for mid <= high {
		switch arr[mid] {
		case 0:
			arr[low], arr[mid] = arr[mid], arr[low]
			low++
			mid++
		case 1:
			mid++
		case 2:
			arr[mid], arr[high] = arr[high], arr[mid]
			high--
		}
	}
}

func main() {
	// Example array
	arr := []int{2, 0, 2, 1, 1, 0}
	fmt.Println("Original array:", arr)
	sortColors(arr)
	fmt.Println("Sorted array:", arr)
}
