package leetcode150

// 后序遍历，左右中
func postorderTraversal(root *TreeNode) []int {
	res := make([]int, 0)
	st := &myStack[*TreeNode]{}
	st.Push(root)
	for st.length != 0 {
		r := st.Pop()
		if r != nil {
			res = append(res, r.Val)
			st.Push(r.Left)
			st.Push(r.Right)
		} else {
			continue
		}
	}
	n := len(res)
	if n == 0 {
		return res
	}
	i := 0
	for j := n - 1; i < j; {
		res[i], res[j] = res[j], res[i]
	}
	return res
}
