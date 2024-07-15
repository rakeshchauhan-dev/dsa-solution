Move All Negative Elements to One Side of the Array
Problem Statement
Move all the negative elements to one side of the array.

Approach
To move all negative elements to one side of the array, we can use a two-pointer approach. This method ensures that the negative elements are moved to one side (either left or right) efficiently in a single pass through the array.

Steps:
Initialize Two Pointers:

left starts at the beginning of the array.
right starts at the end of the array.
Loop Through the Array:

Increment the left pointer until a positive element is found.
Decrement the right pointer until a negative element is found.
Swap the elements at left and right pointers if left is less than right.
Continue Until left Meets right:

Stop when left is greater than or equal to right.