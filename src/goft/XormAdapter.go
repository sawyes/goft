package goft

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	"log"
)

type XormAdapter struct {
	*xorm.Engine
}

func NewXormAdapter() *XormAdapter {
	db, err := xorm.NewEngine("mysql",
		"root:root@(localhost:3306)/test?charset=utf8mb4&parseTime=True&loc=Local")
	
	if err != nil {
		log.Fatal(err)
	}
	
	db.DB().SetMaxIdleConns(5)
	db.DB().SetMaxOpenConns(10)
	
	return &XormAdapter{Engine: db}
}

func (this *XormAdapter) Name() string {
	return "XormAdapter"
}