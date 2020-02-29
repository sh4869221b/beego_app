package main

import (
	"github.com/astaxie/beego"
	"github.com/sh4869221b/beego_app/filters"
)

type MainController struct {
	beego.Controller
}

func (this *MainController) Get() {

	this.Ctx.WriteString("hello world")
}

func main() {

	beego.Router("/", &MainController{})
	beego.InsertFilter("/*", beego.BeforeRouter, filters.LogManager)
	beego.Run()
}
