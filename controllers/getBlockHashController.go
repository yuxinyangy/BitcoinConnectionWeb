package controllers

import (
	"BitcoinConnectionWeb/entity"
	"BitcoinConnectionWeb/getbitcion"
	"github.com/astaxie/beego"
)


type GetBlockHashController struct {
	beego.Controller
}

func (g *GetBlockHashController) Get()  {
	g.TplName="GetBlockHashParameter.html"
}

func (g *GetBlockHashController) Post()  {
	var height entity.Heights
	err := g.ParseForm(&height)
	if err != nil {
		g.TplName = "error.html"
		return
	}
	g.Data["Height"]=height.Height
	g.Data["BlockHash"]=getbitcion.GetBlockHash(height.Height)
	g.TplName="GetBlockHash.html"
}
