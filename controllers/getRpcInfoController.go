package controllers

import (
	"BitcoinConnectionWeb/getbitcion"
	"github.com/astaxie/beego"
)

type GetRpcInfoController struct {
	beego.Controller
}

func (g *GetRpcInfoController) Get()  {
	g.Data["RpcInfo"]=getbitcion.GetRpcInfo()
	g.TplName="GetRpcInfo.html"
}
