package controllers

import (
	"BitcoinConnectionWeb/getbitcion"
	"github.com/astaxie/beego"
)

type GetDifficultyController struct {
	beego.Controller
}

func (g *GetDifficultyController) Get()  {
	g.Data["Difficulty"]=getbitcion.GetDifficulty()
	g.TplName="GetDifficulty.html"
}
