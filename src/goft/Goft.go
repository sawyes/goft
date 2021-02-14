package goft

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"mygin/src/funcs"
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
		log.Println("[Goft] html FuncMap: " + config.Server.Html)
		// 自动生成funcmap
		g.FuncMap = funcs.FuncMap
		
		log.Println("[Goft] config.server.html: " + config.Server.Html)
		
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
	getCronTask().Start()
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

//Field name   | Mandatory? | Allowed values  | Allowed special characters
//----------   | ---------- | --------------  | --------------------------
//Seconds      | Yes        | 0-59            | * / , -
//Minutes      | Yes        | 0-59            | * / , -
//Hours        | Yes        | 0-23            | * / , -
//Day of month | Yes        | 1-31            | * / , - ?
//Month        | Yes        | 1-12 or JAN-DEC | * / , -
//Day of week  | Yes        | 0-6 or SUN-SAT  | * / , - ?
//
//@yearly (or @annually) | Run once a year, midnight, Jan. 1st        | 0 0 0 1 1 *
//@monthly               | Run once a month, midnight, first of month | 0 0 0 1 * *
//@weekly                | Run once a week, midnight between Sat/Sun  | 0 0 0 * * 0
//@daily (or @midnight)  | Run once a day, midnight                   | 0 0 0 * * *
//@hourly                | Run once an hour, beginning of hour        | 0 0 * * * *
func (this *Goft) Task(expr string, f func()) *Goft {
	_, err := getCronTask().AddFunc(expr, f)
	if err != nil {
		log.Println(err)
	}
	return this
}