package main

import (
	"firstbeego/models"
	_ "firstbeego/routers"

	_ "github.com/go-sql-driver/mysql"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)


func init() {
	// Register the database
	orm.RegisterDataBase("default", "mysql", "root:@tcp(127.0.0.1:3306)/test?charset=utf8")

	// Enable ORM debugging
	orm.Debug = true

	// Register models
	orm.RegisterModel(new(models.User))

	// Perform Auto Migration to sync your model with the database schema
	orm.RunSyncdb("default", false, true)
}

func main() {

	beego.Run()
}

