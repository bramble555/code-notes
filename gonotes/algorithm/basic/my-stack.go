/*
我的栈
1. 构造方法
2. 入栈(push)
3. 出栈(pop)
4.获取栈顶元素(peek)
5.获取元素个数(getSize)
6.判断栈是否为空(isEmpty)
7. search
8. clear
9. 获取所有值
*/
package basic

// 切片实现栈
type MyStack struct {
	arr []any
}

func ConstructorMystack() *MyStack {
	return &MyStack{arr: make([]any, 0)}
}
func (mystack *MyStack) Push(x any) {
	mystack.arr = append(mystack.arr, x)
}

// 删除栈顶元素，也就是最后一个元素，并且返回栈顶元素
func (mystack *MyStack) Pop() any {
	res := mystack.arr[mystack.GetSize()-1]
	mystack.arr = mystack.arr[:mystack.GetSize()-1]
	return res
}
func (mystack *MyStack) Peek() any {
	return mystack.arr[mystack.GetSize()-1]
}
func (mystack *MyStack) GetSize() int {
	return len(mystack.arr)
}

func (mystack *MyStack) IsEmpty() bool {
	return mystack.GetSize() == 0
}

// search 0  -   n-1
func (mystack *MyStack) Search(x int) any {
	n := mystack.GetSize() - 1
	if x > n-1 || x < 0 {
		return nil
	}
	return mystack.arr[n-x-1]
}
func (mystack *MyStack) Clear() {
	mystack.arr = make([]any, 0)
}

// 获所有取值
func (mystack *MyStack) GetValues() []any {
	return append([]any{}, mystack.arr...)
}
