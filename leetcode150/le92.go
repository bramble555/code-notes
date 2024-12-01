package leetcode150

func reverseBetween(head *ListNode, left int, right int) *ListNode {
	count := 0
	//     存放 下标和 List
	headMap := make(map[int]*ListNode, 0)
	for head != nil {
		headMap[count] = head
		head = head.Next
		count++
	}
	left = left - 1
	right = right - 1
	for left < right {
		// 获取左右两端的节点
		leftNode := headMap[left]
		rightNode := headMap[right]

		// 确保不会访问 nil 节点
		if leftNode == nil || rightNode == nil {
			panic("Invalid index in headMap")
		}

		// 交换节点
		headMap[left], headMap[right] = rightNode, leftNode

		// 更新索引
		left++
		right--
	}
	var dummy = &ListNode{}
	var cur = dummy
	for i := 0; i <= count; i++ {
		cur.Next = headMap[i]
		cur = cur.Next
	}
	return dummy.Next
}
