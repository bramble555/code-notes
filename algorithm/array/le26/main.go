// 删除重复元素
// 思路1：和remove-element一样，快慢指针
package main

import (
	"fmt"
	"sort"
)

func removeDuplicates(nums []int) int {
	// 第一个元素(索引为0的元素)一定不重复
	slow := 1
	for fast := 1; fast < len(nums); fast++ {
		// 举例
		//	1	2 	2	2	3	3
		//		fast	fast	fast
		// 		slow	slow
		// 可以看出nums[fast]一直和nums[slow-1]比较
		if nums[fast] != nums[slow-1] {
			nums[slow] = nums[fast]
			slow++
		}
	}
	return slow
}

// 思路2：使用set，(map替代set),由于题目是非递减顺序，所以可以排序
// 效率没有第一种方法高，时间上进行了排序还有拷贝，空间上创建了set还有res
func removeDuplicates2(nums []int) int {
	set := make(map[int]struct{}, 0)
	for i := 0; i < len(nums); i++ {
		set[nums[i]] = struct{}{}
	}
	res := make([]int, 0)
	for v, _ := range set {
		res = append(res, v)
	}
	// 题目还要求改变nums
	sort.Ints(res)
	copy(nums, res)
	return len(res)
}
func main() {
	nums := []int{1, 1, 2}
	fmt.Println(removeDuplicates(nums))
	fmt.Println(removeDuplicates2(nums))
}
