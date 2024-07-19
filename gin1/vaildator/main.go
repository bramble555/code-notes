package main

import (
	"reflect"

	"github.com/gin-gonic/gin"               // 引入Gin框架
	"github.com/gin-gonic/gin/binding"       // 引入Gin的binding包，用于处理数据绑定
	"github.com/go-playground/validator/v10" // 引入go-playground的validator库，用于数据验证
)

// GetValidMsg 函数尝试从验证错误中提取并返回自定义的错误消息
func GetValidMsg(err error, obj interface{}) string {
	// obj为结构体指针
	getObj := reflect.TypeOf(obj) // 使用反射获取obj的类型
	// 断言为具体的类型，err是一个接口，这里尝试将其断言为validator.ValidationErrors类型
	if errs, ok := err.(validator.ValidationErrors); ok {
		for _, e := range errs { // 遍历所有验证错误
			if f, exist := getObj.Elem().FieldByName(e.Field()); exist { // 检查结构体中是否存在对应字段
				return f.Tag.Get("msg") // 如果存在，返回该字段的自定义错误消息
			}
		}
	}
	// 如果没有找到自定义错误消息或错误类型不是预期的，返回原始错误消息
	return err.Error()
}

// signValid 是一个自定义验证函数，检查用户名是否等于"fengfeng"
func signValid(fl validator.FieldLevel) bool {
	name := fl.Field().Interface().(string) // 获取当前验证字段的值
	return name == "fengfeng"               // 返回验证结果
}

func main() {
	router := gin.Default() // 创建一个默认的Gin引擎

	// 定义一个路由，处理POST请求到根URL
	router.POST("/", func(c *gin.Context) {
		// 定义一个UserInfo结构体，用于接收前端传来的JSON数据
		type UserInfo struct {
			Name string `json:"name" binding:"sign" msg:"用户名错误"` // 字段Name绑定自定义验证器"sign"，并设置自定义错误消息
			Age  int    `json:"age" binding:""`                  // 字段Age不绑定任何验证器
		}

		var user UserInfo              // 创建一个UserInfo变量用于存储绑定后的数据
		err := c.ShouldBindJSON(&user) // 将请求体中的JSON数据绑定到user变量上
		if err != nil {
			// 如果绑定失败（即验证失败），则调用GetValidMsg获取错误消息
			msg := GetValidMsg(err, &user)
			c.JSON(200, gin.H{"msg": msg}) // 将错误消息以JSON格式返回给客户端
			return
		}
		c.JSON(200, user) // 如果绑定成功，将user变量以JSON格式返回给客户端
	})

	// 注册自定义验证器"sign"到Gin的验证器引擎中
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("sign", signValid)
	}

	router.Run(":8080") // 启动HTTP服务器，监听8080端口
}
