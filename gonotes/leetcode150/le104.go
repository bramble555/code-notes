package leetcode150

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func maxDepth(root *TreeNode) int {
	maxDep := 0
	if root == nil {
		return maxDep
	}
	cur := root
	que := Queue[*TreeNode]{}
	que.Enqueue(cur)
	for que.IsEmpty() {
		maxDep++
		length := que.Length()
		for length > 0 {
			r := que.Dequeue()
			length--
			if r.Left != nil {
				que.Enqueue(r.Left)
			}
			if r.Right != nil {
				que.Enqueue(r.Right)
			}
		}
	}
	return maxDep
}
