package leetcode150

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func rotateRight(head *ListNode, k int) *ListNode {
	cur := head
	if head == nil || head.Next == nil {
		return head
	}
	count := 1
	for cur != nil && cur.Next != nil {
		cur = cur.Next
		count++
	}
	k %= count
	if k == 0 {
		return head
	}
	// cur 是最后一个结点 把最后一个结点，指向 head
	cur.Next = head
	index := count - k - 1
	newLastNode := head
	temp := 0
	for temp < index {
		newLastNode = newLastNode.Next
		temp++
	}
	newHeadNode := newLastNode.Next
	newLastNode.Next = nil
	return newHeadNode
}
