package controllers

import (
	"BitcoinConnectionWeb/getbitcion"
	"github.com/astaxie/beego"
)

type GetBlockCountController struct {
	beego.Controller
}

func (g *GetBlockCountController) Get()  {
	g.Data["BlockCount"]=getbitcion.GetBlockCount()
	g.TplName="GetBlockCount.html"
}
