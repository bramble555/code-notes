package main

import (
	"encoding/json"
	"fmt"
)

type status int

// 存放数据库的时候，可能需要json类型？转换成json类型？实现某个方法，具体看枫枫blog，目前没遇到，暂不谈
// 假设我现在想在数据库节约空间，把用户状态定义成smallint(1)类型，并且返回给前端成对应的类型
type Host struct {
	ID     int
	Name   string
	Status status `gorm:"type:smallint(1)"`
}

func (s status) MarshalJSON() ([]byte, error) {
	var str string
	switch s {
	case 1:
		str = "Running"
	case 2:
		str = "Except"
	case 3:
		str = "Status?"
	}
	return json.Marshal(str)
}
func CustomTpye() {
	// 创表
	err := db.AutoMigrate(Host{})
	if err != nil {
		panic(err)
	}
	fmt.Println("Create Host succeed")
	// 插入数据
	host := Host{Name: "枫枫", Status: 2}
	err = db.Create(&host).Error
	if err != nil {
		panic(err)
	}
	data, _ := json.Marshal(host)
	fmt.Println(string(data)) // {"id":1,"name":"枫枫","status":"Running"}
}
