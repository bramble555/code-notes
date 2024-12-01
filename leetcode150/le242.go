package leetcode150

func isAnagram(s string, t string) bool {
	n := len(s)
	if n != len(t) {
		return false
	}
	mapS := make(map[uint8]int, n)
	for i := 0; i < n; i++ {
		mapS[s[i]]++
	}
	for i := 0; i < n; i++ {
		_, ok := mapS[t[i]]
		if !ok {
			return false
		}
		mapS[t[i]]--
		if mapS[t[i]] < 0 {
			return false
		}
	}
	return true
}
