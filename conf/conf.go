package conf

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"path/filepath"
)

type Configer interface {
	GetIPAddress() (string, error)
	GetPort() (string, error)
}

type ServerConfig struct {
	Address string `json:"address"`
	Port    string `json:"port"`
}

type Config struct {
	Server ServerConfig `json:"server"`
}

func ParseJsonConfig() (config Config, err error) {
	// 获取当前运行目录
	currentDir, err := os.Getwd()
	if err != nil {
		fmt.Println("Error getting current directory:", err)
		return
	}

	fmt.Println("Current directory:", currentDir)

	// 假设 jsonFile 是指向 JSON 文件的路径 使用filepath.Join支持跨平台
	dir := "conf"
	fileName := "conf.json"
	jsonFile := filepath.Join(".", dir, fileName)

	// 打开 JSON 文件
	file, err := os.Open(jsonFile)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}
	defer file.Close()

	// 读取 JSON 文件内容
	data, err := io.ReadAll(file)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	// 解析 JSON 数据到 Config 结构体
	err = json.Unmarshal(data, &config)
	if err != nil {
		fmt.Println("Error unmarshalling JSON:", err)
		return
	}

	// 打印解析后的数据
	fmt.Printf("ParseJsonConfig Server Address: %s\n", config.Server.Address)
	fmt.Printf("ParseJsonConfig Server Port: %s\n", config.Server.Port)

	return
}

// 获取ip
func (config *Config) GetIPAddress() (string, error) {
	return config.Server.Address, nil
}

// 获取端口
func (config *Config) GetPort() (string, error) {
	return config.Server.Port, nil
}
