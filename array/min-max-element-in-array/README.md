Find the maximum and minimum element in an array

Explanation:
Function findMaxMin(arr []int) (int, int):

This function takes an array of integers as input and returns two integers: the maximum and minimum elements in the array.
It first checks if the array is empty and handles this case by returning the minimum and maximum possible integers.
It initializes max and min with the first element of the array.
It iterates through the rest of the array, updating max and min whenever a larger or smaller element is found.
Main Function:

An example array is defined.
The array is printed.
The findMaxMin function is called to find the maximum and minimum elements.
The maximum and minimum elements are printed.
