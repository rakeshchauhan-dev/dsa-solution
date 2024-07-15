Question: Find the "Kth" Maximum and Minimum Element of an Array
Understanding the Question:

Array: A collection of elements (numbers) stored in a specific order.
Kth Minimum Element: The Kth smallest element in the array. For example, if K=3, we want to find the 3rd smallest element in the array.
Kth Maximum Element: The Kth largest element in the array. For example, if K=3, we want to find the 3rd largest element in the array.
Example:
Let's take an example array and understand the problem:

Array: [3, 5, 1, 8, -2, 7, 10, 4]

If we sort this array in ascending order, we get:

Sorted Array: [-2, 1, 3, 4, 5, 7, 8, 10]

If K=3:
The 3rd smallest element is 3 (Kth minimum).
The 3rd largest element is 7 (Kth maximum).
Steps to Solve:
Sort the Array:

First, sort the array in ascending order.
After sorting, you can easily find the Kth smallest and Kth largest elements.
Find the Kth Elements:

The Kth smallest element will be at index K-1 in the sorted array.
The Kth largest element will be at index len(array) - K in the sorted array.


**Explanation of the Code:**
Function findKthMaxMin:

Checks if K is within valid bounds.
Sorts the array in ascending order.
Finds the Kth smallest element at index k-1.
Finds the Kth largest element at index len(arr)-k.
Main Function:

Defines an example array and the value of K.
Calls the findKthMaxMin function to find and print the Kth maximum and minimum elements.
