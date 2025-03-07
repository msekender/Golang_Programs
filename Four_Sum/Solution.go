/*
Approach
Sort the input array → Helps in efficient duplicate handling.
Fix the first two numbers (i and j) using nested loops.
Use the two-pointer approach for the remaining two numbers (c and d):
If the sum of all four numbers is equal to the target, add it to the result.
If the sum is too small, move the left pointer (c) forward.
If the sum is too large, move the right pointer (d) backward.
Skip duplicate numbers to ensure unique quadruplets.
*/

package main

import (
	"fmt"
	"sort"
)

func fourSum(nums []int, target int) [][]int {
	sort.Ints(nums) // Step 1: Sort the array
	n := len(nums)
	var result [][]int

	for i := 0; i < n-3; i++ {
		// Skip duplicate i
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		for j := i + 1; j < n-2; j++ {
			// Skip duplicate j
			if j > i+1 && nums[j] == nums[j-1] {
				continue
			}

			// Two-pointer approach
			c, d := j+1, n-1
			for c < d {
				sum := nums[i] + nums[j] + nums[c] + nums[d]
				if sum == target {
					result = append(result, []int{nums[i], nums[j], nums[c], nums[d]})

					// Move c and avoid duplicates
					for c < d && nums[c] == nums[c+1] {
						c++
					}
					// Move d and avoid duplicates
					for c < d && nums[d] == nums[d-1] {
						d--
					}

					c++
					d--
				} else if sum < target {
					c++ // Increase sum
				} else {
					d-- // Decrease sum
				}
			}
		}
	}
	return result
}

func main() {
	// Example test cases
	fmt.Println(fourSum([]int{1, 0, -1, 0, -2, 2}, 0)) // Output: [[-2,-1,1,2],[-2,0,0,2],[-1,0,0,1]]
	fmt.Println(fourSum([]int{2, 2, 2, 2, 2}, 8))      // Output: [[2,2,2,2]]
}

/*
Time Complexity Analysis
Sorting the array: O(nlogn)
Iterating over pairs (i, j): O(n*n)
Two-pointer approach per pair: O(n)
Total complexity: O(n*n*n)
This is significantly faster than brute-force O(n*n*n*n)

Space Complexity Analysis
Sorting takes 𝑂(1)
O(1) additional space (in-place sorting).
Result list stores valid quadruplets, worst case O(n*n)
Overall Space Complexity: O(n*n)
*/
