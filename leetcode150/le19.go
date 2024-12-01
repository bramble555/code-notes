package leetcode150

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	count := 1
	headMap := make(map[int]*ListNode, 0)
	for head != nil {
		headMap[count] = head
		count++
		head = head.Next
	}
	// 假如只有一个元素 ,结束后 count = 2
	// 要删除的结点的 索引为 count  - n
	var dummy = &ListNode{}
	res := dummy
	for i := 1; i < count; i++ {
		if i == count-n {
			continue
		}
		res.Next = headMap[i]
		res = res.Next
	}
	return dummy.Next
}
