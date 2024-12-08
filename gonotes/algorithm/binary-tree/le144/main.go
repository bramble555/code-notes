package main

import (
	"fmt"
	"gonotes/algorithm/stack1/le225"
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

// 先序遍历，递归实现
func preorderTraversal(root *TreeNode) []int {
	res := make([]int, 0)
	traversal(root, &res)
	return res
}

// 因为traversal函数里面要增加cap(res),而且不能返回res，所以要传入 *res
func traversal(node *TreeNode, res *[]int) {
	if node == nil {
		return
	}
	// 先序遍历是中左右
	*res = append(*res, node.Val)
	traversal(node.Left, res)
	traversal(node.Right, res)
}

// 以上是俩个函数，可以变为一个函数,这样就避免了传入res
func preorderTraversal11(root *TreeNode) []int {
	res := make([]int, 0)
	var traversal11 func(root *TreeNode)
	traversal11 = func(root *TreeNode) {
		if root == nil {
			return
		}
		// 先序遍历是中左右
		res = append(res, root.Val)
		traversal11(root.Left)
		traversal11(root.Right)
	}
	// 调用
	traversal11(root)
	return res
}

// 非递归实现  中左右
func PreorderTraversal111(root *TreeNode) []any {
	mystack := le225.Constructor()
	mystack.Push(root)
	res := make([]any, 0)
	for len(mystack.Queue) != 0 {
		// 需要断言获取值
		node := mystack.Pop().(*TreeNode)
		res = append(res, node.Val)
		// 先加入右子树，递归的时候先遍历出左子树
		if node.Right != nil {
			mystack.Push(node.Right)
		}
		if node.Left != nil {
			mystack.Push(node.Left)
		}

	}
	return res
}
func main() {
	tree1 := &TreeNode{1, nil, nil}
	tree21 := &TreeNode{2, nil, nil}
	tree22 := &TreeNode{3, nil, nil}
	tree31 := &TreeNode{4, nil, nil}
	tree1.Left = tree21
	tree1.Right = tree22
	tree21.Left = tree31
	fmt.Println(PreorderTraversal111(tree1))
}
