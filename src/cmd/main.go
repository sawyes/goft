package main

import (
	. "mygin/src/classes"
	"mygin/src/goft"
	. "mygin/src/middleware"
)

func main() {
	goft.Ignite().
		DB(goft.NewGormAdapter()).
		Attach(NewUserMiddleware()).
		Mount("/v1", NewIndexClasses(), NewUserClasses()).
		Lanuch()
	
}
