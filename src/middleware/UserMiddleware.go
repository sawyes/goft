package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

type UserMiddleware struct {

}

func NewUserMiddleware() *UserMiddleware {
	return &UserMiddleware{}
}

func (this *UserMiddleware) OnRequest(context *gin.Context) error {
	fmt.Println("用户中间件->start")
	context.Next()
	fmt.Println("用户中间件->end")
	return nil
}