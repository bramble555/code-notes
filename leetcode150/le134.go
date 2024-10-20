package leetcode150

import "fmt"

// gas = [2,3,4], cost = [3,4,3]
// 1-1
func canCompleteCircuit(gas []int, cost []int) int {
	n := len(gas)
	// 查看是否循环了
	count := 0
	for i := 0; i < n; i++ {
		// 每次重置 cur
		cur := 0
		for j := i; ; j++ {
			j = j % n

			if cur+gas[j] <= cost[j] {
				fmt.Println(i, "退出了")
				break
			}
			count++
			cur = cur + gas[j] - cost[j]
			fmt.Println("cur:", cur)
			fmt.Println("count:", count)
			// 满足条件
			if count == n {
				// fmt.Println(count)
				fmt.Println(i)
				fmt.Println(j)
				return i
			}
			// if cur == 0 {
			// 	break
			// }
		}
	}
	return -1

}
