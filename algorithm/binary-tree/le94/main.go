package main

import (
	"fmt"
	"gonotes/algorithm/basic"
)

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func inorderTraversal(root *TreeNode) []int {
	res := make([]int, 0)
	var traversal func(node *TreeNode)
	traversal = func(node *TreeNode) {
		if node == nil {
			return
		}
		traversal(node.Left)
		res = append(res, node.Val)
		traversal(node.Right)
	}
	// 调用
	traversal(root)
	return res
}

// 左中右
func inorderTraversal111(root *TreeNode) []int {
	res := make([]int, 0)
	stack := basic.ConstructorMystack()
	node := root
	for node != nil || stack.GetSize() != 0 {
		if node != nil {
			// 一路向左
			stack.Push(node)
			node = node.Left
		} else {
			node = stack.Pop().(*TreeNode)
			res = append(res, node.Val)
			// 这里关键：因为node已经判断了node.left是否为空，所以这里暗含了判断node.right是否为空，然后为空，那么到下一轮就删除上一轮node的父结点(判断2结点的right是否为空，然后删除了1结点)，也就是当前node的父父结点
			node = node.Right
		}
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
	fmt.Println(inorderTraversal111(tree1))
}
