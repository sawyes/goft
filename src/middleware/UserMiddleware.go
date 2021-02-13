package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
)

type UserMiddleware struct {
}

func NewUserMiddleware() *UserMiddleware {
	return &UserMiddleware{}
}

func (this *UserMiddleware) OnRequest(context *gin.Context) error {
	
	params, _ := ioutil.ReadAll(context.Request.Body)
	
	fmtString := fmt.Sprintf("[UserMiddleware] Method: %s, Uri: %s, Params: %s",
		context.Request.Method,
		context.Request.URL,
		string(params),
	)
	
	fmt.Println(fmtString)
	
	context.Next()
	
	return nil
}