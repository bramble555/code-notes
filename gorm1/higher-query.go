package main

import (
	"encoding/json"
	"fmt"
)

func HigerQuery() {
	var studentList []Student
	// 这三个结果相等，，但是底层不太一样
	// db.Where("name <> ?", "lcc").Find(&studentList)
	// db.Not("name = ?", "lcc").Find(&studentList)
	db.Order("age desc").Where("name != ?", "lcc").Find(&studentList) // 并且进行排序 Order需要在前面 asc 或者 desc
	for _, student := range studentList {
		v, _ := json.Marshal(student)
		fmt.Println(string(v))
	}
	fmt.Println("not 查询")

	// 分页查询
	studentList = []Student{}
	// offset 0和-1是一样的
	db.Limit(1).Offset(1).Find(&studentList) // SELECT * FROM `students` LIMIT 1 OFFSET 1
	for _, student := range studentList {
		v, _ := json.Marshal(student)
		fmt.Println(string(v))
	}
	fmt.Println("分页查询")
	// 去重
	studentList = []Student{}
	age := []int{}
	// 如果想把查到的去重后的年龄存放起来，需要用到Scan，不能是Find
	db.Model(studentList).Select("age").Distinct("age").Scan(&age)
	fmt.Println(age)
	fmt.Println("去重查询")
	// 查询男女生的人数以及男生姓名和女生姓名
	type AggeGroup struct {
		Count  int // `gorm:"column:count(pk_id)"`
		GName  string
		Gender int
	}
	aggeGroup := []AggeGroup{}
	// 当然也可以用db.Model.........进行操作
	db.Raw("select count(pk_id) as count,GROUP_CONCAT(name) as g_name,gender from students GROUP BY gender").Scan(&aggeGroup)
	fmt.Println(aggeGroup)
	fmt.Println("查询男女生的人数以及男生姓名和女生姓名")
	studentList = []Student{}
	// 原生
	// ("select * FROM students where age >(SELECT AVG(age) FROM  students)").Scan(&studentList)
	var avgAge float64
	db.Table("students").Select("AVG(age)").Scan(&avgAge)
	db.Model(Student{}).Where("age > ?", avgAge).Find(&studentList)
	fmt.Println(studentList)
	fmt.Println("查询大于平均年龄的学生信息")
	// 查询name=oo，或者age=23的学生
	studentList = []Student{}
	db.Where("name = @name or age = @age",
		map[string]any{"name": "oo", "age": 23}).Find(&studentList)
	fmt.Println(studentList)
	// 查询引用Scopes。假设已经查出来年龄大于18的学生，现在要查询年龄大于18的男生
	// .Scopes("年龄大于18岁学生的函数").Group(gender).......
	// Scopes相当于把一部分查询信息封装起来了
}
