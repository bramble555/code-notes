package main

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

// 先序遍历
func preorderTraversal(root *TreeNode) []int {
	res := make([]int, 0)
	traversal(root, &res)
	return res
}

// 因为traversal函数里面要增加cap(res),而且不能返回res，所以要传入 *res
func traversal(node *TreeNode, res *[]int) {
	if node == nil {
		return
	}
	// 先序遍历是中左右
	*res = append(*res, node.Val)
	traversal(node.Left, res)
	traversal(node.Right, res)
}

// 以上是俩个函数，可以变为一个函数,这样就避免了传入res
func preorderTraversal11(root *TreeNode) []int {
	res := make([]int, 0)
	var traversal11 func(root *TreeNode)
	traversal11 = func(root *TreeNode) {
		if root == nil {
			return
		}
		// 先序遍历是中左右
		res = append(res, root.Val)
		traversal11(root.Left)
		traversal11(root.Right)
	}
	// 调用
	traversal11(root)
	return res
}
