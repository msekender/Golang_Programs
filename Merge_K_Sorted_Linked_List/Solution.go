/*
Solution Approach
The most efficient way to merge k sorted linked lists is using a min-heap (priority queue) because:

We always extract the smallest element from the lists efficiently.
We push the next element of the extracted list into the heap.
Steps:

Use a min-heap (priority queue) to store the smallest elements from each list.
Pop the smallest element from the heap and append it to the merged list.
If the popped node has a next node, push it into the heap.
Repeat until all nodes are processed.
*/

package main

import (
	"container/heap"
	"fmt"
)

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

// Priority queue (min-heap) implementation
type MinHeap []*ListNode

func (h MinHeap) Len() int            { return len(h) }
func (h MinHeap) Less(i, j int) bool  { return h[i].Val < h[j].Val } // Min-heap based on node value
func (h MinHeap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *MinHeap) Push(x interface{}) { *h = append(*h, x.(*ListNode)) }
func (h *MinHeap) Pop() interface{} {
	n := len(*h)
	x := (*h)[n-1]
	*h = (*h)[:n-1]
	return x
}

// Merges k sorted linked lists using a priority queue (min-heap)
func mergeKLists(lists []*ListNode) *ListNode {
	minHeap := &MinHeap{}
	heap.Init(minHeap)

	// Push the head of each linked list into the heap
	for _, list := range lists {
		if list != nil {
			heap.Push(minHeap, list)
		}
	}

	// Dummy head for the merged linked list
	dummy := &ListNode{}
	curr := dummy

	// Extract min element and push next node from the same list
	for minHeap.Len() > 0 {
		smallest := heap.Pop(minHeap).(*ListNode)
		curr.Next = smallest
		curr = curr.Next

		if smallest.Next != nil {
			heap.Push(minHeap, smallest.Next)
		}
	}

	return dummy.Next
}

// Helper function to create a linked list from a slice
func createLinkedList(arr []int) *ListNode {
	if len(arr) == 0 {
		return nil
	}
	head := &ListNode{Val: arr[0]}
	curr := head
	for _, val := range arr[1:] {
		curr.Next = &ListNode{Val: val}
		curr = curr.Next
	}
	return head
}

// Helper function to print a linked list
func printLinkedList(head *ListNode) {
	curr := head
	for curr != nil {
		fmt.Print(curr.Val, " -> ")
		curr = curr.Next
	}
	fmt.Println("nil")
}

func main() {
	// Example test case
	lists := []*ListNode{
		createLinkedList([]int{1, 4, 5}),
		createLinkedList([]int{1, 3, 4}),
		createLinkedList([]int{2, 6}),
	}

	mergedList := mergeKLists(lists)
	printLinkedList(mergedList)
}

/*
Time Complexity Analysis
Insertion and extraction operations from a min-heap take O(log k).
Since we process all nodes (N total across all lists), each node is pushed and popped from the heap once.
Therefore, the total complexity is O(N log k), where:
N is the total number of nodes across all lists.
k is the number of linked lists.

Space Complexity Analysis
The min-heap stores at most k elements at any time, leading to O(k) space.
The output linked list is formed in O(N) space, but this is the required output, so it doesn’t count as extra space.
Thus, the additional space complexity is O(k).
*/
