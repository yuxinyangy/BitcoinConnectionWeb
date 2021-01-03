package controllers

import (
	"BitcoinConnectionWeb/getbitcion"
	"github.com/astaxie/beego"
)

type VerifyChainController struct {
	beego.Controller
}

func (v *VerifyChainController) Get()  {
	v.Data["VerifyChain"]=getbitcion.VerifyChain()
	v.TplName="VerifyChain.html"
}
