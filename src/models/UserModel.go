package models

import "fmt"

type UserModel struct {
	//gin文档tag方式绑定`uri`,`binding`
	//gorm tag
	UserId int `gorm:"column:id" uri:"id" binding:"required,gt=0"`
	UserName string `gorm:"column:user_name"`
}

func NewUserModel() *UserModel {
	return &UserModel{}
}

func (this *UserModel) String() string {
	return fmt.Sprintf("userid:%d, username:%s", this.UserId, this.UserName)
}