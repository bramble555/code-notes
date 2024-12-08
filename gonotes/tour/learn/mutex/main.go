package main

import (
	"fmt"
	"sort"
	"sync"
)

// map存放26个小写字母对应的个数
type safeCount struct {
	mu sync.Mutex
	v  map[string]int
}

func (s *safeCount) inc(key string) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.v[key]++
}

func (s *safeCount) getValue(key string) int {
	s.mu.Lock()
	defer s.mu.Unlock()
	return s.v[key]
}

func (s *safeCount) printMap() {
	// 如果想要顺序打印，可以先把i存储在切片里面
	letter := make([]string, 0)
	// 把s.v里面的每个字母添加到letter切片里面
	for l, _ := range s.v {
		letter = append(letter, l)
	}
	// 对letter进行排序
	sort.Strings(letter)
	for i := 0; i < 26; i++ {
		fmt.Print(letter[i], "=", s.v[letter[i]], "\t")
	}
}
func main() {
	b := 'a'
	countLetter := make(map[string]int, 26)
	// 初始化countLetter
	for i := b; i < b+26; i++ {
		countLetter[string(i)] = 0
	}
	sc := safeCount{v: countLetter}

	var wg sync.WaitGroup
	wg.Add(2) // 增加两个等待的goroutine

	// 同步增加1000次“a”
	go func() {
		defer wg.Done()
		for i := 0; i < 1000; i++ {
			sc.inc("a")
		}
	}()
	// 同步增加1000次“b”
	go func() {
		defer wg.Done()
		for i := 0; i < 1000; i++ {
			sc.inc("b")
		}
	}()
	wg.Wait()
	// time.Sleep(time.Second)
	sc.printMap()
}
