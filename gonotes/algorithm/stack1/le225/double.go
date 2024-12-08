package le225

// // 假设arr1是 123
// // Push arr1 1234
// // Pop arr1 4  arr2 123  then arr1 = arr2  len(arr2)=0
// // Top arr1 4 arr2 123 then arr1 = arr2  arr1=append(arr1,arr2[0])
// type MyStack struct {
// 	arr1 []int // 队列，先进先出
// 	arr2 []int // 队列，先进先出
// }

// func Constructor() MyStack {
// 	return MyStack{
// 		arr1: make([]int, 0),
// 		arr2: make([]int, 0),
// 	}
// }

// func (this *MyStack) Push(x int) {
// 	this.arr1 = append(this.arr1, x)
// }
// func (this *MyStack) Pop() int {
// 	this.to()
// 	// swap arr1 and arr2
// 	this.arr1, this.arr2 = this.arr2, this.arr1
// 	res := this.arr2[0]
// 	// clear arr2
// 	this.arr2 = make([]int, 0)
// 	return res
// }

// func (this *MyStack) Top() int {
// 	this.to()
// 	// swap arr1 and arr2
// 	this.arr1, this.arr2 = this.arr2, this.arr1
// 	this.arr1 = append(this.arr1, this.arr2[0])
// 	res := this.arr2[0]
// 	return res
// }

// func (this *MyStack) Empty() bool {
// 	n := len(this.arr1)
// 	if n != 0 {
// 		return false
// 	}
// 	return true

// }

// // 这个函数负责把arr1里面的元素只剩一个，其他都移到arr2里面
// func (this *MyStack) to() {
// 	n := len(this.arr1)
// 	this.arr2 = make([]int, n-1)
// 	for j := 0; j <= n-2; j++ {
// 		this.arr2[j] = this.arr1[0]
// 		this.arr1 = this.arr1[1:]
// 	}
// }

// /**
//  * Your MyStack object will be instantiated and called as such:
//  * obj := Constructor();
//  * obj.Push(x);
//  * param_2 := obj.Pop();
//  * param_3 := obj.Top();
//  * param_4 := obj.Empty();
//  */
