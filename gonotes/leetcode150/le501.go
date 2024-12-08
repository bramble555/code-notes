package leetcode150

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func findMode(root *TreeNode) []int {
	var pre *TreeNode
	var dfs func(root *TreeNode)
	countMax := 0
	curCount := 0
	res := make([]int, 0)
	dfs = func(root *TreeNode) {
		if root == nil {
			return
		}
		dfs(root.Left)
		if pre != nil && pre.Val == root.Val {
			curCount++
		} else {
			curCount = 1
		}

		if curCount == countMax {
			res = append(res, root.Val)
		} else if curCount > countMax {
			res = []int{root.Val}
			countMax = curCount
		} else {

		}

		pre = root
		dfs(root.Right)
	}
	dfs(root)

	return res
}
