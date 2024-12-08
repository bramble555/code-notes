package leetcode150

type Node22 struct {
	key   int
	value int
	Pre   *Node22
	Next  *Node22
}
type LRUCache struct {
	Dummy    *Node22
	Tail     *Node22
	curLen   int
	cap      int
	orderMap map[int]*Node22
}

func Constructor22(capacity int) LRUCache {
	dummy := Node22{}
	tail := Node22{}
	// 将 Dummy 和 Tail 节点连接起来
	dummy.Next = &tail
	tail.Pre = &dummy
	return LRUCache{
		Dummy:    &dummy,
		Tail:     &tail,
		curLen:   0,
		cap:      capacity,
		orderMap: make(map[int]*Node22),
	}
}
func (this *LRUCache) deleteNode(node *Node22) {
	node.Pre.Next = node.Next
	node.Next.Pre = node.Pre
	delete(this.orderMap, node.key)
	this.curLen--
}

// 添加到末尾
func (this *LRUCache) addNode(node *Node22) {
	node.Pre = this.Tail.Pre
	this.Tail.Pre.Next = node
	this.Tail.Pre = node
	node.Next = this.Tail
	this.orderMap[node.key] = node
	this.curLen++
}
func (this *LRUCache) Get(key int) int {
	if node, ok := this.orderMap[key]; ok {
		this.deleteNode(node)
		this.addNode(node)
		return node.value
	}
	return -1
}

func (this *LRUCache) Put(key int, value int) {
	if node, ok := this.orderMap[key]; ok {
		this.deleteNode(node)
		newNode := &Node22{
			key:   key,
			value: value,
		}
		this.addNode(newNode)
		return
	}
	if this.curLen < this.cap {
		newNode := &Node22{
			key:   key,
			value: value,
		}
		this.addNode(newNode)
		return
	}
	newNode := &Node22{
		key:   key,
		value: value,
	}
	this.deleteNode(this.Dummy.Next)
	this.addNode(newNode)

}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
