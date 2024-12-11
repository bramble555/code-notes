package leetcode150

func deleteNode(root *TreeNode, key int) *TreeNode {
	if root == nil {
		return nil
	}
	if root.Val == key {
		if root.Left == nil && root.Right == nil {
			return nil
		}
		if root.Left != nil && root.Right == nil {
			return root.Left
		}
		if root.Right != nil && root.Left == nil {
			return root.Right
		}
		if root.Right != nil && root.Left != nil {
			cur := root.Left
			for cur.Right != nil {
				cur = cur.Right
			}
			cur.Right = root.Right
			return root.Left
		}
	}
	if root.Val > key {
		root.Left = deleteNode(root.Left, key)
	}
	if root.Val < key {
		root.Right = deleteNode(root.Right, key)
	}
	return root
}
