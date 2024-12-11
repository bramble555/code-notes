package leetcode150

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
var pre *TreeNode

func dfsd(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	root.Right = dfsd(root.Right)
	if pre != nil {
		root.Val += pre.Val
	}
	pre = root
	root.Left = dfsd(root.Left)
	return root
}
func convertBST(root *TreeNode) *TreeNode {
	pre = nil
	return dfsd(root)
}
