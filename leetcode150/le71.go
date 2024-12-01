package leetcode150

import "strings"

type T any
type myStack[T any] struct {
	arr    []T
	length int
}

func Constructord[T any]() *myStack[T] {
	return &myStack[T]{
		arr: make([]T, 0),
	}
}
func (m *myStack[T]) Push(val T) {
	m.arr = append(m.arr, val)
	m.length++
}
func (m *myStack[T]) Pop() T {
	if m.length > 0 {
		temp := m.arr[len(m.arr)-1]
		m.arr = m.arr[:len(m.arr)-1]
		m.length--
		return temp
	}
	var zeroValue T // 返回类型的零值
	return zeroValue
}
func simplifyPath(path string) string {
	_pathArr := strings.Split(path, "/")
	pathArr := make([]string, 0)
	for i := 0; i < len(_pathArr); i++ {
		if _pathArr[i] != "" {
			pathArr = append(pathArr, _pathArr[i])
		}
	}

	d := Constructord[string]()
	for i := 0; i < len(pathArr); i++ {
		if pathArr[i] == "." {
			continue
		}
		if pathArr[i] == ".." {
			d.Pop()
			continue
		}
		d.Push(pathArr[i])
	}
	if d.length == 0 {
		return "/"
	}
	res := ""
	n := d.length
	arr := make([]string, 0, n)
	for d.length > 0 {
		arr = append(arr, d.Pop())
	}
	for i := 0; i < n; i++ {
		arr[i], arr[n-1-i] = arr[n-1-i], arr[i]
	}
	for i := 0; i < n; i++ {
		res = "/" + arr[i]
	}
	// for d.length > 0 {
	// 	temp := d.Pop()
	// 	res = "/" + temp + res
	// }
	return res
}
