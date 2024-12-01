package leetcode150

import (
	"fmt"
)

const maxSize = int(1e4)

type mystack struct {
	data   []any
	length int // 当前长度
	top    int // 栈顶指针
	base   int // 栈底指针
}

func Init() *mystack {
	var s mystack
	s.base = 0
	s.top = 0
	s.length = 0
	s.data = make([]any, maxSize)
	return &s
}
func (s *mystack) push(str rune) {
	// 判断是否满了
	if s.top == maxSize {
		fmt.Println("栈满了")
		return
	}
	s.data[s.top] = str
	s.top++
	s.length++
}
func (s *mystack) pop() rune {
	// 判断是否空了
	if s.top == 0 || s.length == 0 {
		fmt.Println("栈空了")
		return 0
	}
	s.top--
	s.length--
	return s.data[s.top].(rune)

}

func isValid(s string) bool {
	st := Init()
	sStr := []rune(s)
	n := len(sStr)
	for i := 0; i < n; i++ {
		if sStr[i] == '[' || sStr[i] == '{' || sStr[i] == '(' {
			st.push(sStr[i])
		}

		if sStr[i] == ']' || sStr[i] == '}' || sStr[i] == ')' {
			popV := st.pop()
			if (sStr[i] == ']' && popV != '[') ||
				(sStr[i] == '}' && popV != '{') ||
				(sStr[i] == ')' && popV != '(') {
				return false
			}
		}
	}
	return st.length == 0
}
