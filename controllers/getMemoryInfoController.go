package controllers

import (
	"BitcoinConnectionWeb/getbitcion"
	"github.com/astaxie/beego"
)

type GetMemoryInfoController struct {
	beego.Controller
}

func (g *GetMemoryInfoController) Get()  {
	g.Data["MemoryInfo"]=getbitcion.GetMemoryInfo()
	g.TplName="GetMemoryInfo.html"
}
