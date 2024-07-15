package main

import (
	"fmt"
	"sort"
)

func findKthMaxMin(arr []int, k int) (int, int, error) {
	if k <= 0 || k > len(arr) {
		return 0, 0, fmt.Errorf("k is out of bounds")
	}

	// Sort the array
	sort.Ints(arr)

	// Kth minimum is at index k-1
	kthMin := arr[k-1]

	// Kth maximum is at index len(arr)-k
	kthMax := arr[len(arr)-k]

	return kthMax, kthMin, nil
}

func main() {
	// Example array
	arr := []int{3, 5, 1, 8, -2, 7, 10, 4}
	k := 3
	fmt.Println("Array:", arr)
	kthMax, kthMin, err := findKthMaxMin(arr, k)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Printf("%dth maximum element: %d\n", k, kthMax)
		fmt.Printf("%dth minimum element: %d\n", k, kthMin)
	}
}
