package le225

// 一个队列模拟
type MyStack struct {
	Queue []any // 创建一个队列
}

func Constructor() MyStack {
	return MyStack{Queue: make([]any, 0)}
}

func (this *MyStack) Push(x any) {
	this.Queue = append(this.Queue, x)
}

// 思路是 queue 1234
//        123重新加入队列，4123
// val = 4 queue=queue[1:]
func (this *MyStack) Pop() any {
	n := len(this.Queue) - 1
	for n != 0 {
		val := this.Queue[0]
		this.Queue = this.Queue[1:]
		this.Queue = append(this.Queue, val)
		n--
	}
	val := this.Queue[0]
	this.Queue = this.Queue[1:]
	return val
}

func (this *MyStack) Top() any {
	val := this.Pop()
	this.Queue = append(this.Queue, val)
	return val
}

func (this *MyStack) Empty() bool {
	n := len(this.Queue)
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
