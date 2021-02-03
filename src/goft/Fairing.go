package goft

import "github.com/gin-gonic/gin"

// 规范中间件
type Fairing interface {
	OnRequest(context *gin.Context) error
}
