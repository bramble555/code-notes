package main

import (
	"fmt"
	"runtime"
)

func main() {
	// 获取CPU数量
	cpuNum := runtime.NumCPU()
	fmt.Printf("Number of CPU(s): %d\n", cpuNum)

	// 设置可并行运行的goroutine的最大数目等于CPU数量
	// runtime.GOMAXPROCS(cpuNum + 10)
	fmt.Printf("GOMAXPROCS set to: %d\n", runtime.GOMAXPROCS(0)) // 传递0表示查询当前设置
}
