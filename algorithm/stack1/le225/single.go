package le225

// 一个队列模拟
type MyStack struct {
	queue []int // 创建一个队列
}

func Constructor() MyStack {
	return MyStack{queue: make([]int, 0)}
}

func (this *MyStack) Push(x int) {
	this.queue = append(this.queue, x)
}

// 思路是 queue 1234
//        123重新加入队列，4123
// val = 4 queue=queue[1:]
func (this *MyStack) Pop() int {
	n := len(this.queue) - 1
	for n != 0 {
		val := this.queue[0]
		this.queue = this.queue[1:]
		this.queue = append(this.queue, val)
		n--
	}
	val := this.queue[0]
	this.queue = this.queue[1:]
	return val
}

func (this *MyStack) Top() int {
	val := this.Pop()
	this.queue = append(this.queue, val)
	return val
}

func (this *MyStack) Empty() bool {
	n := len(this.queue)
	if n != 0 {
		return false
	}
	return true
}

/**
 * Your MyStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.Empty();
 */
