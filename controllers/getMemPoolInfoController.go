package controllers

import (
	"BitcoinConnectionWeb/getbitcion"
	"github.com/astaxie/beego"
)

type GetMemPoolInfoController struct {
	beego.Controller
}

func (g *GetMemPoolInfoController) Get()  {
	g.Data["MemPoolInfo"]=getbitcion.GetMemPoolInfo()
	g.TplName="GetMemPoolInfo.html"
}
