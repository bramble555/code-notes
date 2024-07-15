package main

// 双指针

func reverseString(s []byte) {
	// i不等于，可以举例子
	for i, j := 0, len(s)-1; i < len(s)/2; {
		swap(s, i, j)
		i++
		j--
	}
	// 或者
	/*
		for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
				swap(sStr, i, j)
			}
	*/
}
func swap(s []byte, i, j int) {
	s[i], s[j] = s[j], s[i]
}
