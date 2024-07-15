package main

import (
	"container/heap"
	"fmt"
)

// MinHeap struct that implements heap.Interface
type MinHeap []int

func (h MinHeap) Len() int            { return len(h) }
func (h MinHeap) Less(i, j int) bool  { return h[i] < h[j] }
func (h MinHeap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *MinHeap) Push(x interface{}) { *h = append(*h, x.(int)) }

func (h *MinHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func sortColorsUsingHeap(arr []int) {
	h := &MinHeap{}
	heap.Init(h)

	// Build the heap by pushing all elements of the array
	for _, num := range arr {
		heap.Push(h, num)
	}

	// Extract elements from the heap and place them back into the array
	for i := 0; i < len(arr); i++ {
		arr[i] = heap.Pop(h).(int)
	}
}

func main() {
	// Example array
	arr := []int{2, 0, 2, 1, 1, 0}
	fmt.Println("Original array:", arr)
	sortColorsUsingHeap(arr)
	fmt.Println("Sorted array:", arr)
}
