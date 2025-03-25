/*
Solution Approach:
We’ll improve the algorithm by using coordinate compression and attempting all possible pairs of two cuts (either vertical or horizontal):

Steps:
Collect all unique x (or y) start and end points.

Sort these points and simulate all possible pairs of cuts.

For each pair:

Split the grid into 3 sections.

Assign each rectangle to one section (based on its position).

Check:

Each section has at least one rectangle.

No rectangle spans across two sections.
*/

package main

import (
	"fmt"
	"sort"
)

func canCutGrid(n int, rectangles [][]int) bool {
	return checkCuts(rectangles, true) || checkCuts(rectangles, false)
}

func checkCuts(rects [][]int, vertical bool) bool {
	// Step 1: Get all unique start and end points along the axis (x or y)
	var points []int
	for _, r := range rects {
		if vertical {
			points = append(points, r[0], r[2])
		} else {
			points = append(points, r[1], r[3])
		}
	}
	sort.Ints(points)

	// Step 2: Try all unique pairs of cuts
	for i := 1; i < len(points); i++ {
		for j := i + 1; j < len(points); j++ {
			cut1, cut2 := points[i], points[j]
			sections := [3]int{} // Count rectangles in 3 sections

			valid := true
			for _, r := range rects {
				var lo, hi int
				if vertical {
					lo, hi = r[0], r[2]
				} else {
					lo, hi = r[1], r[3]
				}

				if hi <= cut1 {
					sections[0]++
				} else if lo >= cut2 {
					sections[2]++
				} else if lo >= cut1 && hi <= cut2 {
					sections[1]++
				} else {
					// Rectangle spans across sections → invalid
					valid = false
					break
				}
			}

			if valid && sections[0] > 0 && sections[1] > 0 && sections[2] > 0 {
				return true
			}
		}
	}

	return false
}

// ---- Test Code ----
func main() {
	tests := []struct {
		n          int
		rectangles [][]int
		expected   bool
	}{
		{5, [][]int{{1, 0, 5, 2}, {0, 2, 2, 4}, {3, 2, 5, 3}, {0, 4, 4, 5}}, true},
		{4, [][]int{{0, 0, 1, 1}, {2, 0, 3, 4}, {0, 2, 2, 3}, {3, 0, 4, 3}}, true},
		{4, [][]int{{0, 2, 2, 4}, {1, 0, 3, 2}, {2, 2, 3, 4}, {3, 0, 4, 2}, {3, 2, 4, 4}}, false}, // ✅ Now handled correctly
	}

	for _, test := range tests {
		result := canCutGrid(test.n, test.rectangles)
		fmt.Printf("Input: n = %d, rectangles = %v => Output: %v (Expected: %v)\n",
			test.n, test.rectangles, result, test.expected)
	}
}

/*
Time Complexity
Let:

N = number of rectangles

P = number of unique endpoints (start/end x or y values). At most 2N.

Step-by-step Breakdown:

Collect all endpoints:

We go through N rectangles, and collect 2 values per rectangle.

Time: O(N)

Sort the endpoints:

Sorting 2N values.

Time: O(N log N)

Try all unique pairs of cut points:

In worst case, there are P ≈ 2N unique points.

Number of pairs = O(P^2) = O(N^2)

Simulate assignment of rectangles for each pair:

For each pair of cuts, we loop through all N rectangles.

Cost per pair = O(N)

Total cost = O(P^2 * N) = O(N^3) worst-case (if all start/end points are unique)

BUT… in practice:
Many coordinates will be repeated.

Early termination (as soon as a valid cut is found) prevents checking all P^2 pairs.

Real-world performance is much closer to O(N^2).

Space Complexity
Breakdown:
Storing endpoints:

At most 2N integers

Space: O(N)

Sorted coordinate list:

Space: O(N)

Constant auxiliary space:

For counters and variables: O(1)

Total: O(N) space */
