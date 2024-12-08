package leetcode150

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func deleteDuplicates(head *ListNode) *ListNode {
	var dummy = &ListNode{}
	res := dummy

	for head != nil {
		// 找到第一个循环结点，一直跳到最后一个循环结点
		// 保证 head.Next ！= nil  可以进行比较
		if head.Next != nil && head.Val == head.Next.Val {

			for head.Next != nil && head.Val == head.Next.Val {
				head = head.Next
			}
			// 跳过最后一个循环结点
			head = head.Next
		} else {
			res.Next = head
			res = res.Next
			head = head.Next
		}
	}
	res.Next = nil
	return dummy.Next
}
