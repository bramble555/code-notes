package main

import "fmt"

// 单指针快排

// 分治法在每一层递归上都有三个步骤：
// 分解：将原问题分解为若干个规模较小，相互独立，与原问题形式相同的子问题；
// 解决：若子问题规模较小而容易被解决则直接解，否则递归地解各个子问题；
// 合并：将各个子问题的解合并为原问题的解
func sort(arr []int) {
	recursion(arr, 0, len(arr)-1)
}
func recursion(arr []int, left, right int) {
	// 如果left<right,则继续递归,直至left=right,也就是只有一个元素了
	// 分解操作
	if left < right {
		// 把left作为基准元素下标
		pivotIndex := partition(arr, left, right)
		partition(arr, left, pivotIndex-1)
		partition(arr, pivotIndex+1, right)
	}
}
func partition(arr []int, left, right int) int {
	// 解决操作
	pivotIndex := left
	swapIndex := left + 1
	for i := swapIndex; i <= right; i++ {
		// 如果基准元素大
		if arr[i] < arr[pivotIndex] {
			swap(arr, swapIndex, i)
			swapIndex++
		}
	}
	swap(arr, pivotIndex, swapIndex-1)
	return swapIndex - 1
}
func swap(arr []int, i, j int) {
	arr[i], arr[j] = arr[j], arr[i]
}
func main() {
	arr := []int{5, 3, 2, 10, 1, 9}
	fmt.Println(arr)
	sort(arr)
	fmt.Println(arr)
}
