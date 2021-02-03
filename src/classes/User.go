package classes

import (
	"github.com/gin-gonic/gin"
	"mygin/src/goft"
	"mygin/src/models"
)

type User struct {
	*goft.GormAdapter
}

func NewUserClasses() *User {
	return &User{}
}

func (this *User) Build(goft *goft.Goft) {
	goft.Handle("GET", "/user", this.UserLists)
	goft.Handle("GET", "/user/detail/:id", this.UserDetail)
}

func (this *User) UserLists(context *gin.Context) goft.Models {
	users:=[]*models.UserModel{
		&models.UserModel{UserId: 101, UserName: "a"},
		&models.UserModel{UserId: 102, UserName: "b"},
	}
	return goft.Makemodels(users)
}

func (this *User) UserDetail(ctx *gin.Context) goft.Model {
	//return &models.UserModel{
	//	UserId: 13,
	//	UserName: "steven",
	//}
	user:=models.NewUserModel()
	
	// 原理, 通过反射匹配,遍历请求参数的key和user的tag(反射获得)值匹配,
	// 那么把请求参数的值复制到目标对象
	err := ctx.BindUri(user)
	
	goft.Error(err, "参数不合法")
	this.Table("user").Where(user.UserId).
		Find(user)
	return user
}