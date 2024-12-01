package leetcode150

func romanToInt(s string) int {
	ri := make(map[rune]int, 0)
	ri['I'] = 1
	ri['V'] = 5
	ri['X'] = 10
	ri['L'] = 50
	ri['C'] = 100
	ri['D'] = 500
	ri['M'] = 1000
	sRune := []rune(s)
	n := len(sRune)
	var res int
	var i int = 0
	for i < n-1 {
		val := ri[sRune[i]]
		valLat := ri[sRune[i+1]]
		if val >= valLat {
			res += val
			i++
			continue
		}
		if sRune[i] == 'I' && (sRune[i+1] == 'V' || sRune[i+1] == 'X') {
			res += valLat - val
			i += 2
			continue
		}
		if sRune[i] == 'X' && (sRune[i+1] == 'L' || sRune[i+1] == 'C') {
			res += valLat - val
			i += 2
			continue
		}
		if sRune[i] == 'C' && (sRune[i+1] == 'D' || sRune[i+1] == 'M') {
			res += valLat - val
			i += 2
			continue
		}
	}
	if i == n-1 {
		res += ri[sRune[n-1]]
	}
	return res
}
