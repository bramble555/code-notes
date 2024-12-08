package main

import "fmt"

func main() {
	fmt.Println(sortArray([]int{1, 3, 0, 5, 2}))
}

// 快排 acwing 这个并没有保证每次找到的 pivotal 元素，左边都比 pivotal 小，右边都比 pivotal 大
func sortArray(nums []int) []int {
	quick(nums, 0, len(nums)-1)
	return nums
}
func quick(nums []int, l, r int) {
	if l >= r {
		return
	}
	var i = l - 1
	var j = r + 1
	var pivotal = nums[l]
	for i < j {
		i++
		for nums[i] < pivotal {
			i++
		}
		j--
		for nums[j] > pivotal {
			j--
		}
		if i < j {
			temp := nums[j]
			nums[j] = nums[i]
			nums[i] = temp
		}
	}
	quick(nums, l, j)
	quick(nums, j+1, r)
}
