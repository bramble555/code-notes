package flagtest

import (
	sys_flag "flag"
	"fmt"
)

// 命令行 用 gotest 测试不了
func Flagdemo1() {
	// 用户权限，默认是 用户权限
	permission := sys_flag.String("u", "user", "创建用户权限")
	username := sys_flag.String("n", "", "用户名")
	password := sys_flag.String("p", "", "密码")
	sys_flag.Parse()
	if *username == "" {
		fmt.Println("请输入用户名")
		return
	}
	if *password == "" {
		fmt.Println("请输入密码")
		return
	}
	fmt.Println(*permission)
	fmt.Println(*username)
	fmt.Println(*password)
}
