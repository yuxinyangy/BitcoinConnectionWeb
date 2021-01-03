package controllers

import "github.com/astaxie/beego"

type LeftHomeController struct {
	beego.Controller
}

func (l *LeftHomeController) Get()  {
	l.TplName="lefthome.html"
}
