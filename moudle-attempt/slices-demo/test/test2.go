package test

import "fmt"

func foo(a []int) {
	// 传参时拷贝了新的切片，因此当新切片的长度发生改变时，原切片并不会发生改变。而且在函数 foo 中，新切片 a 增加了 8 个元素，原切片对应的底层数组不够放置这 8 个元素，因此申请了新的空间来放置扩充后的底层数组。
	// 这个时候新切片和原切片指向的底层数组就不是同一个了。
	fmt.Println(len(a), cap(a))
	a = append(a, 1, 2, 3)

	a[0] = 200
}

func Slices2() {
	// slices陷阱，
	a := []int{1, 2}
	foo(a)
	fmt.Println(a) // 这样会输出1,2
}

// 修改  去返回切片这样就可以了
func foo2(a []int) []int {
	// 传参时拷贝了新的切片，因此当新切片的长度发生改变时，原切片并不会发生改变。而且在函数 foo 中，新切片 a 增加了 8 个元素，原切片对应的底层数组不够放置这 8 个元素，因此申请了新的空间来放置扩充后的底层数组。
	// 这个时候新切片和原切片指向的底层数组就不是同一个了。
	fmt.Println(len(a), cap(a))
	a = append(a, 1, 2, 3, 4, 5, 6, 7, 8)

	a[0] = 200
	return a
}
func Slices22() {
	a := []int{1, 2}
	a = foo2(a)
	fmt.Println(a)
}
