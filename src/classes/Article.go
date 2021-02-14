package classes

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"log"
	"mygin/src/goft"
	"mygin/src/models"
	"strconv"
)

type ArticleClass struct {
	*goft.GormAdapter
}
func NewArticleCalss() *ArticleClass {
	return &ArticleClass{}
}
//php tp
func(this *ArticleClass) ArticleDetail(ctx *gin.Context) goft.Model{
	news:=models.NewArticleModel()
	goft.Error(ctx.ShouldBindUri(news))
	goft.Error(this.Table("mynews").Where("id=?",news.NewsId).Find(news).Error)

	goft.Task(this.UpdateViews, func() {
		log.Println("ok id:" + strconv.Itoa(news.NewsId))
	} ,news.NewsId) //代表执行一个协程任务
	
	return news
}
//增加点击量
func(this *ArticleClass) UpdateViews(params ...interface{}){
	this.Table("mynews").Where("id=?",params[0]).
		Update("views",gorm.Expr("views+1"))
}


func(this *ArticleClass)  Build(goft *goft.Goft){
	goft.Handle("GET","/article/:id",this.ArticleDetail)
}
