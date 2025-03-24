package controllers

import (
	"log"
	"strconv"

	"github.com/astaxie/beego"
)

type ApiController struct {
	beego.Controller
}

type Response struct {
	Data interface{} `json:"data"`
	Metadata interface{} `json:"metadata"`
	}

func (c *ApiController) Get() {
	c.Data["json"] = map[string]string{"message": "Hello, world!"}
	c.ServeJSON()
}

func (c *ApiController) Post() {
	var (
		dam []string
		res Response
	)

	for x:=0;x<4;x++{
		dam = append(dam, strconv.Itoa(x))
	}

	log.Println("dam",dam)

	res.Data = dam

	c.Data["json"] = res
	c.ServeJSON()
}