package goft

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"log"
)

type ServerConfig struct {
	Name string
	Port int32
}

// 系统配置 
type SysConfig struct {
	Server *ServerConfig
}

func NewSysConfig() *SysConfig {
	// 默认配置
	return &SysConfig{Server: &ServerConfig{Name: "myweb", Port: 8080}}
}

// 初始化配置, 存在配置文件则用配置, 否则加载默认
func InitConfig() *SysConfig  {
	conf:=NewSysConfig()
	if b:= LoadConfigFile();b!=nil {
		
		err := yaml.Unmarshal(b, conf)
		if err!=nil {
			log.Println("SysConfig load yaml failed")
			log.Fatal(err)
		}
		
		log.Println("SysConfig load success")
		log.Println(fmt.Sprintf("%+v", *conf.Server))
	}
	return conf
}