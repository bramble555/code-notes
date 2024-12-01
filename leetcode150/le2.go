package leetcode150

/**
 * Definition for singly-linked list.

 */

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	// 头节点
	var res *ListNode
	// 当前结点，用于添加元素
	var cur *ListNode
	var carry int
	for l1 != nil || l2 != nil || carry > 0 {
		var num1 int
		var num2 int
		if l1 != nil {
			num1 = l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			num2 = l2.Val
			l2 = l2.Next
		}

		sum := num1 + num2 + carry
		carry /= sum
		var node = ListNode{
			Val: sum % 10,
		}
		if res == nil {
			res = &node
			cur = res
		} else {
			cur.Next = &node
			cur = cur.Next
		}

	}
	return res
}
