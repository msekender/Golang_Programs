/*
Solution: Container With Most Water (Two-Pointer Approach)
The problem can be solved efficiently using the two-pointer approach instead of brute force. The key observation is that moving the shorter height pointer inward has a chance to find a larger area, while moving the taller height pointer does not help.
*/
package main

import (
	"fmt"
)

func maxArea(height []int) int {
	left, right := 0, len(height)-1
	maxArea := 0

	for left < right {
		// Calculate current area
		h := min(height[left], height[right])
		area := (right - left) * h
		if area > maxArea {
			maxArea = area
		}

		// Move the pointer pointing to the shorter height
		if height[left] < height[right] {
			left++
		} else {
			right--
		}
	}
	return maxArea
}

// Helper function to find minimum of two numbers
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	// Example test cases
	fmt.Println(maxArea([]int{1, 8, 6, 2, 5, 4, 8, 3, 7})) // Output: 49
	fmt.Println(maxArea([]int{1, 1}))                      // Output: 1
}

/*
Time Complexity Analysis
We only traverse the array once using two pointers.
Each iteration takes constant time (O(1) operations).
Total complexity: O(n) where n is the number of elements in height.

Space Complexity Analysis
We only use a few integer variables (left, right, maxArea).
No extra arrays or data structures are used.
Total space complexity: O(1).

*/
