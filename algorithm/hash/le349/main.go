package main

import (
	"fmt"
)

// 思路：把nums1添加到map1里面，然后遍历num2，查找map1，如果出现过，加入到另外一个map,最后把map转换为数组输出

// 提示：
// 1 <= nums1.length, nums2.length <= 1000
// 0 <= nums1[i], nums2[i] <= 1000
// 用数组解决
// 头有点发昏，明天再写一遍
func intersection12(nums1 []int, nums2 []int) []int {
	res1 := make([]int, 10001)
	res2 := make([]int, 10001)
	// 思路，不能像上面加减判断了，如果+1，-1判断，那么和原始0(没有出现过的数组无法区分)
	// 那么nums1出现的值置为1,nums2出现过的值也置为1，最后相加=2的那么就是答案了
	for i := 0; i < len(nums1); i++ {
		res1[nums1[i]] = 1
	}
	for i := 0; i < len(nums2); i++ {
		res2[nums2[i]] = 1
	}
	// for _, v := range nums2 {
	// 	res1[v] = 1
	// }
	res := make([]int, 0)
	for i := 0; i < 1001; i++ {
		if res1[i]+res2[i] == 2 {
			res = append(res, i)
		}
	}
	return res

}

// 由于没有set，用map解决
func intersection11(nums1 []int, nums2 []int) []int {
	resMap := make(map[int]int, 0)
	map1 := make(map[int]int, 0)
	// 把nums1添加到map1里面
	for _, v := range nums1 {
		map1[v]++
	}
	// 然后遍历num2，查找上一个map
	for _, v := range nums2 {
		for i, _ := range map1 {
			if v == i {
				resMap[i]++
			}
		}
	}
	res := make([]int, 0)
	for i, _ := range resMap {
		res = append(res, i)
	}
	return res

}

// 对于上面的优化：
// 1.浪费了map后面的val
// 2.查找map1的时候使用了for循环，应该让底层自己查找，效率更高
func intersection(nums1 []int, nums2 []int) []int {
	// 值是空的struct类型（作为占位符）
	resMap := make(map[int]struct{}, 0)
	map1 := make(map[int]struct{}, 0)
	// 把nums1添加到map1里面
	for _, v := range nums1 {
		map1[v] = struct{}{}
	}
	// // 然后遍历num2，查找上一个map
	for _, v := range nums2 {
		if _, ok := map1[v]; ok {
			resMap[v] = struct{}{}
		}
	}
	res := make([]int, 0)
	for i, _ := range resMap {
		res = append(res, i)
	}
	return res
}

func main() {
	nums1 := []int{1, 2}
	nums2 := []int{1}
	fmt.Println(intersection11(nums1, nums2))
	fmt.Println(intersection(nums1, nums2))
	fmt.Println(intersection12(nums1, nums2))
}
