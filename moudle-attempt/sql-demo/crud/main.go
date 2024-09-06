package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

const (
	username = "root"
	password = "123456"
	dbname   = "store"
)

var db *sql.DB
var err error

type customer struct {
	customerID int
	firstName  string
	birthDate  string
	phone      string
}

func initMysql() (err error) {
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
	return
}

func main() {
	err := initMysql()
	if err != nil {
		panic(err)
	}
	defer db.Close()
	queryRow()
	query()
	fmt.Println()
	// insert()
	update()
	delete()

}

// 在表中插入数据
func insert() {
	sqlStr := "insert into order_statuses (order_status_id,name) values(?,?)"
	// 插入、更新和删除操作都使用Exec方法。
	_, err = db.Exec(sqlStr, 4, "ck")
	if err != nil {
		panic(err)
	}
	fmt.Println("Insert succeed")
}

// 查询一行数据
func queryRow() {
	sqlStr := "select customer_id,first_name,birth_date,phone from customers where customer_id=?"
	var c customer
	// 非常重要：确保QueryRow之后调用Scan方法，否则持有的数据库链接不会被释放
	err = db.QueryRow(sqlStr, 1).Scan(&c.customerID, &c.firstName, &c.birthDate, &c.phone)

	if err != nil {
		panic(err)
	}
	fmt.Println("customerID:", c.customerID, "|firstName:", c.firstName, "|birthDate:", c.birthDate, "|phone:", c.phone)
}
func query() {
	sqlStr := "select customer_id, first_name, birth_date, phone from customers"

	rows, err := db.Query(sqlStr)
	if err != nil {
		panic(err)
	}
	// 需要释放查询
	defer rows.Close()

	for rows.Next() {
		var c customer
		var phone sql.NullString
		err = rows.Scan(&c.customerID, &c.firstName, &c.birthDate, &phone)
		if err != nil {
			panic(err)
		}
		if phone.Valid {
			c.phone = phone.String // 直接赋值给 c.phone（现在是 string 类型）
		}
		fmt.Printf("customerID: %d\t | firstName: %s\t | birthDate: %s\t | phone: %s\n", c.customerID, c.firstName, c.birthDate, c.phone)
	}
}
func update() {
	sqlStr := "UPDATE order_statuses SET name = 'dadadad' WHERE order_status_id = ?;"
	res, err := db.Exec(sqlStr, 4)
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
	sqlStr := "delete from order_statuses where order_status_id = ?"
	res, err := db.Exec(sqlStr, 4)
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
