package leetcode150

func countNodes(root *TreeNode) int {
	count := 0
	var traversal func(root *TreeNode)
	traversal = func(root *TreeNode) {
		if root == nil {
			return
		}
		count++
		traversal(root.Left)
		traversal(root.Right)
	}
	traversal(root)
	return count
}
