package controllers

import (
	"BitcoinConnectionWeb/entity"
	"BitcoinConnectionWeb/getbitcion"
	"github.com/astaxie/beego"
)

type GetBlockHeaderController struct {
	beego.Controller
}

func (g *GetBlockHeaderController) Get()  {
	g.TplName="GetBlockHeaderParameter.html"
}

func (g *GetBlockHeaderController) Post()  {
	var hash entity.HASH
	err := g.ParseForm(&hash)
	if err != nil {
		g.TplName = "error.html"
		return
	}
	g.Data["BlockHeader"]=getbitcion.GetBlockHeader(hash.Hash)
	if getbitcion.GetBlockHeader(hash.Hash).Previousblockhash == "" {
		g.Data["Previousblockhash"]="这是创世区块，无上一区块的hash"
	}else{
		g.Data["Previousblockhash"]=getbitcion.GetBlockHeader(hash.Hash).Previousblockhash
	}
	if getbitcion.GetBlockHeader(hash.Hash).Nextblockhash == "" {
		g.Data["Nextblockhash"]="这是最高区块，无下一区块的hash"
	}else{
		g.Data["Nextblockhash"]=getbitcion.GetBlockHeader(hash.Hash).Nextblockhash
	}
	g.TplName="GetBlockHeader.html"
}
