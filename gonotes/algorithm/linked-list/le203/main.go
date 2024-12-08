package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
type ListNode struct {
	Val  int
	Next *ListNode
}

func removeElements(head *ListNode, val int) *ListNode {
	// 设置一个哨兵结点
	dummy := &ListNode{-1, head}
	// 记录前一个结点
	pre := dummy
	// 记录当前结点
	cur := dummy.Next
	// 当cur == nil 的时候，说明前一个结点有效，当前cur无效
	for cur != nil {
		if cur.Val == val {
			pre.Next = cur.Next
			cur = cur.Next
		} else {
			pre = cur
			cur = cur.Next
		}
	}
	return dummy.Next

}
func main() {

}
