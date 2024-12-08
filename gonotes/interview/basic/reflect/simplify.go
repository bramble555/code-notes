package reflect

import (
	"fmt"
	"os"
	"reflect"
	"strings"
)

// 反射如何简化代码
// 极客兔兔
type Config struct {
	// 注释是配置中对应的环境变量
	Name    string `json:"server-name"` // CONFIG_SERVER_NAME
	IP      string `json:"server-ip"`   // CONFIG_SERVER_IP
	URL     string `json:"server-url"`  // CONFIG_SERVER_URL
	Timeout string `json:"timeout"`     // CONFIG_TIMEOUT
}

func writeConfig() {
	// os.Setenv 只改变当前 Go 程序进程内的环境变量设置
	os.Setenv("CONFIG_SERVER_NAME", "global_server")
	os.Setenv("CONFIG_SERVER_IP", "10.0.0.1")
	os.Setenv("CONFIG_SERVER_URL", "geektutu.com")
}
func ReadConfig() *Config {
	writeConfig()
	config := Config{}
	// 获取reflect.Config
	typ := reflect.TypeOf(config)
	// reflect.Indirect需要传入指针类型
	value := reflect.Indirect(reflect.ValueOf(&config))
	for i := 0; i < typ.NumField(); i++ {
		f := typ.Field(i)
		if v, ok := f.Tag.Lookup("json"); ok {
			key := fmt.Sprintf("CONFIG_%s", strings.ReplaceAll(strings.ToUpper(v), "-", "_"))
			if envVal, exist := os.LookupEnv(key); exist {
				// env 和 reflect.ValueOf(env) 打印出来的是一样的，但是类型不一样 env是string，后者是reflect.Value
				value.Field(i).Set(reflect.ValueOf(envVal))
				// 下面根据name查找性能会更差一些
				// value.FieldByName(f.Name).Set(reflect.ValueOf(envVal))
			}
		}
	}
	return &config
}
