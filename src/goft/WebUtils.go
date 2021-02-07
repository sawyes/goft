package goft

import (
	"io/ioutil"
	"log"
	"os"
)

// 读取配置文件
func LoadConfigFile() []byte {
	getwd, err := os.Getwd()
	file := getwd + "/application.yaml"
	bytes, err:=ioutil.ReadFile(file)
	if err != nil {
		log.Println(err)
		return nil
	}
	return bytes
}
