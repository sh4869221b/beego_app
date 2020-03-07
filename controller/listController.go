package controller

import "github.com/astaxie/beego"

type ListController struct {
	beego.Controller
}

func (controller *ListController) Get() {
	controller.Ctx.WriteString("list go")
}
