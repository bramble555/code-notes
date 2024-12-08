package leetcode150

import "container/list"

func findBottomLeftValue(root *TreeNode) int {
	var maxDep int = -1
	var res int
	var dfs func(root *TreeNode, dep int)
	dfs = func(root *TreeNode, dep int) {
		list.New()
		if root == nil {
			return
		}
		// 根节点
		if root.Left == nil && root.Right == nil {
			if maxDep < dep {
				maxDep = dep
				res = root.Val
			}
		}
		if root.Left != nil {
			dep++
			dfs(root.Left, dep)
			dep--
		}
		if root.Right != nil {
			dep++
			dfs(root.Right, dep)
			dep--
		}
	}
	dfs(root, 0)
	return res
}
