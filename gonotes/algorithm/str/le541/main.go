package main

import (
	"fmt"
)

// 题目：不能看懂，看了评论：每次去2K个字符串，前k个翻转，后k个不反转
// 如果后面不够2k个，有俩种情况：如果剩余字符少于 k 个，则将剩余字符全部反转。如果剩余字符小于 2k 但大于或等于 k 个，则反转前 k 个字符，其余字符保持原样。
func reverseStr(s string, k int) string {
	// 字符串不可变，所以要变成切片
	sStr := []byte(s)
	n := len(sStr)
	// 如果k<len(s),先计算要循环的次数
	count := n / (2 * k)
	if n%(2*k) != 0 {
		count++
	}
	i := 0
	// 注意j，比如k为2，则只需要翻转0和1
	j := i + k - 1
	// 注意，如果第一次的k直接超出了范围，那么直接翻转全部
	if n <= j {
		j = n - 1
		for ; i < j; i, j = i+1, j-1 {
			swap(sStr, i, j)
		}
	}
	// 如果第一次的k超出了范围，count为0
	for count > 0 {
		// 特殊情况，后面不够2k,此时j超出范围
		if j >= n {
			// 分析，无论后面<k还是k<后面<2k，都只需要把n-1与i之间的元素进行翻转
			j = n - 1
			// exceed := j - n
			// 如果刚好的话exceed=0,如果exceed=k说明超出了
			for ; i < j; i, j = i+1, j-1 {
				swap(sStr, i, j)
			}
		}
		// 在for循环里面i的值会变，所以要先保存起来
		temp := i
		// 每次翻转
		for ; i < j; i, j = i+1, j-1 {
			swap(sStr, i, j)
		}
		// 翻转第一次是0,第二次是2k
		i = temp + 2*k
		j = i + k - 1
		count--
	}
	return string(sStr)

}
func swap(s []byte, i, j int) {
	s[i], s[j] = s[j], s[i]
}
func main() {
	// n=5,k=2,反转了2次，就是5/4=1 +1？？
	// n=8,k=2,反转了2次，就是8/4=2
	// n=9,k=2,反转了3次，就是9/4=2	+1
	s := "abcdefg"
	fmt.Println(s)
	fmt.Println(reverseStr(s, 8))

}
