package controllers

import (
	"BitcoinConnectionWeb/getbitcion"
	"github.com/astaxie/beego"
)

type UpTimeController struct {
	beego.Controller
}

func (u *UpTimeController) Get()  {
	u.Data["UpTime"]=getbitcion.UpTime()
	u.TplName="UpTime.html"
}
