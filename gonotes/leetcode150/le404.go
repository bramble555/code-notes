package leetcode150

func sumOfLeftLeaves(root *TreeNode) int {
	if root == nil || (root.Left == nil && root.Right == nil) {
		return 0
	}
	leftVal := 0
	rigVal := 0
	if root.Left != nil && root.Left.Left == nil && root.Left.Right == nil {
		leftVal = root.Left.Val
	} else {
		leftVal = sumOfLeftLeaves(root.Left)
	}
	rigVal = sumOfLeftLeaves(root.Right)
	return leftVal + rigVal
}
