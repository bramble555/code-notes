package splicestrmode

import (
	"math/rand"
	"strings"
)

const letterBtyes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func randomString(n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = letterBtyes[rand.Intn(len(letterBtyes))]
	}
	return string(b)
}

// 把那个randomString字符串用+字符加上n次
func PlusConcat(n int, str string) string {
	s := ""
	for i := 0; i < n; i++ {
		s += str
	}
	return s
}

// Builder
func BuilderConcat(n int, str string) string {
	var b strings.Builder
	for i := 0; i < n; i++ {
		b.WriteString(str)
	}
	return b.String()
}
