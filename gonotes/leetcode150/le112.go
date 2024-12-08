package leetcode150

func hasPathSum(root *TreeNode, targetSum int) bool {
	if root == nil {
		return false
	}
	var dfs func(root *TreeNode, curSum int) bool
	dfs = func(root *TreeNode, curSum int) bool {
		if root == nil {
			return false
		}
		curSum += root.Val
		if root.Left == nil && root.Right == nil {
			return curSum == targetSum
		}
		return dfs(root.Left, curSum) || dfs(root.Right, curSum)
	}
	return dfs(root, 0)
}
