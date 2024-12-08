package leetcode150

import "math"

// 递归
func isBalanced(root *TreeNode) bool {
	he := getHeigh(root)
	if he == -1 {
		return false
	}
	return true
}
func getHeigh(root *TreeNode) int {
	if root == nil {
		return 0
	}
	leftHei := getHeigh(root.Left)
	rigHei := getHeigh(root.Right)
	if leftHei == -1 || rigHei == -1 {
		return -1
	}
	res := math.Abs(float64(leftHei) - float64(rigHei))
	if res > 1 {
		return -1
	}
	return max(leftHei, rigHei) + 1
}
