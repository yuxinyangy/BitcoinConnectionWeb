package controllers

import (
	"BitcoinConnectionWeb/entity"
	"BitcoinConnectionWeb/getbitcion"
	"github.com/astaxie/beego"
)

type VaLiDateAddressController struct {
	beego.Controller
}

func (v *VaLiDateAddressController) Get()  {
	v.TplName="VaLiDateAddressParameter.html"
}

func (v *VaLiDateAddressController) Post()  {
	var address entity.DataAddress
	err := v.ParseForm(&address)
	if err != nil {
		v.TplName = "error.html"
		return
	}
	v.Data["DateAddress"]=getbitcion.VaLiDateAddress(address.Address)
	v.TplName="VaLiDateAddress.html"
}