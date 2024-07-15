package main

import (
	"fmt"
	"math"
)

func findMaxMin(arr []int) (int, int) {
	if len(arr) == 0 {
		// Handle the case where the array is empty
		return math.MinInt64, math.MaxInt64
	}

	max, min := arr[0], arr[0]

	for _, value := range arr[1:] {
		if value > max {
			max = value
		}
		if value < min {
			min = value
		}
	}

	return max, min
}

func main() {
	// Example array
	arr := []int{3, 5, 1, 8, -2, 7, 10, 4}
	fmt.Println("Array:", arr)
	max, min := findMaxMin(arr)
	fmt.Printf("Maximum element: %d\n", max)
	fmt.Printf("Minimum element: %d\n", min)
}
