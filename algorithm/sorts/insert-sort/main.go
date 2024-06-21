package main

import "fmt"

// 带有哨兵
func sort(arr []int) []int {
	n := len(arr)
	arr2 := make([]int, n+1)
	// 哨兵位置
	arr2[0] = 0
	// 直接新切片中第一个元素
	arr2[1] = arr[0]
	// 逐一遍历arr数组
	for i := 1; i < n; i++ {
		// 第二个for循环进行将arr[i] 在arr2中找到正确位置
		for j := i; j >= 0; j-- {
			// 如果已经j到了哨兵位置，则说明arr[i]是在arr2数组是最小的
			if j == 0 {
				arr2[1] = arr[i]
			}
			// 如果arr[i] 较小
			if arr[i] < arr2[j] {
				// arr2元素向后移动
				arr2[j+1] = arr2[j]
			} else {
				// 如果arr[i]较大,则直接添加
				arr2[j+1] = arr[i]
				break
			}

		}
	}
	// 如果没有返回值，这样返回是错的，为什么呢？？
	// arr = make([]int, n)
	// arr = arr2[1:len(arr2)]

	// 创建一个新的空切片，长度和容量都为0
	// 注意：这个新切片的地址与原始切片不同
	//arr = arr[:0]

	// 正确返回方式
	return arr2[1:]
}

func main() {
	arr := []int{5, 3, 2, 10, 1, 9}
	fmt.Println(arr)
	arr = sort(arr)
	fmt.Println(arr)

}
