package main

import (
	"fmt"
)

// le704
func search(nums []int, target int) int {
	return search3(nums, 0, len(nums)-1, target)
}

func main() {
	arr := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println(search(arr, 0))

}

// 标准版
// 前俩个版本均会多查询一次而且当只有一个元素的时候，会发生错误
// 注意循环不变量，不要把已经查找过的再加入区间
func search3(arr []int, left, right, target int) int {
	for left <= right { // 注意这里通常使用<=而不是<
		mid := left + (right-left)/2 // 避免溢出
		// target在mid左边
		if arr[mid] > target {
			right = mid - 1
			// target 在mid右边
		} else if arr[mid] < target {
			left = mid + 1
			// 找到了target
		} else {
			return mid
		}
	}
	// 未找到
	return -1
}
