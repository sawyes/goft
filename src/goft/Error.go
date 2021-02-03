package goft

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func ErrorHandler() gin.HandlerFunc {
	return func(context *gin.Context) {
		defer func() {
			if e := recover(); e != nil {
				context.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
					"error": e,
				})
			}
		}()
		context.Next()
	}
}

// defaultMsg 默认信息, 用于代码触发异常
func Error(err error, defaultMsg ...string) {
	if err == nil {
		return
	}
	errMsg := err.Error()
	if len(defaultMsg) > 0 {
		errMsg = defaultMsg[0]
	}
	panic(errMsg)
}
