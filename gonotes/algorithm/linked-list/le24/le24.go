package le24

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

// to do 不知道哪里有问题
// func SwapPairs(head *ListNode) *ListNode {
// 	// 	1 	2 	3 	4
// 	// 	cur	back
// 	//			cur	back
// 	//            				cur
// 	//	1	2	3
// 	// 	cur	back
// 	//			cur 	back

// 	// 采用虚拟头节点，当然也可以不使用
// 	// 首先知道循环终止条件 cur.next == nil
// 	// 不采用虚拟头节点
// 	cur := head
// 	// 如果传来空指针
// 	if cur == nil {
// 		return cur
// 	}
// 	// 如果有元素
// 	back := cur.Next
// 	res := back
// 	// 只有一个元素
// 	if back == nil {
// 		// 不用交换，直接返回第一个元素
// 		return cur
// 	}
// 	for cur != nil && cur.Next != nil {
// 		if cur == nil || cur.Next == nil {
// 			break
// 		}
// 		temp := back.Next
// 		// cur.Next = temp
// 		// back.Next = cur
// 		// 交换
// 		back.Next = cur
// 		cur.Next = temp
// 		// 更新
// 		cur = temp
// 		if cur != nil {
// 			back = cur.Next
// 		}

// 	}

// 	return res
// }
// 有虚拟头节点
func SwapPairs(head *ListNode) *ListNode {
	dummy := &ListNode{-1, head}
	//dummy 1	2	3	4
	//cur 	back	temp
	//			cur	back

	//dummy	1	2	3
	//
	//			cur	back
	cur := dummy
	back := cur.Next
	temp := back.Next
	for back != nil && temp != nil {
		// 交换
		cur.Next = temp
		back.Next = temp.Next
		temp.Next = back
		// 更新
		cur = back
		back = cur.Next

	}
	return dummy.Next

}
