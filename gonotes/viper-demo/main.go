package main

import (
	"fmt"

	"github.com/fsnotify/fsnotify"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

type DatabaseConfig struct {
	Database string `mapstructure:"database"`
	Username string `mapstructure:"username"`
	Password string `mapstructure:"password"`
	Port     int    `mapstructure:"port"`
}

func main() {
	// 设置默认值，等级最低
	viper.SetDefault("fileDir", "./")
	viper.SetConfigName("mysql")
	viper.SetConfigType("yaml")
	// 可以设置多个路径
	viper.AddConfigPath(".")
	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			// 如果找不到配置文件，可以选择打印一条消息，但继续执行（或进行其他处理）
			fmt.Println("未找到配置文件，使用默认配置...")
			// 这里可以继续执行其他代码，或者设置一些默认配置
		} else {
			// 如果错误不是配置文件未找到，则使用 panic 或其他错误处理方式
			panic(err)
		}
	}
	fmt.Println("配置文件加载成功")
	// 监听配置
	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		// 配置文件发生变更之后会调用的回调函数
		fmt.Println("Config file changed:", e.Name)
	})
	// 当然也可以从io.Reader读取配置，使用一个切片，bytes.NewBuffer(slices) 转换为buffer类型
	viper.RegisterAlias("loud", "Verbose")
	// 注册别名（此处loud和Verbose建立了别名）
	r := gin.Default()

	r.GET("/database", func(c *gin.Context) {
		c.String(200, viper.GetString("mysql.database"))
	})
	var mysqlConfig DatabaseConfig
	// 注意读取嵌套的时候，用的方法是UnmarshalKey，而不是Unmarshal
	if err := viper.UnmarshalKey("mysql", &mysqlConfig); err != nil {
		panic(err)
	}
	fmt.Println(mysqlConfig)
	r.Run(":8080")
}
