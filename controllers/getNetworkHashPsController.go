package controllers

import (
	"BitcoinConnectionWeb/entity"
	"BitcoinConnectionWeb/getbitcion"
	"github.com/astaxie/beego"
)

type GetNetworkHashPsController struct {
	beego.Controller
}

func (g *GetNetworkHashPsController) Get() {
	g.TplName="GetNetworkHashPSParameter.html"
}

func (g *GetNetworkHashPsController) Post() {
	var HashPs entity.HashPs
	err := g.ParseForm(&HashPs)
	if err != nil {
		g.TplName = "error.html"
		return
	}
	g.Data["NetworkHashPs"]=getbitcion.GetNetworkHashPS(HashPs.Blocks,HashPs.Height)
	g.TplName="GetNetworkHashPS.html"
}