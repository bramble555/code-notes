package leetcode150

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseKGroup(head *ListNode, k int) *ListNode {
	count := 0
	head1 := head
	headMap := make(map[int]*ListNode, 0)
	// 计算链表有多少个,记录对应索引和 Node
	for head != nil {
		headMap[count] = head
		count++
		head1 = head1.Next
	}
	// 计算反转次数
	times := count / k
	var time int = 1
	for time <= times {
		// 反转 k 个,分别是 从 (time-1)*k - (time)*k-1
		// 啥意思？ 假设 time 是 1，k 是 3，第一次是 从0-2
		// 第二次 time 是 2，从 3-5
		i := (time - 1) * k
		j := (time)*k - 1
		for i < j {
			headMap[i], headMap[j] = headMap[j], headMap[i]
			i++
			j--
		}
		time++
	}
	var res = &ListNode{}
	var cur = res
	for i := 0; i <= count; i++ {
		cur.Next = headMap[i]
		cur = cur.Next
	}
	
	return res.Next
}
