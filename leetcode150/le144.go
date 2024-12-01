package leetcode150

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func preorderTraversal(root *TreeNode) []int {
	var st = &myStack[*TreeNode]{}
	res := make([]int, 0)
	st.Push(root)
	for st.length != 0 {
		r := st.Pop()
		res = append(res, r.Val)
		st.Push(r.Right)
		st.Push(r.Left)
	}
	return res
}
