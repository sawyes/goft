package classes

import (
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
}

// gin业务逻辑实现
func (this *Index) GetIndex(context *gin.Context) goft.View {
	return "index"
}

func (this *Index) GetHello(context *gin.Context) string {
	return "world"
}