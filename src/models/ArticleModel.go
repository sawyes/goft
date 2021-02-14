package models

type ArticleModel struct {
	//gin文档tag方式绑定`uri`,`binding`
	//gorm tag
	NewsId int `json:"id" gorm:"column:id" uri:"id" binding:"required,gt=0"`
	NewsTitle string `json:"title" gorm:"column:newstitle"`
	NewsContent string `json:"content" gorm:"column:newscontent"`
	Views string `json:"views" gorm:"column:views"`
	Addtime string `json:"addtime" gorm:"column:addtime"`
}

func NewArticleModel() *ArticleModel {
	return &ArticleModel{}
}


func (this *ArticleModel) String() string {
	return "ArticleModel"
}