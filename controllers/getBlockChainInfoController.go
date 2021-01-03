package controllers

import (
	"BitcoinConnectionWeb/getbitcion"
	"github.com/astaxie/beego"
)

type GetBlockChainInfoController struct {
	beego.Controller
}

func (g *GetBlockChainInfoController) Get()  {
	g.Data["BlockChainInfo"] = getbitcion.GetBlockChainInfo()
	g.TplName="GetBlockChainInfo.html"
}
