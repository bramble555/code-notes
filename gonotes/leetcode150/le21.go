package leetcode150

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	// 创建一个哨兵节点，方便处理边界情况
	dummy := &ListNode{}
	res := dummy
	for list1 != nil && list2 != nil {
		if list1.Val <= list2.Val {
			res.Next = list1
			list1 = list1.Next
		} else {
			res.Next = list2
			list2 = list2.Next
		}
		res = res.Next
	}
	for list1 != nil {
		res.Next = list1
		list1 = list1.Next
		res = res.Next
	}
	for list2 != nil {
		res.Next = list2
		list2 = list2.Next
		res = res.Next
	}
	return res
}
