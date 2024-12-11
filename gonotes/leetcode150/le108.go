package leetcode150

func sortedArrayToBST(nums []int) *TreeNode {
	var build func(nums []int, start, end int) *TreeNode
	build = func(nums []int, start, end int) *TreeNode {
		index := start + end/2
		var root = &TreeNode{Val: nums[index]}
		root.Left = build(nums, start, index-1)
		root.Right = build(nums, index+1, end)
		return root
	}
	return build(nums, 0, len(nums)-1)
}
