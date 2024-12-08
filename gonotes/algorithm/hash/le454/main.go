package main

import "fmt"

// 四个数相加为0，注意题目：不用去重，求的是次数
// 暴力循环，n的四次方。
// 哈希思路：把nums1和nums2的值存入到第一个map1里面，同理。。第二个map2
// 由于要统计次数，map[v]count
// 这个方法为3*n²
func fourSumCount(nums1 []int, nums2 []int, nums3 []int, nums4 []int) int {
	map1 := make(map[int]int, 0)
	map2 := make(map[int]int, 0)
	count := 0
	for i := 0; i < len(nums1); i++ {
		for j := 0; j < len(nums2); j++ {
			map1[nums1[i]+nums2[j]]++
		}
	}
	for i := 0; i < len(nums3); i++ {
		for j := 0; j < len(nums4); j++ {
			map2[nums3[i]+nums4[j]]++
		}
	}
	// 注意次数
	// 比如map1里面-1出现了3次，map2里面1出现了6次
	//
	for v1, c1 := range map1 {
		if c2, ok := map2[0-v1]; ok {
			count += c1 * c2
		}
	}
	return count
}

// 优化，四个数字长度均为n
// 简洁一点，就是不要第三次的for循环,也不需要第二个map了变成2*n²
func fourSumCount1(nums1 []int, nums2 []int, nums3 []int, nums4 []int) int {
	n := len(nums1)
	map1 := make(map[int]int, n*n)
	count := 0
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			map1[nums1[i]+nums2[j]]++
		}
	}
	for _, v1 := range nums3 {
		for _, v2 := range nums4 {
			temp := 0 - v1 - v2
			// map查找时间复杂度为o(1),最坏情况下为o(n)
			if c, ok := map1[temp]; ok {
				count += c
			}
		}
	}
	return count
}
func main() {
	fmt.Println(fourSumCount([]int{0}, []int{0}, []int{0}, []int{0}))
	fmt.Println(fourSumCount1([]int{0}, []int{0}, []int{0}, []int{0}))
}
