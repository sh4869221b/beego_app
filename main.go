package main

import (
	"github.com/astaxie/beego"
	"github.com/sh4869221b/beego_app/controller"
	"github.com/sh4869221b/beego_app/filters"
)

type mainController struct {
	beego.Controller
}

func (controller *mainController) Get() {
	controller.Ctx.WriteString("hello world!!")
}

func main() {
	beego.Router("/", &mainController{})
	beego.Router("/list.txt", &controller.ListController{})
	beego.InsertFilter("/*", beego.BeforeRouter, filters.LogManager)
	beego.Run()
}
