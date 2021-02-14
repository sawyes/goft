package classes

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"mygin/src/goft"
)

type Index struct {
}

// 接收外部的gin.Engine, 构造IndexClasses类
func NewIndexClasses() *Index {
	return &Index{}
}

// 注册路由方法
func (this *Index) Build(goft *goft.Goft) {
	
	goft.Handle("GET", "/index", this.GetIndex)
	goft.Handle("GET", "/hello", this.GetHello)
	goft.Handle("GET", "/expr", this.GetExpr)
	goft.Handle("GET", "/expr2", this.GetExpr2)
}

// gin业务逻辑实现
func (this *Index) GetIndex(context *gin.Context) goft.View {
	context.Set("name", "hello varriable")
	return "index"
}

func (this *Index) GetHello(context *gin.Context) string {
	return "world"
}

// 可比较表达式
func (this *Index) GetExpr(context * gin.Context) string {
	expr, _ := goft.ExecExpr("gt .age 19", map[string]interface{}{
		"age": 19,
	})
	
	return fmt.Sprintf("%v", expr)
}

// 可比较表达式
func (this *Index) GetExpr2(context * gin.Context) string {
	// 模板调用变量对象方法
	expr, _ := goft.ExecExpr(".method.GetAge", map[string]interface{}{
		"method": &Me{},
	})
	
	return fmt.Sprintf("%v", expr)
}

type Me struct {
}
func (this *Me) GetAge() int {
	return 21
}

func (this *Index) Name() string {
	return "IndexClass"
}