/*
Approach
To compute the minimum number of operations efficiently, we can use a two-pass approach:

Left-to-Right Pass:
Keep track of the number of balls and their cumulative moves to reach each box from the left.
Right-to-Left Pass:
Compute the total operations by incorporating contributions from the right side.
*/

package main

import (
	"fmt"
)

func minOperations(boxes string) []int {
	n := len(boxes)
	answer := make([]int, n)
	leftCount, leftMoves := 0, 0

	// Left-to-right pass
	for i := 0; i < n; i++ {
		answer[i] = leftMoves
		if boxes[i] == '1' {
			leftCount++
		}
		leftMoves += leftCount
	}

	rightCount, rightMoves := 0, 0

	// Right-to-left pass
	for i := n - 1; i >= 0; i-- {
		answer[i] += rightMoves
		if boxes[i] == '1' {
			rightCount++
		}
		rightMoves += rightCount
	}

	return answer
}

func main() {
	boxes := "110"
	fmt.Println(minOperations(boxes)) // Output: [1,1,3]

	boxes2 := "001011"
	fmt.Println(minOperations(boxes2)) // Output: [11,8,5,4,3,4]
}

/*
Time Complexity Analysis
First pass (Left-to-Right):
We iterate through the string once (O(n)) and compute the cumulative moves.
Second pass (Right-to-Left):
Again, a single pass (O(n)) updates the results.
Overall, the solution runs in O(n) time complexity, which is efficient given n ≤ 2000.

Space Complexity Analysis
We use a single answer array of size n, so the space complexity is O(n).
Other variables (leftCount, leftMoves, rightCount, rightMoves) use O(1) space.
Thus, the overall space complexity is O(n) due to the output array.
*/
