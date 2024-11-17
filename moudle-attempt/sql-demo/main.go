package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)
const (
	user     = "root"
	password = "123456"
	database = "t"
)

// 用来操作数据库，并发安全
// 不需要进行关闭
func migration() {
	migStr := fmt.Sprintf("%s:%s@tcp(127.0.0.1:3306)/t?multiStatements=true",
		user, password,
	)
	db, err := sql.Open("mysql", migStr)
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	driver, err := mysql.WithInstance(db, &mysql.Config{})

	if err != nil {
		log.Fatal(err)
	}

	m, err := migrate.NewWithDatabaseInstance(
		"file://migration",
		"t",
		driver,
	)
	if err != nil {
		log.Fatal(err)
	}
	err = m.Up() //or m.Down()
	if err != nil {
		log.Fatal(err)
	}
	_ = m.Steps(1) //执行的文件数
}
func main() {
	// // 连接字符串，也就是数据源
	// connStr := fmt.Sprintf("%s:%s@tcp(127.0.0.1:3306)/%s",
	// 	user, password, database,
	// )

	// // 不会连接数据库，不会进行验证参数。只是把连接到数据库的struct给设置了
	// // 真正的连接是被需要的时候才进行懒设置的
	// db, err := sql.Open("mysql", connStr)
	// if err != nil {
	// 	log.Fatalln(err)
	// }
	// defer db.Close()
	// // context.Context这个类型可以携带截止时间，取消信号
	// ctx := context.Background() // 此函数连接的时候，不会被取消，也没有截止时间
	// err = db.PingContext(ctx)
	// if err != nil {
	// 	log.Fatalln(err)
	// }
	// fmt.Println("Connected")
	migration()

}
