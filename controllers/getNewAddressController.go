package controllers

import (
	"BitcoinConnectionWeb/entity"
	"BitcoinConnectionWeb/getbitcion"
	"github.com/astaxie/beego"
)


type GetNewAddressController struct {
	beego.Controller
}

func (g *GetNewAddressController)Get(){
	g.TplName="GetNewAddressParameter.html"
}
func (g *GetNewAddressController)Post()  {
	var address entity.NewAddress
	err := g.ParseForm(&address)
	if err != nil {
		g.TplName = "error.html"
		return
	}
	g.Data["Type"]=address.Type
	g.Data["NewAddress"]=getbitcion.GetNewAddress(address.Label,address.Type)
	g.TplName="GetNewAddress.html"
}
