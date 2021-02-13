package goft

import (
	"fmt"
	"reflect"
	"strings"
)

//注解处理
type Annotations interface {
	SetTag(tag reflect.StructTag)
	String() string
}
var AnnotationList []Annotations
//判断当前的 注入对象 是否是注解
func IsAnnotation(t reflect.Type) bool {
	for _,item:=range AnnotationList{
		if reflect.TypeOf(item)==t{
			return true
		}
	}
	return false
}
func init()  {
	AnnotationList=make([]Annotations,0)
	AnnotationList=append(AnnotationList,new(Value))
}

type Value struct {
	tag reflect.StructTag
	Beanfactory *BeanFactory
}
func(this *Value) SetTag(tag reflect.StructTag){
		this.tag=tag
}
func(this *Value) String() string {
	// tag 由程序自动注入
	get_prefix:=this.tag.Get("prefix")
	if get_prefix==""{
		return ""
	}
	prefix:=strings.Split(get_prefix,".")
	if config:=this.Beanfactory.GetBean(new(SysConfig));config!=nil{
		get_value:=GetConfigValue(config.(*SysConfig).Config,prefix,0)
		if get_value!=nil{
			return fmt.Sprintf("%v",get_value)
		}else{
			return ""
		}
	}else{
		return ""
	}
	//return "21"
}

