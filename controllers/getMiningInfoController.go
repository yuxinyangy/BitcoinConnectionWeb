package controllers

import (
	"BitcoinConnectionWeb/getbitcion"
	"github.com/astaxie/beego"
)

type GetMiningInfoController struct {
	beego.Controller
}

func (g *GetMiningInfoController) Get()  {
	g.Data["MiningInfo"]=getbitcion.GetMiningInfo()
	g.TplName="GetMiningInfo.html"
}
