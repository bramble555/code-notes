package leetcode150

// 只 push 当前元素
// 先一路向左，当 当前元素的左孩子为 nil 的时候，那就 pop 元素，然后再跳到 cur 的右孩子
func inorderTraversal(root *TreeNode) []int {
	res := make([]int, 0)
	st := &myStack[*TreeNode]{}
	cur := root
	for cur != nil || st.length != 0 {
		if cur != nil {
			st.Push(cur)
			cur = cur.Left
		} else {
			cur = st.Pop()
			res = append(res, cur.Val)
			cur = cur.Right
		}
	}
	return res
}
