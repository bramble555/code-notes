package test

import (
	"fmt"
	"reflect"
	"time"
)

type User struct {
	ID         int    `gorm:"column:id;primaryKey"`
	PassWd     string `json:"passwd" gorm:"column:password"`
	Name       string
	FamilyName int       `gorm:"-"`
	CreateTime time.Time `form:"create_time" binding:"required" time_format:"2006-01-02"`
}

func PrintStruct(object any) {
	tp := reflect.TypeOf(object)
	fieldNum := tp.NumField()
	for i := 0; i < fieldNum; i++ {
		field := tp.Field(i)
		fmt.Printf("%d %s offset:%d  anonymous: %t  tpye:%s exported:%t gorm tag= %s",
			i,
			field.Name,            // 变量名称
			field.Offset,          // 相对于结构体首地址的内存偏移量，string会占据16个字节
			field.Anonymous,       // 是否为匿名成员变量
			field.Type,            // 数据类型
			field.IsExported(),    // 是否可导出
			field.Tag.Get("gorm"), //得到tag里面的gorm信息
		)
		fmt.Println()
	}
}
