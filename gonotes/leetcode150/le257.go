package leetcode150

import "strconv"

func binaryTreePaths(root *TreeNode) []string {
	res := []string{}
	dfs(root, "", &res)
	return res
}
func dfs(root *TreeNode, path string, res *[]string) {
	if root == nil {
		return
	}
	curPath := path + strconv.Itoa(root.Val)
	// 根节点
	if root.Left == nil && root.Right == nil {
		*res = append(*res, curPath)
	} else {
		curPath += "->"
		if root.Left != nil {
			dfs(root.Left, curPath, res)
		}
		if root.Right != nil {
			dfs(root.Right, curPath, res)
		}
	}
}
