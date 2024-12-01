package leetcode150

/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Next *Node
 *     Random *Node
 * }
 */
type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

func copyRandomList(head *Node) *Node {
	head2 := head
	// 第一次遍历只添加 val，并且添加 oldNode 和 newNode 映射关系
	nodeMap := make(map[*Node]*Node, 0)
	dummy := &Node{}
	// res 永远指向最后一个元素
	res := dummy
	for head != nil {
		node := &Node{
			Val:    head.Val,
			Random: nil,
			Next:   nil,
		}
		// 添加结点
		res.Next = node
		// 建立映射关系
		nodeMap[head] = res.Next
		// 继续往后添加
		head = head.Next
		res = res.Next
	}
	// 第二次遍历
	head = head2
	for head != nil {
		newNode, _ := nodeMap[head]
		// 查看旧结点的 random
		newNode.Random = nodeMap[head.Random]
		head = head.Next
	}
	return dummy.Next

}
