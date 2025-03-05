package main

import (
	"fmt"
)

// Function to find the longest palindromic substring
func longestPalindrome(s string) string {
	if len(s) == 0 {
		return ""
	}

	start, end := 0, 0

	for i := 0; i < len(s); i++ {
		len1 := expandAroundCenter(s, i, i)   // Odd length palindrome
		len2 := expandAroundCenter(s, i, i+1) // Even length palindrome
		maxLen := max(len1, len2)

		if maxLen > end-start {
			start = i - (maxLen-1)/2
			end = i + maxLen/2
		}
	}

	return s[start : end+1]
}

// Helper function to expand around center and return palindrome length
func expandAroundCenter(s string, left, right int) int {
	for left >= 0 && right < len(s) && s[left] == s[right] {
		left--
		right++
	}
	return right - left - 1
}

// Helper function to find the max of two numbers
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	s1 := "babad"
	fmt.Println("Input:", s1)
	fmt.Println("Output:", longestPalindrome(s1))

	s2 := "cbbd"
	fmt.Println("Input:", s2)
	fmt.Println("Output:", longestPalindrome(s2))
}

/*
Explanation:
Expand Around Center: We check palindromes by expanding from each character and from each pair of consecutive characters.
Tracking Longest Palindrome: The start and end indices of the longest found palindrome are updated.

Time Complexity: O(n²) since we expand for each character.

The space complexity of the provided expand around center approach for finding the longest palindromic substring is O(1).
Why is the Space Complexity O(1)?
The algorithm does not use any extra data structures (like a DP table or additional arrays) that scale with the input size.
It only uses a few integer variables (start, end, len1, len2, maxLen, left, and right), which occupy constant space.
The function expandAroundCenter only uses local integer variables and does not allocate any extra memory proportional to the input size.
*/
