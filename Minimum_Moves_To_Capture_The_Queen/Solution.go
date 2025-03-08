/*
Solution Approach
We need to determine the minimum number of moves required for either the rook or the bishop to capture the black queen on an 8x8 chessboard. This problem can be solved using Breadth-First Search (BFS) since we are looking for the shortest path.

Steps to Solve
Model Chessboard as a Graph:
The Rook moves horizontally or vertically any number of squares.
The Bishop moves diagonally any number of squares.
Neither piece can jump over other pieces.
Use BFS for Shortest Path:
Start BFS from both the rook and bishop.
Move in all valid directions until capturing the queen.
Track the number of moves taken to reach (e, f).
Return the minimum moves among the two pieces.
*/

package main

import (
	"fmt"
)

// Position struct to represent a board position
type Position struct {
	x, y, moves int
}

// Directions for rook (horizontal & vertical)
var rookDirs = [][2]int{
	{0, 1}, {0, -1}, {1, 0}, {-1, 0},
}

// Directions for bishop (diagonal)
var bishopDirs = [][2]int{
	{1, 1}, {-1, -1}, {1, -1}, {-1, 1},
}

// Function to find the minimum moves to capture the queen
func minMovesToCaptureQueen(a, b, c, d, e, f int) int {
	boardSize := 8

	// BFS queue
	queue := []Position{
		{a, b, 0}, // Rook starting position
		{c, d, 0}, // Bishop starting position
	}

	// Visited set
	visited := make(map[[3]int]bool)

	for len(queue) > 0 {
		curr := queue[0]
		queue = queue[1:]

		// If we reach the queen, return the number of moves
		if curr.x == e && curr.y == f {
			return curr.moves
		}

		// Get valid moves based on piece type
		var moves [][2]int
		if curr.x == a && curr.y == b {
			moves = rookDirs // Rook movement
		} else {
			moves = bishopDirs // Bishop movement
		}

		// Try all valid moves
		for _, dir := range moves {
			newX, newY := curr.x, curr.y

			// Move in a straight line until hitting board limits or an obstacle
			for 1 <= newX+dir[0] && newX+dir[0] <= boardSize && 1 <= newY+dir[1] && newY+dir[1] <= boardSize {
				newX += dir[0]
				newY += dir[1]

				// Skip if reaching another piece's position (rook can't move over bishop, and vice versa)
				if (newX == c && newY == d) || (newX == a && newY == b) {
					break
				}

				// If not visited, add to the queue
				state := [3]int{newX, newY, curr.moves + 1}
				if !visited[state] {
					visited[state] = true
					queue = append(queue, Position{newX, newY, curr.moves + 1})
				}
			}
		}
	}
	return -1 // Should never reach here based on problem constraints
}

// Main function for testing
func main() {
	fmt.Println(minMovesToCaptureQueen(1, 1, 8, 8, 2, 3)) // Output: 2
	fmt.Println(minMovesToCaptureQueen(5, 3, 3, 4, 5, 2)) // Output: 1
}

/*
Time Complexity Analysis
Each piece can move in 4 or 8 directions (rook: 4, bishop: 4).
Each move scans up to 7 squares, so at worst, a BFS will traverse O(8 × 7) = O(56) ~ O(1) in a fixed 8x8 grid.
Total worst-case complexity is O(1) because the board size is constant.

Space Complexity Analysis
Visited states store at most O(64) = O(1) positions.
Queue storage has at most O(64) = O(1).
Thus, space complexity is O(1).
*/
