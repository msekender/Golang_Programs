/*
Solution Approach:
We use backtracking to generate all possible letter combinations.
We define a mapping of digits (2-9) to their respective letter groups.
We recursively explore all possible letter combinations for the given digits.
*/

package main

import (
	"fmt"
)

// Mapping of digits to letters
var phoneMap = map[byte]string{
	'2': "abc", '3': "def", '4': "ghi", '5': "jkl",
	'6': "mno", '7': "pqrs", '8': "tuv", '9': "wxyz",
}

// Backtracking function
func backtrack(index int, digits string, current string, result *[]string) {
	// Base case: if we have processed all digits, add the combination to result
	if index == len(digits) {
		*result = append(*result, current)
		return
	}

	// Get the letters corresponding to the current digit
	letters := phoneMap[digits[index]]

	// Iterate through each letter and recurse
	for i := 0; i < len(letters); i++ {
		backtrack(index+1, digits, current+string(letters[i]), result)
	}
}

// Main function to generate letter combinations
func letterCombinations(digits string) []string {
	if len(digits) == 0 {
		return []string{}
	}

	var result []string
	backtrack(0, digits, "", &result)
	return result
}

// Testing the function
func main() {
	digits := "23"
	output := letterCombinations(digits)
	fmt.Println(output) // Expected: ["ad","ae","af","bd","be","bf","cd","ce","cf"]
}

/*
Complexity Analysis
Time Complexity:
Each digit in the input string can represent 3 or 4 letters.
If there are n digits, the total number of combinations is approximately:
3ⁿ to 4ⁿ possible combinations.
Since we generate all combinations recursively, the worst-case time complexity is:
O(4ⁿ) (since the digit '7' and '9' map to 4 letters, and others map to 3).

Space Complexity:
The space required for storing the result is O(4ⁿ) in the worst case.
The recursive stack depth is O(n) due to backtracking.
Total Space Complexity: O(4ⁿ + n) ≈ O(4ⁿ).
*/
