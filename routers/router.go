package routers

import (
	"firstbeego/controllers"

	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
	beego.Router("/api/hello", &controllers.ApiController{})
	beego.Router("/api/testing", &controllers.ApiController{})
	beego.Router("/api/users", &controllers.UserController{})
    beego.Router("/api/users/:id", &controllers.UserController{}, "get:GetById")
    beego.Router("/api/users/create", &controllers.UserController{}, "post:Post")
}
