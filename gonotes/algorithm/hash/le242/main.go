package main

import (
	"fmt"
)

// 不用hashmap,用数组
// 思路是记录s中每个字符出现的个数，然后在t里面查看是否一样
func isAnagram1(s string, t string) bool {
	n := len(s)
	m := len(t)
	if n != m {
		return false
	}
	arr := make([]int, 26)
	// 记录s中每个字符出现的个数
	for i := 0; i < n; i++ {
		arr[s[i]-'a']++
	}
	// 减去t中每个字符出现的个数
	for i := 0; i < m; i++ {
		arr[t[i]-'a']--
	}
	// 查看arr数组中是否全为0
	for i := 0; i < 26; i++ {
		if arr[i] != 0 {
			return false
		}
	}
	return true
}

// 如果要遍历个数小的时候，使用切片
// 使用hashmap,如果要遍历的个数非常大的时候，每次遍历非常耗时，此时使用map
func isAnagram(s string, t string) bool {
	n := len(s)
	m := len(t)
	if n != m {
		return false
	}
	// 可以使用map[int]int
	// 也可以使用map[string]int
	hashmap := make(map[string]int, 26)
	// 记录s中每个字符出现的个数
	for i := 0; i < n; i++ {
		letter := string(s[i])
		hashmap[letter]++
	}
	// 减去t中每个字符出现的个数
	for i := 0; i < m; i++ {
		letter := string(t[i])
		hashmap[letter]--
	}
	for _, v := range hashmap {
		if v != 0 {
			return false
		}
	}
	return true
}
func main() {
	s := "ada"
	t := "daa"
	fmt.Println(string(s[0]))
	fmt.Println(isAnagram1(s, t))
	fmt.Println(isAnagram(s, t))
}
