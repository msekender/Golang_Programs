/*
Corrected Approach
Sliding Window (Two-Pointer Technique):
Use two pointers (left and right) to maintain a valid window that contains:
All vowels at least once.
Exactly k consonants.
Use a vowel frequency map to track vowels.
Use consonantCount to track consonants.
Count Valid Substrings Properly:
Once a valid window is found (consonantCount == k and all vowels exist):
Count substrings starting from left.
Increment left cautiously, ensuring we still have a valid substring.
*/

package main

import (
	"fmt"
)

// Function to count valid substrings
func countSubstrings(word string, k int) int {
	vowels := map[byte]bool{'a': true, 'e': true, 'i': true, 'o': true, 'u': true}
	vowelCount := map[byte]int{'a': 0, 'e': 0, 'i': 0, 'o': 0, 'u': 0}
	n := len(word)
	left, consonantCount, totalCount, validStart := 0, 0, 0, -1

	for right := 0; right < n; right++ {
		char := word[right]

		// Update vowel count if it’s a vowel
		if vowels[char] {
			vowelCount[char]++
		} else {
			consonantCount++
		}

		// Shrink the window if consonants exceed k
		for consonantCount > k {
			if vowels[word[left]] {
				vowelCount[word[left]]--
			} else {
				consonantCount--
			}
			left++
		}

		// Ensure all vowels are in the window
		if isValid(vowelCount) && consonantCount == k {
			// Find the first position where all vowels appear
			if validStart == -1 {
				validStart = left
			}
			totalCount += (left - validStart + 1)
		}
	}

	return totalCount
}

// Helper function to check if all vowels are present at least once
func isValid(vowelCount map[byte]int) bool {
	return vowelCount['a'] > 0 && vowelCount['e'] > 0 && vowelCount['i'] > 0 && vowelCount['o'] > 0 && vowelCount['u'] > 0
}

// Testing the function
func main() {
	fmt.Println(countSubstrings("aeioqq", 1))         // Output: 0
	fmt.Println(countSubstrings("aeiou", 0))          // Output: 1
	fmt.Println(countSubstrings("ieaouqqieaouqq", 1)) // Output: 3
}

/*
Time Complexity Analysis
Expanding right pointer: O(n).
Shrinking left pointer: O(n).
Checking vowel validity: O(1).
Total Complexity: O(n) + O(n) = O(n) (Linear time).

Space Complexity Analysis
Vowel map: O(1).
Fixed integer variables: O(1).
Total Complexity: O(1) (Constant extra space).
*/
