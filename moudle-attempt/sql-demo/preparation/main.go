package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

// 1.什么是预处理？
// 把SQL语句分成两部分，命令部分与数据部分。
// 先把命令部分发送给MySQL服务端，MySQL服务端进行SQL预处理。
// 然后把数据部分发送给MySQL服务端，MySQL服务端对SQL语句进行占位符替换。
// MySQL服务端执行完整的SQL语句并将结果返回给客户端
// 2.预处理可以解决
// 优化MySQL服务器重复执行SQL的方法，可以提升服务器性能，提前让服务器编译，一次编译多次执行，节省后续编译的成本。
// 避免SQL注入问题
const (
	user     = "root"
	password = "123456"
	dbName   = "go_test"
)

var db *sql.DB
var err error

type student struct {
	id   int
	name string
	age  int
}

func initMysql() (err error) {
	dns := fmt.Sprintf("%s:%s@tcp(127.0.0.1:3306)/%s",
		user, password, dbName)
	db, err = sql.Open("mysql", dns)
	if err != nil {
		panic(err)
	}
	err = db.Ping()
	if err != nil {
		panic(err)
	}
	fmt.Println("Connected")
	return
}
func main() {
	err := initMysql()
	if err != nil {
		panic(err)
	}
	defer db.Close()
	preInsert()
	preQuery()
}
func preQuery() {
	sqlStr := "select * from student "
	stmt, err := db.Prepare(sqlStr)
	if err != nil {
		panic(err)
	}
	defer stmt.Close() // 使用defer来确保资源被关闭
	// 查询用Query(),其他用Exec()
	rows, err := stmt.Query()
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	for rows.Next() {
		var s student
		err = rows.Scan(&s.id, &s.name, &s.age)
		if err != nil {
			panic(err)
		}
		fmt.Println("student id:", s.id, "|student name:", s.name, "|student age:", s.age)
	}
}
func preInsert() {
	sqlStr := "insert into student values (?,?,?)"
	stmt, err := db.Prepare(sqlStr)
	if err != nil {
		panic(err)
	}
	defer stmt.Close() // 使用defer来确保资源被关闭
	// 查询，删除，插入用Exec
	res, err := stmt.Exec(3, "oo", 20)
	if err != nil {
		panic(err)
	}
	n, err := res.RowsAffected()
	if err != nil {
		panic(err)
	}
	fmt.Println("插入成功，影响了", n, "行")

	// 我又来进行插入啦
	_, err = stmt.Exec(5, "小王子", 16)
	if err != nil {
		fmt.Printf("insert failed, err:%v\n", err)
		return
	}

}
