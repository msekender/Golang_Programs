/*
Solution: Regular Expression Matching
This problem can be efficiently solved using dynamic programming (DP). We define a 2D DP table where dp[i][j] represents whether the first i characters of s match the first j characters of p.

Approach
Define a DP table dp[i][j]:

dp[i][j] is true if s[0:i] matches p[0:j], otherwise false.
dp[0][0] = true since an empty string matches an empty pattern.
Handle the '*' wildcard:

If p[j-1] == '*', it can:
Ignore the preceding element (dp[i][j] = dp[i][j-2]).
Match multiple occurrences (dp[i][j] = dp[i-1][j] if p[j-2] == s[i-1] or p[j-2] == '.').
Handle direct character matches and '.' wildcard:

If p[j-1] == s[i-1] or p[j-1] == '.', then dp[i][j] = dp[i-1][j-1].
*/

package main

import (
	"fmt"
)

func isMatch(s string, p string) bool {
	m, n := len(s), len(p)

	// Create a DP table
	dp := make([][]bool, m+1)
	for i := range dp {
		dp[i] = make([]bool, n+1)
	}

	// Empty string matches empty pattern
	dp[0][0] = true

	// Pre-fill the table for patterns with '*' that can match empty string
	for j := 2; j <= n; j++ {
		if p[j-1] == '*' {
			dp[0][j] = dp[0][j-2] // '*' means removing the preceding character
		}
	}

	// Fill the DP table
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if p[j-1] == s[i-1] || p[j-1] == '.' {
				// If the characters match or pattern has '.', check the previous state
				dp[i][j] = dp[i-1][j-1]
			} else if p[j-1] == '*' {
				// '*' can either:
				// - Remove the preceding element
				// - Match multiple occurrences if the preceding character matches
				dp[i][j] = dp[i][j-2] // Case: '*' removes previous element
				if p[j-2] == s[i-1] || p[j-2] == '.' {
					dp[i][j] = dp[i][j] || dp[i-1][j] // Case: '*' repeats preceding char
				}
			}
		}
	}

	return dp[m][n]
}

func main() {
	// Example test cases
	fmt.Println(isMatch("aa", "a"))  // Output: false
	fmt.Println(isMatch("aa", "a*")) // Output: true
	fmt.Println(isMatch("ab", ".*")) // Output: true
}

/*
Time Complexity Analysis
Building the DP table requires O(m × n) operations, where:
m = len(s), n = len(p).
Each cell dp[i][j] is computed in O(1) time.
Given 1 ≤ s.length, p.length ≤ 20, this runs efficiently.

Space Complexity Analysis
O(m × n) space is required for the DP table.
This can be optimized to O(n) space by keeping only the previous row.
*/
