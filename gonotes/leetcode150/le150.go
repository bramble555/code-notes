package leetcode150

import "strconv"

func evalRPN(tokens []string) int {
	st := Constructord[string]()
	res := 0
	for _, token := range tokens {

		if token == "+" || token == "-" || token == "*" || token == "/" {
			temp := 0
			b, _ := strconv.Atoi(st.Pop())
			a, _ := strconv.Atoi(st.Pop())
			if token == "+" {
				temp = a + b
			}
			if token == "-" {
				temp = a - b
			}
			if token == "*" {
				temp = a * b
			}
			if token == "/" {
				temp = a / b
			}
			st.Push(strconv.Itoa(temp))
		} else {
			st.Push(token)
		}
	}

	res, _ = strconv.Atoi(st.Pop())
	return res
}
