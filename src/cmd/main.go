package main

import (
	"fmt"
	. "mygin/src/classes"
	"mygin/src/goft"
	. "mygin/src/middleware"
)

func main() {
	
	fmt.Println(goft.InitConfig().Config)
	// map[user:map[score:100]]
	//return
	goft.GenTplFunc("src/funcs")
	
	
	goft.Ignite().
		Beans(goft.NewGormAdapter(), goft.NewXormAdapter()).
		Attach(NewUserMiddleware()).
		Mount("/v1",
			NewIndexClasses(),
			NewUserClasses(),
			NewArticleCalss(),
		).
		//Task("0/3 * * * * *", func() {
		//	log.Println("执行定时任务")
		//}).
		//Task("0/3 * * * * *", goft.Expr(".ArticleClass.Test")).
		Lanuch()
	
}
