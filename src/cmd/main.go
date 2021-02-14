package main

import (
	"fmt"
	"log"
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
		Task("@daily", func() {
			log.Println("执行定时任务")
		}).
		Lanuch()
	
}
