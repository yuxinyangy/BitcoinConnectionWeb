package controllers

import (
	"BitcoinConnectionWeb/entity"
	"BitcoinConnectionWeb/getbitcion"
	"github.com/astaxie/beego"
)


type GetBlockController struct {
	beego.Controller
}

func (g *GetBlockController) Get() {
	g.TplName="GetBlockParameter.html"
}

func (g *GetBlockController) Post()  {
	var hash entity.HASH
	err := g.ParseForm(&hash)
	if err != nil {
		g.TplName = "error.html"
		return
	}
	g.Data["Block"]=getbitcion.GetBlock(hash.Hash)
	if getbitcion.GetBlock(hash.Hash).Previousblockhash == "" {
		g.Data["Previousblockhash"]="这是创世区块，无上一区块的hash"
	}else{
		g.Data["Previousblockhash"]=getbitcion.GetBlock(hash.Hash).Previousblockhash
	}
	if getbitcion.GetBlock(hash.Hash).Nextblockhash == "" {
		g.Data["Nextblockhash"]="这是最高区块，无下一区块的hash"
	}else{
		g.Data["Nextblockhash"]=getbitcion.GetBlock(hash.Hash).Nextblockhash
	}
	g.TplName="GetBlock.html"
}
