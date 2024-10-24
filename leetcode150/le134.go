package leetcode150

import "fmt"

// gas = [2,3,4],
// cost= [3,4,3]
// 	   -1 1 -1

// gas = [1,2,3,4,5],
// cost= [3,4,5,1,2]
// 	   -2-2-2 3 3

// gas = [4, 5, 2, 6, 5, 3]
// cost= [3, 2, 7, 3, 2, 9]
// 	    1  3 -5  3  3  -6

// gas = [5,1,2,4,4]
// cost= [4,4,1,5,1]
//  	    1-3 1-1 3
func canCompleteCircuit(gas []int, cost []int) int {
	start := 0
	n := len(gas)
	curSum := 0
	totalSum := 0
	for i := 0; i < n; i++ {
		totalSum += gas[i] - cost[i]
		curSum += gas[i] - cost[i]
		fmt.Println(curSum)
		if curSum < 0 {
			curSum = 0
			start = i + 1
		}
	}
	if totalSum < 0 {
		return -1
	}
	return start

}
