package main

import (
	"fmt"
)

// Function to find the length of the longest substring without repeating characters
func lengthOfLongestSubstring(s string) int {
	charIndex := make(map[byte]int)
	maxLength, left := 0, 0

	for right := 0; right < len(s); right++ {
		if index, found := charIndex[s[right]]; found && index >= left {
			left = index + 1
		}
		charIndex[s[right]] = right
		maxLength = max(maxLength, right-left+1)
	}

	return maxLength
}

// Helper function to get max of two integers
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	s1 := "abcabcbb"
	fmt.Println("Input:", s1)
	fmt.Println("Output:", lengthOfLongestSubstring(s1))

	s2 := "bbbbb"
	fmt.Println("Input:", s2)
	fmt.Println("Output:", lengthOfLongestSubstring(s2))

	s3 := "pwwkew"
	fmt.Println("Input:", s3)
	fmt.Println("Output:", lengthOfLongestSubstring(s3))
}

/*
Time Complexity Analysis
We iterate through the string once with right pointer (O(n)).
The left pointer also moves forward at most n times (O(n)).
Each character is added/updated in the hashmap in constant time (O(1)).
Overall complexity: O(n).

Space Complexity Analysis
The hashmap stores at most unique characters (limited to 128 ASCII characters).
Worst case: If all characters are unique, the hashmap holds O(min(n, 128)) entries.
Since 128 is a constant, the space complexity is O(n) in the worst case.
*/
