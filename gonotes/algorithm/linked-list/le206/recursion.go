package le206

func ReverseList1(head *ListNode) *ListNode {
	var first *ListNode = nil
	return reverse(first, head)

}
func reverse(pre *ListNode, cur *ListNode) *ListNode {
	if cur == nil {
		return pre
	}
	temp := cur.Next
	// 顺序不能改变
	cur.Next = pre
	pre = cur
	cur = temp
	return reverse(pre, cur)
}
