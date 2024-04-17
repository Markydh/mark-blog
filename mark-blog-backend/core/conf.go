package core

import (
	"fmt"
	"mark-blog-backend/config"
	"mark-blog-backend/global"
	"os"

	"gopkg.in/yaml.v3"
)

// InitConfig 读取配置文件 settings.ymal
func InitConfig() {
	dataBytes, err := os.ReadFile("settings.yaml")
	if err != nil {
		fmt.Println("读取文件失败：", err)
		return
	}
	fmt.Println("yaml 文件的内容: \n", string(dataBytes))
	config := &config.Config{}
	err = yaml.Unmarshal(dataBytes, &config)
	if err != nil {
		fmt.Println("解析 yaml 文件失败：", err)
		return
	}
	fmt.Printf("config → %+v\n", config) // config → {Mysql:{Url:127.0.0.1 Port:3306} Redis:{Host:127.0.0.1 Port:6379}}
	mp := make(map[string]any, 2)
	err = yaml.Unmarshal(dataBytes, mp)
	if err != nil {
		fmt.Println("解析 yaml 文件失败：", err)
		return
	}
	global.Config = config
	fmt.Printf("map → %+v", config) // config → {Mysql:{Url:127.0.0.1 Port:3306} Redis:{Host:127.0.0.1 Port:6379}}
}
