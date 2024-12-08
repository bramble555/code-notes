package main

import (
	"encoding/json"
	"fmt"
)

func Insert() {
	// 插入lcc
	email := "999333.com"
	// ID可以自增
	studenthope := Student{
		Name:   "lwx",
		Age:    33,
		Gender: false,
		Email:  &email,
	}
	// 当有钩子函数，需要传入结构体指针，否则会报错
	// invalid value, should be pointer to struct or slice
	err := db.Create(studenthope).Error
	panic(err)
	// studentlcc := Student{
	// 	ID:     1,
	// 	Name:   "lcc",
	// 	Age:    18,
	// 	Gender: true,
	// 	Email:  nil,
	// }
	// db.Create(studentlcc)
	// // 插入lc
	// lcEmail := "3478034126qq.com"
	// studentlc := Student{
	// 	ID:     2,
	// 	Name:   "lc",
	// 	Age:    19,
	// 	Gender: true,
	// 	Email:  &lcEmail,
	// }
	// db.Create(studentlc)
	// // 插入ck
	// ckEmail := "8881515606qq.com"
	// studentck := Student{
	// 	ID:     4,
	// 	Name:   "jack",
	// 	Age:    16,
	// 	Gender: true,
	// 	Email:  &ckEmail,
	// }
	// err := db.Create(studentck).Error
	// if err == nil {
	// 	fmt.Println("插入成功")
	// }
}

// 查询一条数据
func QueryRow() {
	// var student Student
	// SELECT * FROM `students` LIMIT 1
	// db.Take(&student)
	// fmt.Println(student)
	// SELECT * FROM `students` ORDER BY `students`.`pk_id` LIMIT 1
	// db.First(&student)
	// fmt.Println(student)
	// SELECT * FROM `students` ORDER BY `students`.`pk_id` DESC LIMIT 1
	// db.Last(&student)
	// fmt.Println(*student.Email) // 是指针类型的需要用指针去获取值
	// Take 默认传入主键，数字或者字符串都行
	// db.Take(&student,"2")
	// fmt.Println(student)
	// Take 当然也可以传入where条件
	// db.Take(&student, "name = ? ", "lcc") // 不能自己用Sprintf去拼接，这样会产生sql注入问题
	// fmt.Println(student)
	// 可以根据结构体查询,只能设置主键，不能设置其他的，比如Age
	var student1 = Student{}
	student1.ID = 3
	db.Take(&student1)
	fmt.Println(student1)
}

// 查询多条数据
func Query() {
	var studentList []Student
	// 无where
	db.Find(&studentList)
	// 根据主键查询
	// db.Find(&studentList,1,3) // 查询1和3
	// 简单查询
	// db.Find(&studentList, "name = ? or age = ?", "lcc",16)
	// 复杂查询(就是需要用到一个列表)，其他例子去官网查看
	// db.Where("name IN ?", []string{"lcc", "lc"}).Find(&studentList)
	for _, student := range studentList {
		// 由于email是指针类型，所以看不到实际的内容
		// 但是序列化之后，会转换为我们可以看得懂的方式
		data, _ := json.Marshal(student)
		fmt.Println(string(data))
	}
}
func Update() {
	// 保存所有字段,先找到某个student，还是更改
	var student Student
	db.Find(&student, 3)
	// data,_ := json.Marshal(student)
	// fmt.Println(string(data))
	student.Gender = false
	db.Save(&student)
	// 这个就相当于插入数据了，id自增，不用输入，email可以为空
	db.Save(&Student{Name: "kenzhu", Age: 10})
	db.Save(&Student{Name: "jinjiao", Age: 15})
	// 这里有一个问题，就是我Age设置的不能为空，我没有进行保存，默认给了0
	db.Save(&Student{Name: "liming"})
	// 其他的看官网，多条更新,仍然先找，现在要更新未成年 为成年
	var studentList []Student
	// 不能用Save,用Update
	db.Where("age < ?", 18).Find(&studentList).Update("age", 18)
	student = Student{}
	// map[string]interface{}可以更改成默认字段，如果是结构体，不能
	db.Model(&student).Where("pk_id = ?", 4).Updates(map[string]interface{}{"name": "hong", "age": 0})
}
func Delete() {
	// 删除最后一行
	var student Student
	db.Last(&student)
	deleteID := student.ID
	fmt.Println(deleteID)
	// student = Student{} // 这行可有有无
	db.Delete(&student, deleteID)
}
