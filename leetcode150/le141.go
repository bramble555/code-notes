package leetcode150

/**
 * Definition for singly-linked list.

 */
type ListNode struct {
	Val  int
	Next *ListNode
}

func hasCycle(head *ListNode) bool {
	temp := head
	n := int(1e4 + 1)
	for i := 0; i < n; i++ {
		temp = temp.Next
		if temp == nil {
			return false
		}
	}
	return true
}
