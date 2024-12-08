package leetcode150

func getMask(word string) int {
	mask := 0
	for i := 0; i < len(word); i++ {
		mask |= 1 << (word[i] - 'a')
	}
	return mask
}
func maxProduct(words []string) int {
	n := len(words)
	// 每一个 word 生成 mask 和 wordsLen
	masks := make([]int, n)
	wordsLen := make([]int, n)
	for i := 0; i < n; i++ {
		masks[i] = getMask(words[i])
		wordsLen[i] = len(words[i])
	}
	// 比较每一对单词
	// & 运算 == 0 说明每一位都不同
	res := 0
	for i := 0; i < n-1; i++ {
		for j := i + 1; j < n; j++ {
			if masks[i]&masks[j] == 0 {
				res = max(wordsLen[i]*wordsLen[j], res)
			}
		}
	}
	return res
}
