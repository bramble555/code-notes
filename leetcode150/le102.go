package leetcode150

type Queue[T any] struct {
	items []T
}

// 向队列尾部添加元素

func (q *Queue[T]) Enqueue(item T) {
	q.items = append(q.items, item)
}

// 从队列头部移除并返回元素
func (q *Queue[T]) Dequeue() T {
	item := q.items[0]
	q.items = q.items[1:]
	return item
}

// 检查队列是否为空
func (q *Queue[T]) IsEmpty() bool {
	return len(q.items) == 0
}

// 返回队列的长度
func (q *Queue[T]) Length() int {
	return len(q.items)
}

// 查看队列头部元素但不移除
func (q *Queue[T]) Peek() (T, bool) {
	if q.IsEmpty() {
		var zeroValue T // 获取零值
		return zeroValue, false
	}
	return q.items[0], true
}

func levelOrder(root *TreeNode) [][]int {
	res := make([][]int, 0)
	var que = Queue[*TreeNode]{}
	que.Enqueue(root)
	for !que.IsEmpty() {
		size := que.Length()
		tempRes := make([]int, 0)
		for size > 0 {
			r := que.Dequeue()
			tempRes = append(tempRes, r.Val)
			if r.Left != nil {
				que.Enqueue(r.Left)
			}
			if r.Right != nil {
				que.Enqueue(r.Right)
			}
			size--
		}
		res = append(res, tempRes)
	}
	return res
}
