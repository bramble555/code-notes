package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	_ "github.com/go-sql-driver/mysql"
	// 想要连接到SQL数据库，首先需要加载目标数据库驱动，驱动里面包含者与数据库交互的逻辑
	// 一般使用sql.Register()函数注册数据库驱动,github.com/denisenkom/go-mssqldb里面的init注册了
)

const (
	user     = "root"
	password = "123456"
	database = "go_test"
)

// 用来操作数据库，并发安全
// 不需要进行关闭

func main() {
	// 连接字符串，也就是数据源
	connStr := fmt.Sprintf("%s:%s@tcp(127.0.0.1:3306)/%s",
		user, password, database,
	)
	
	// 不会连接数据库，不会进行验证参数。只是把连接到数据库的struct给设置了
	// 真正的连接是被需要的时候才进行懒设置的
	db, err := sql.Open("mysql", connStr)
	if err != nil {
		log.Fatalln(err)
	}
	defer db.Close()
	// context.Context这个类型可以携带截止时间，取消信号
	ctx := context.Background() // 此函数连接的时候，不会被取消，也没有截止时间
	err = db.PingContext(ctx)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println("Connected")

}
