package goft

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"log"
)

// type定义struct变量必须以大写字母开头，否则会出现key解析不出来
// 此处定义的是项目启动的名字和端口
type ServerConfig struct {
	Name string
	Port int32
	Html string
}

// 系统配置 
type SysConfig struct {
	//接收yaml中server的配置
	Server *ServerConfig
	
	//接收yaml中config的配置
	Config UserConfig
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

type UserConfig map[interface{}]interface{}
//递归读取用户配置文件
func GetConfigValue(m UserConfig,prefix []string,index int) interface{}  {
	key:=prefix[index]
	if v,ok:=m[key];ok{
		if index==len(prefix)-1{ //到了最后一个
			return v
		}else{
			index=index+1
			if mv,ok:=v.(UserConfig);ok{ //值必须是UserConfig类型
				return GetConfigValue(mv,prefix,index)
			}else{
				return  nil
			}
			
		}
	}
	return  nil
}