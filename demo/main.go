package main

import (
	"log"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type MODEL struct {
	ID        uint      `gorm:"primarykey" json:"id" structs:"-"` //主键
	CreatedAt time.Time `json:"created_at" structs:"-"`           //创建时间
	UpdatedAt time.Time `json:"-" structs:"-"`                    //更新时间
}
type UserModel struct {
	//gorm.Model 需要逻辑删除时可用
	MODEL
	NickName string `gorm:"size:36" json:"nick_name"` //昵称
	UserName string `gorm:"size:36" json:"user_name"`
	PassWord string `gorm:"size:128" json:"-"`
	Avatar   string `gorm:"sign:256" json:"avatar_id"`
	Email    string `gorm:"sign:128" json:"email"`
	Tel      string `gorm:"size:18" json:"tel"`
	Addr     string `gorm:"size:64" json:"addr"`
	Token    string `gorm:"size:64" json:"token"`
	IP       string `gorm:"size:20" json:"ip"`
}
type UserCollectModel struct {
	MODEL
	UserID    uint      `gorm:"primaryKey"`
	UserModel UserModel `gorm:"foreignKey:UserID"`
	ArticleID uint      `gorm:"primaryKey"`
	CreatedAt time.Time
}

func main() {
	dsn := "root:123456@tcp(127.0.0.1:3306)/test_gvb"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalln(err)
	}
	// db.SetupJoinTable(&UserModel{}, "CollectsModels", &UserCollectModel{})
	db.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(&UserModel{}, &UserCollectModel{})
}
