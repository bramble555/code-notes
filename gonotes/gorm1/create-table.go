package main

import "fmt"

// gorm采用的命名策略是，表名是蛇形复数，字段名是蛇形单数
type Student struct {
	// 注意自增要放在tpye里面，否则不生效
	ID     int    `gorm:"column:pk_id;primary_key;type:bigint(20) auto_increment;comment:主键学生学号"`
	Name   string `gorm:"type:varchar(30);comment:学生名字;not null"`
	Age    int    `gorm:"typte:smallint;comment:学生年龄;not null"`
	Gender bool   `gorm:"default:true;comment:学生性别;not null"`
	// 指针类型，因为可以为空
	Email *string `gorm:"type:varchar(30);comment:学生邮箱;null"`
}

func Create() {
	db.AutoMigrate(Student{})
	fmt.Println("Create succeed")
}
