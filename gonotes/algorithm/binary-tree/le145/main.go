package main

import (
	"fmt"
	"gonotes/algorithm/basic"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 后序遍历
func postorderTraversal(root *TreeNode) []int {
	res := make([]int, 0)
	traversal(root, &res)
	return res
}
func traversal(node *TreeNode, res *[]int) {
	if node == nil {
		return
	}
	traversal(node.Left, res)
	traversal(node.Right, res)
	*res = append(*res, node.Val)
}

// 左右中
func postorderTraversal111(root *TreeNode) []int {
	stack := basic.ConstructorMystack()
	var node *TreeNode
	res := make([]int, 0)

	for root != nil || !stack.IsEmpty() {
		// 首先尽可能地向左走，并将沿途的节点压入栈中
		for root != nil {
			stack.Push(root)
			root = root.Left
		}
		// 当无法继续向左走时，从栈中弹出一个节点
		node = stack.Pop().(*TreeNode)
		// 将弹出的节点的值添加到结果列表中
		res = append(res, node.Val)
		// 如果右子节点存在，先将其压入栈中，以便后续访问
		if node.Right != nil {
			stack.Push(node.Right)
		}
		// 注意：这里不需要设置 node = node.Right，因为下一次循环会从栈中弹出右子节点（如果存在的话）
	}

	return res
}
func main() {
	tree1 := &TreeNode{1, nil, nil}
	tree21 := &TreeNode{2, nil, nil}
	tree22 := &TreeNode{3, nil, nil}
	tree31 := &TreeNode{4, nil, nil}
	tree32 := &TreeNode{5, nil, nil}
	tree33 := &TreeNode{6, nil, nil}
	tree1.Left = tree21
	tree1.Right = tree22
	tree21.Left = tree31
	tree21.Right = tree32
	tree22.Left = tree33
	fmt.Println(postorderTraversal111(tree1))
}
