package goft

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"reflect"
)

var ResponderList []Responder

// 必须初始化类型
func init()  {
	ResponderList = []Responder{
		new(StringResponder),
		new(ModelGoftResponder),
		new(ModelsGoftResponder),
		new(ViewResponder),
	}
}

type Responder interface {
	RespondTo() gin.HandlerFunc
}

// 转换目标类型
func convert(handler interface{}) gin.HandlerFunc {
	h_ref := reflect.ValueOf(handler)
	for _, r := range ResponderList {
		r_val := reflect.ValueOf(r).Elem()
		if h_ref.Type().ConvertibleTo(r_val.Type()) {
			r_val.Set(h_ref)
			return r_val.Interface().(Responder).RespondTo()
		}
	}
	return nil
}

// 返回类型为字符串
type StringResponder func(*gin.Context) string
func (this StringResponder) RespondTo() gin.HandlerFunc {
	return func(context *gin.Context) {
		context.String(http.StatusOK, this(context))
	}
}

// 返回类型为Goft.Model
type ModelGoftResponder func(*gin.Context) Model
func (this ModelGoftResponder) RespondTo() gin.HandlerFunc {
	return func(context *gin.Context) {
		context.JSON(http.StatusOK, this(context))
	}
}

// 返回类型为Goft.Models
type ModelsGoftResponder func(*gin.Context) Models
func (this ModelsGoftResponder) RespondTo() gin.HandlerFunc {
	return func(context *gin.Context) {
		context.Writer.Header().Set("Content-type", "application/json")
		context.Writer.WriteString(string(this(context)))
	}
}

// 返回类型为Goft.View
type View string
type ViewResponder func(*gin.Context) View
func (this ViewResponder) RespondTo() gin.HandlerFunc {
	return func(context *gin.Context) {
		
		//obj:= map[string]string{
		//	"name":"hello world",
		//}
		
		context.HTML( http.StatusOK, string(this(context)) + ".html", context.Keys)

		log.Println("[Responder] ViewResponder keys ", fmt.Sprintf("%+v", context.Keys))
	}
}