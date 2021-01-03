package controllers

import (
	"BitcoinConnectionWeb/getbitcion"
	"github.com/astaxie/beego"
)

type GetTxOutSetInfoController struct {
	beego.Controller
}

func (g *GetTxOutSetInfoController) Get()  {
	g.Data["TxOutSetInfo"]=getbitcion.GetTxOutSetInfo()
	g.TplName="GetTxOutSetInfo.html"
}
