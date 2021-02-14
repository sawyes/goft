package goft

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
)

type Goft struct {
	// 你们继承
	*gin.Engine
	g           *gin.RouterGroup
	beanFactory *BeanFactory
}

// 构建函数
func Ignite() *Goft {
	g := &Goft{Engine: gin.New(), beanFactory: NewBeanFactory()}
	g.Use(ErrorHandler()) // 强迫加载的异常处理中间件
	
	config := InitConfig()
	g.beanFactory.setBean(config) //整个配置加载进bean中
	
	// 试图配置
	if config.Server.Html != "" {
		log.Println("Html glob path: " + config.Server.Html)
		g.LoadHTMLGlob(config.Server.Html)
	}
	
	return g
}

// 启动函数
func (this *Goft) Lanuch() {
	var port int32 = 8080
	if config := this.beanFactory.GetBean(new(SysConfig)); config != nil {
		port = config.(*SysConfig).Server.Port
	}
	this.Run(fmt.Sprintf(":%d", port)) // 套接字, 暂时写死
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
		
		// 这一步是关键, 直接调用对象的build方法, 一般为路由声明
		class.Build(this)
		
		// 反射注入
		this.beanFactory.inject(class)
	}
	return this
}

// 中间件
func (this *Goft) Attach(f Fairing) *Goft {
	this.Use(func(context *gin.Context) {
		f.OnRequest(context)
	})
	return this
}

// 依赖注入, 用props保存bean对象 如:数据库连接对象
func (this *Goft) Beans(beans ...interface{}) *Goft {
	
	this.beanFactory.setBean(beans...)
	
	return this
}
