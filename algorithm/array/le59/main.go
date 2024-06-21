package main

import (
	"fmt"
)

// 螺旋正方形，全都采取左闭右开的原则遍历
// le54是长方形，晚点再看
// 需要的遍历的次数为n/2,n为行数,如果n为3,n/2=1,只需要转一圈。
// 1	2	3
// 8	9	4
// 7	6	5

func generateMatrix(n int) [][]int {
	// 结果
	res := make([][]int, n)
	// 初始化
	for i := 0; i < n; i++ {
		res[i] = make([]int, n)
	}
	// 要加入的数字
	count := 1
	// 当前遍历的行
	i := 0
	// 当前遍历的列
	j := 0
	// 偏移量，由于是正方形，行和列的偏移量都是offset
	offset := 1
	// 记住全都采取左闭右开的原则遍历
	needCount := n / 2
	for needCount > 0 {
		// 遍历上边
		for ; j < n-offset; j++ {
			res[i][j] = count
			count++
		}
		// 结束后 i = 0	j = n-offset
		// 遍历右边
		for ; i < n-offset; i++ {
			res[i][j] = count
			count++
		}
		// 结束后 i = n-offset	j = n-offset
		// 遍历下边
		for ; j >= offset; j-- {
			res[i][j] = count
			count++
		}
		// 结束后 i = n-offset 	j = offset-1
		// 遍历左边
		for ; i >= offset; i-- {
			res[i][j] = count
			count++
		}
		// 圈圈结束后 i = offset-1	 j = offset-1
		i++
		j++
		offset++
		needCount--

	}
	// 假设n为3，则上述for循环后 offset = 2,i = 1 j = 1 还有(1,1)这个位置没有遍历
	// 假设n为5, 则上述for循环后 offset = 3,i = 2 j = 2 还有(2,2)这个位置没有遍历
	// fmt.Println(i)
	// fmt.Println(j)
	// fmt.Println(offset)
	if n%2 == 1 {
		res[i][j] = count
		count++
	}
	return res
}

func main() {
	fmt.Println(generateMatrix(3))
}
