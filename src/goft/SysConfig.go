package goft

import (
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
	return &SysConfig{Server: &ServerConfig{Name: "myweb", Port: 8080}}
}

func InitConfig() *SysConfig  {
	conf:=NewSysConfig()
	if b:= LoadConfigFile();b!=nil {
		err := yaml.Unmarshal(b, conf)
		if err!=nil {
			log.Fatal(err)
		}
	}
	return conf
}