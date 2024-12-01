package le206

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

// 思路：一个pre，一个cur，
// pre cur
// 	1-> 2-> 3-> 4
//	1<- 2<- 3<- 4
// 一直到            pre cur
func ReverseList(head *ListNode) *ListNode {
	// 不能这样写，需要显式声明
	// pre := nil
	var pre *ListNode = nil
	cur := head
	for cur != nil {
		// 需要定义一个临时变量来存储cur.Next,否则cur不能变成cur的下一个节点
		temp := cur.Next
		// 注意这三个顺序不能变
		cur.Next = pre
		pre = cur
		cur = temp
	}
	return pre
}
