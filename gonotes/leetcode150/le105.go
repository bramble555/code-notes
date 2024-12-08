package leetcode150

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func buildTree2(preorder []int, inorder []int) *TreeNode {
	inMap := make(map[int]int, 0)
	for i := 0; i < len(inorder); i++ {
		inMap[inorder[i]] = i
	}

	var helper func(preorder []int, inorder []int, rootIndex, l, r int) *TreeNode
	helper = func(preorder []int, inorder []int, rootIndex, l, r int) *TreeNode {
		if l > r {
			return nil
		}
		if l == r {
			return &TreeNode{Val: inorder[l]}
		}
		// 查找 rootIndex 在 先序遍历中的 Val
		rootVal := preorder[rootIndex]
		root := &TreeNode{Val: rootVal}
		// 查找 rootVal 在 中序遍历中的 位置,然后拆分
		divideIndex := inMap[rootVal]
		root.Left = helper(preorder, inorder, rootIndex+1, l, divideIndex-1)
		root.Right = helper(preorder, inorder, rootIndex+1+(divideIndex-l), divideIndex+1, r)
		return root
	}
	return helper(preorder, inorder, 0, 0, len(preorder)-1)
}
