package controllers

import (
	"BitcoinConnectionWeb/getbitcion"
	"github.com/astaxie/beego"
)

type GetBestBlockHashController struct {
	beego.Controller
}

func (g *GetBestBlockHashController) Get()  {
	g.Data["BestHash"] = getbitcion.GetBestBlockHash()
	g.TplName="GetBestBlockHash.html"
}
