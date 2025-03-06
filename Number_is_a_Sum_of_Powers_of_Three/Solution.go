package main

import (
	"fmt"
)

// Function to check if a number can be represented as a sum of distinct powers of three
func checkPowersOfThree(n int) bool {
	for n > 0 {
		if n%3 == 2 {
			return false
		}
		n /= 3
	}
	return true
}

func main() {
	examples := []int{12, 91, 21}
	for _, n := range examples {
		fmt.Printf("Input: %d\nOutput: %t\n", n, checkPowersOfThree(n))
	}
}

/*
Time Complexity Analysis:
The algorithm repeatedly divides n by 3, which means it runs in O(log₃ n) time.
Since n is at most 10⁷, the number of iterations is approximately log₃(10⁷) ≈ 15, making it very efficient.

Space Complexity Analysis:
The algorithm uses only a few integer variables, meaning it runs in O(1) space (constant space).
*/
