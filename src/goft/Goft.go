package goft

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"reflect"
)

type Goft struct {
	// 你们继承
	*gin.Engine
	g *gin.RouterGroup
	//dba interface{}
	props []interface{}
}

// 构建函数
func Ignite() *Goft {
	goft := &Goft{Engine: gin.New(), props: make([]interface{}, 0)}
	goft.Use(ErrorHandler())
	return goft
}

// 启动函数
func (this *Goft) Lanuch() {
	this.Run(":9888") // 套接字, 暂时写死
}

// 根据匿名继承, 重载Handle
func (this *Goft) Handle(httpMethod, relativePath string, handler interface{}) *Goft {
	if h := convert(handler); h != nil {
		this.g.Handle(httpMethod, relativePath, h)
	} else {
		fmt.Println(relativePath)
	}
	return this
}

// 加载分组和设置路由等信息
func (this *Goft) Mount(group string, classes ...IClass) *Goft {
	this.g = this.Group(group)
	for _, class := range classes {
		class.Build(this) // 这一步是关键, 直接调用对象的build方法注册函数
		
		// 注入数据源
		this.setProp(class)
	}
	return this
}

// 注入数据源
func (this *Goft) setProp(class IClass) {
	vClass := reflect.ValueOf(class).Elem()
	
	// 遍历IClass中的所有属性, 存在ORM数据源且为nil则注入
	for i := 0; i < vClass.NumField(); i++ {
		
		// 反射定位到属性
		f := vClass.Field(i)
		
		// 目标对象必须为指针, 且为nil
		if !f.IsNil() || f.Kind() != reflect.Ptr {
			continue
		}
		
		// 检查是否为ORM数据源, 是的话自动注入
		if p := this.getProp(f.Type()); p != nil {
			f.Set(reflect.New(f.Type().Elem()))
			f.Elem().Set(reflect.ValueOf(p).Elem())
		}
	}
}

// 如果Goft props存在对应的类型, 那么返回该类型
func (this *Goft) getProp(t reflect.Type) interface{} {
	for _, p := range this.props {
		if t == reflect.TypeOf(p) {
			return p
		}
	}
	return nil
}

// 中间件
func (this *Goft) Attach(f Fairing) *Goft {
	this.Use(func(context *gin.Context) {
		f.OnRequest(context)
	})
	return this
}

// 依赖注入, 如:数据库连接对象
func (this *Goft) Beans(beans ...interface{}) *Goft {

	this.props = append(this.props, beans...)
	
	return this
}
