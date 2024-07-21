package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

const (
	user     = "root"
	password = "123456"
	dbName   = "go_test"
)

var db *sql.DB
var err error

func init() {
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
}
func main() {
	tranUpdate()
}
func tranUpdate() {
	// 更新事务，如果有err，回滚
	sqlStr := "update student set age = 66 where id = ?"
	tx, err := db.Begin()
	if err != nil {
		if tx != nil {
			tx.Rollback()
		}
		panic(err)
	}
	stmt, err := tx.Prepare(sqlStr)
	if err != nil {
		tx.Rollback()
		panic(err)
	}
	defer stmt.Close()
	res, err := stmt.Exec(1)
	if err != nil {
		tx.Rollback()
		panic(err)
	}
	n, err := res.RowsAffected()
	if err != nil {
		tx.Rollback()
		panic(err)
	}
	if n != 1 {
		tx.Rollback()
		fmt.Println("更新不是一行，已经回滚")
	}
	// 如果没有问题提交事务
	if err := tx.Commit(); err != nil {
		fmt.Println("Failed to commit transaction:", err)
		return
	}
	fmt.Println("update succeed")
}
