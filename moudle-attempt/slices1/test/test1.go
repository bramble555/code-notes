package test

import "fmt"

// 测试切片部分索引传递到函数中长度会发生什么变化？？
// &arr 是原始切片 arr 的地址，而 &arr（在 testTransmitLen 函数中）是传递给该函数的切片参数的地址。这两个切片值是不同的，尽管它们可能引用相同的底层数组。
// 要打印切片的底层数组的地址，你可以使用 &arr[0] 而不是 &arr，因为 &arr[0] 会给你切片中第一个元素的地址，这通常是底层数组的地址（除非切片是空的）。但是，注意这只是一个惯例，并且依赖于切片不是nil且非空。
func Slices1() {
	arr := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	testTransmitLen(arr[1:5])
	fmt.Println("全部切片为", arr)
	fmt.Printf("全部切片第0个地址为%p", &arr[0])
	fmt.Printf("全部切片第一个地址为%p", &arr[1]) // 和传入的切片地址一样，引用的数组是一样的

}
func testTransmitLen(arr []int) {
	fmt.Println("切片长度为", len(arr))
	fmt.Println("部分切片为", arr)
	fmt.Printf("部分切片地址为%p", &arr[0])
}

// 结论：引用的是同一个数组,长度是传入多少长度，len(arr)就为多大
