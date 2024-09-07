package main

import (
	"fmt"
	"os"
)

// 检测是否有目标目录，如果目录不在就创建
func main() {
	flag := isExist("./log")
	if flag {
		fmt.Println("log目录存在")
	} else {
		err := os.MkdirAll("./log", os.ModePerm)
		if err != nil {
			fmt.Println("创建目录失败", err)
			return
		}
	}
	filePath := "./log/september.log"
	if isExist(filePath) {
		fmt.Printf("%s is exist!", filePath)
	} else {
		_, err := os.Create(filePath)
		if err != nil {
			fmt.Println("这是什么玩意：", filePath, "\t创建出错了")
			return
		}
	}
	fmt.Println("目标文件成功")
}
func isExist(path string) bool {
	_, err := os.Stat(path)
	if err != nil {
		if os.IsExist(err) {
			return true
		}
		if os.IsNotExist(err) {
			return false
		}
		fmt.Println(err)
		return false
	}
	return true
}
