Using Heaps to Sort 0s, 1s, and 2s
A heap is a binary tree-based data structure that satisfies the heap property:

Min-Heap: The parent node is always smaller than its child nodes.
Max-Heap: The parent node is always larger than its child nodes.
For this problem, we can use a Min-Heap to extract the smallest elements repeatedly until the array is sorted. Here’s how you can do it:

Steps:
Build a Min-Heap from the array.
Extract the elements from the heap one by one.
Detailed Algorithm:
Build a Min-Heap:

Insert all elements of the array into a Min-Heap.
Extract Elements:

Repeatedly extract the minimum element from the heap and place it back into the array.
Complexity:


Explanation of the Code:


MinHeap Struct:

Defines a type MinHeap which is a slice of integers.
Implements heap.Interface methods: Len, Less, Swap, Push, and Pop.
sortColorsUsingHeap Function:

Initializes a MinHeap.
Pushes all elements of the array into the heap.
Pops elements from the heap and places them back into the array in sorted order.
Main Function:

Defines an example array.
Calls the sortColorsUsingHeap function to sort the array.
Prints the original and sorted arrays.
How It Works:
Building the Min-Heap:

The elements [2, 0, 2, 1, 1, 0] are pushed into the heap one by one.
The heap maintains the smallest element at the root.
Extracting Elements:

The smallest element (0) is repeatedly extracted and placed back into the array.
This continues until all elements are extracted and the array is sorted.