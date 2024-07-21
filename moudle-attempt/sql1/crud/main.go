package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

const (
	username = "root"
	password = "123456"
	dbname   = "go_test"
)

var db *sql.DB
var err error

type student struct {
	id   int
	name string
	age  int
}

func init() {
	dataSourceName := fmt.Sprintf("%s:%s@tcp(127.0.0.1)/%s",
		username, password, dbname)
	db, err = sql.Open("mysql", dataSourceName)
	if err != nil {
		panic(err)
	}
	err = db.Ping()
	if err != nil {
		panic(err)
	}
	fmt.Println("Connected")
}

func main() {
	delete()
	query()
}

// 在表中插入数据
func insert() {
	sqlStr := "insert into student values(?,?,?)"
	// 插入、更新和删除操作都使用Exec方法。
	_, err = db.Exec(sqlStr, 3, "ke", 17)
	if err != nil {
		panic(err)
	}
	fmt.Println("Insert succeed")
}

// 查询一行数据
func queryRow() {
	sqlStr := "select id,name,age from student where id=?"
	var s student
	// 非常重要：确保QueryRow之后调用Scan方法，否则持有的数据库链接不会被释放
	err = db.QueryRow(sqlStr, 1).Scan(&s.id, &s.name, &s.age)
	if err != nil {
		panic(err)
	}
	fmt.Println("student id:", s.id, "|student name:", s.name, "|student age:", s.age)
}
func query() {
	sqlStr := "select id,name,age from student"

	rows, err := db.Query(sqlStr)
	if err != nil {
		panic(err)
	}
	// 非常重要：关闭rows释放持有的数据库链接
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
func update() {
	sqlStr := "update student set age = 80 where id = ?"
	res, err := db.Exec(sqlStr, 1)
	if err != nil {
		panic(err)
	}
	// 返回影响的行数
	n, err := res.RowsAffected()
	if err != nil {
		panic(err)
	}
	fmt.Println(n == 1)
}
func delete() {
	sqlStr := "delete from student where id = ?"
	res, err := db.Exec(sqlStr, 3)
	if err != nil {
		panic(err)
	}
	// 返回影响的行数
	n, err := res.RowsAffected()
	if err != nil {
		panic(err)
	}
	fmt.Println(n == 1)
}
