package main

// 去重非常难想到，多思考，多看，
import (
	"fmt"
	"sort"
)

// 注意题目：要求返回的数字，不能下标 答案中不可以包含重复的三元组
// 思路：排序，然后固定一个数字i，然后双指针，j从固定的数字的后面开始，k从len-1开始，找到相加为0的，加入结果
// 去重：我的思路：固定i，j=i+1，如果j >= i+1 并且和 nums[j] = nums[j-1] continue
// 那么k去重？？k <= len(nums)-2 并且 nums[k] = nums[k+1] continue
// 有错误，容易误会：i != j、i != k 且 j != k
// 当nums为{0,0,0,0}的时候，输出的是{{0,0,0}}
//
//	       i,j,k
//	          i j k
//	这时候第二次i=第一次的j了，不能算结果，没有去重
//
// 不对！！ 以上这几行是错误的，是有了{0，0，0}就不能 再有{0,0,0,}了 和i，j，k无关
func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	res := make([][]int, 0)
	// 注意i范围，i，j，k各三个位置
	// k最大为len(nums)-1,j最大为len(nums)-2
	for i := 0; i <= len(nums)-3; i++ {
		// 每次重置
		j := i + 1
		k := len(nums) - 1
		// i的去重 和前一个比较，不能与后一个比较，因为后一个是j
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		for j < k {
			sum := nums[i] + nums[j] + nums[k]
			if sum > 0 {
				k--
			} else if sum < 0 {
				j++
			} else {
				// 把加入结果的现在nums[j]记录下来，方便后面j++，k同理
				l := nums[j]
				r := nums[k]
				res = append(res, []int{nums[i], l, r})
				// 正确做法应该在加入结果的时候去重
				for j < k && nums[j] == l {
					j++
				}
				for j < k && nums[k] == r {
					k--
				}

			}
		}

	}
	return res
}
func main() {
	// 排序后为-4	-1 -1 0 1 2
	// 答案为	i   j     k
	// 		   i j k
	nums := []int{0, 0, 0, 0}
	nums2 := []int{-4, -1, -1, 0, 0, 1, 2}
	fmt.Println(threeSum(nums))
	fmt.Println(threeSum(nums2))
}
