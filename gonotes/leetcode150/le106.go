package leetcode150

func buildTree(inorder []int, postorder []int) *TreeNode {
	// 存放 中序v遍历的 值 和 索引 ,方便后面分割
	inMap := make(map[int]int, 0)
	for i := 0; i < len(inorder); i++ {
		inMap[inorder[i]] = i
	}
	// 这个 rootIndex 是 在后序遍历中的 Index
	// l 和 r 是中序遍历中的 边界
	var builder func(inorder []int, postorder []int, rootIndex int, l, r int) *TreeNode
	builder = func(inorder []int, postorder []int, rootIndex int, l, r int) *TreeNode {
		if l > r {
			return nil
		}
		if l == r {
			return &TreeNode{Val: inorder[l]}
		}
		// 查找 RootIndex 在 后序遍历的 值
		rootVal := postorder[rootIndex]
		root := &TreeNode{Val: rootVal}
		// 查找 rootVal 在 中序遍历中的 Index
		divideIndex := inMap[rootVal]
		root.Left = builder(inorder, postorder, rootIndex-1-(r-divideIndex), l, divideIndex-1)
		root.Right = builder(inorder, postorder, rootIndex-1, divideIndex+1, r)
		return root
	}
	return builder(inorder, postorder, len(postorder)-1, 0, len(inorder)-1)
}
