package controllers

import "github.com/astaxie/beego"

type RightHomeController struct {
	beego.Controller
}

func (r *RightHomeController) Get()  {
	r.TplName="righthome.html"
}
