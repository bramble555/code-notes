package main

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

type DatabaseConfig struct {
	Database string `yaml:"database"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	Port     int    `yaml:"port"`
}

type MySQLConfig struct {
	MySQL DatabaseConfig `yaml:"mysql"`
}

func main() {
	// 读取文件
	data, err := os.ReadFile("mysql.yaml")
	if err != nil {
		panic(err)
	}
	// 解析文件
	var config MySQLConfig
	err = yaml.Unmarshal(data, &config)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Database: %s\n", config.MySQL.Database)
	fmt.Printf("Username: %s\n", config.MySQL.Username)
	fmt.Printf("Password: %s\n", config.MySQL.Password) // 注意：在实际应用中，直接打印密码可能不安全
	fmt.Printf("Port: %d\n", config.MySQL.Port)
}
