Certainly! The Dutch National Flag algorithm is a well-known algorithm for sorting an array consisting of only three distinct values (0, 1, and 2 in this case) in linear time. It was proposed by Edsger Dijkstra. The idea is to partition the array into three sections:

Elements less than the middle value (0 in this case).
Elements equal to the middle value (1 in this case).
Elements greater than the middle value (2 in this case).

Steps of the Algorithm:
Initialize three pointers:

low to track the boundary of elements less than the middle value.

mid to track the current element being processed.

high to track the boundary of elements greater than the middle value.

Iterate through the array with the mid pointer:

If the current element is 0, swap it with the element at the low pointer, and move both low and mid pointers one step forward.

If the current element is 1, just move the mid pointer one step forward.

If the current element is 2, swap it with the element at the high pointer, and move the high pointer one step backward.

Stop the iteration when mid pointer exceeds the high pointer.

**Detailed Algorithm:**

Initialization:

low = 0
mid = 0
high = len(array) - 1
Loop until mid <= high:

If array[mid] == 0:
Swap array[low] and array[mid]
Increment both low and mid
If array[mid] == 1:
Increment mid
If array[mid] == 2:
Swap array[mid] and array[high]
Decrement high

**Example:**
For an array [2, 0, 2, 1, 1, 0]:

Initial state: [2, 0, 2, 1, 1, 0], low=0, mid=0, high=5

After processing first element (2): [0, 0, 2, 1, 1, 2], low=0, mid=0, high=4

After processing second element (0): [0, 0, 2, 1, 1, 2], low=1, mid=1, high=4

After processing third element (0): [0, 0, 2, 1, 1, 2], low=2, mid=2, high=4

After processing fourth element (2): [0, 0, 1, 1, 2, 2], low=2, mid=2, high=3

After processing fifth element (1): [0, 0, 1, 1, 2, 2], low=2, mid=3, high=3

After processing sixth element (1): [0, 0, 1, 1, 2, 2], low=2, mid=4, high=3

Final state: [0, 0, 1, 1, 2, 2]