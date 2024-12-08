package leetcode150

// "badc"

// "baba"
func isIsomorphic(s string, t string) bool {
	//
	n := len(s)
	mapStr := make(map[uint8]uint8)
	// 把 t[i] 加入 arr
	arr := make(map[uint8]struct{}, 0)
	for i := 0; i < n; i++ {
		v, ok := mapStr[s[i]]

		// s[i] 不存在
		if !ok {
			// 但是 t[i] 存在
			if _, ok := arr[t[i]]; ok {
				return false
			}
			mapStr[s[i]] = t[i]
			arr[t[i]] = struct{}{}
		} else {
			if v != t[i] {
				return false
			}
		}
	}
	return true
}
