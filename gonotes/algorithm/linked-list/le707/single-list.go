package le707

// 不循环带有虚拟头节点的单链表
type singleList struct {
	Val  int
	Next *singleList
}
type MyLinkedList struct {
	Dummy *singleList
	Size  int
}

func Constructor() MyLinkedList {
	dummy := &singleList{-1, nil}
	return MyLinkedList{dummy, 0}
}

func (this *MyLinkedList) Get(index int) int {
	// 判断索引是否合法
	if index < 0 || index >= this.Size {
		return -1
	}
	cur := this.Dummy.Next
	for i := 0; i < index; i++ {
		cur = cur.Next
	}
	return cur.Val
}

func (this *MyLinkedList) AddAtHead(val int) {
	dummy := this.Dummy
	node := &singleList{val, nil}
	node.Next = dummy.Next
	dummy.Next = node
	this.Size++
}

func (this *MyLinkedList) AddAtTail(val int) {
	cur := this.Dummy
	for i := 0; i < this.Size; i++ {
		cur = cur.Next
	}
	// 找到了当前要添加的位置 的 前一个位置
	node := &singleList{val, nil}
	cur.Next = node
	this.Size++

}

func (this *MyLinkedList) AddAtIndex(index int, val int) {
	// 题目说 如果 index 比长度更大，该节点将 不会插入 到链表中。
	// 易错，不需要等于号，如果等于，那么相当于在tail位置加入元素
	if index > this.Size {
		return
	}
	cur := this.Dummy
	for i := 0; i < index; i++ {
		cur = cur.Next
	}
	// 找到了当前要添加的位置 的 前一个位置
	node := &singleList{val, cur.Next}
	cur.Next = node
	this.Size++

}

func (this *MyLinkedList) DeleteAtIndex(index int) {
	// 判断索引是否合法
	if index < 0 || index >= this.Size {
		return
	}
	cur := this.Dummy
	for i := 0; i < index; i++ {
		cur = cur.Next
	}
	// 找到了当前要删除的位置 的 前一个位置
	cur.Next = cur.Next.Next
	this.Size--
}

/**
 * Your MyLinkedList object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Get(index);
 * obj.AddAtHead(val);
 * obj.AddAtTail(val);
 * obj.AddAtIndex(index,val);
 * obj.DeleteAtIndex(index);
 */
