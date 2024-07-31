package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 后序遍历
func postorderTraversal(root *TreeNode) []int {
	res := make([]int, 0)
	traversal(root, &res)
	return res
}
func traversal(node *TreeNode, res *[]int) {
	if node == nil {
		return
	}
	traversal(node.Left, res)
	traversal(node.Right, res)
	*res = append(*res, node.Val)
}
