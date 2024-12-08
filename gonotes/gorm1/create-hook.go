package main

import (
	"gorm.io/gorm"
)

// 增删改查都有钩子函数
// 在每次插入之前，我希望做点事情？，让每次增加不用输入ID
func (u *Student) BeforeCreate(tx *gorm.DB) (err error) {

	return nil
}
