package leetcode150

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func partition(head *ListNode, x int) *ListNode {

	var dummy1 = &ListNode{}
	var dummy2 = &ListNode{}
	var lessList = dummy1
	var moreList = dummy2

	for head != nil {
		if head.Val < x {
			lessList.Next = head
			head = head.Next
			lessList = lessList.Next
		} else {
			moreList.Next = head
			head = head.Next
			moreList = moreList.Next
		}
	}
	lessList.Next = dummy2.Next
	moreList.Next = nil
	return dummy1.Next
}
