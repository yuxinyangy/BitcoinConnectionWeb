package controllers

import (
	"BitcoinConnectionWeb/models"
	"github.com/astaxie/beego"
)

type RegisterController struct {
	beego.Controller
}

func (r *RegisterController) Post() {
	//解析请求数据
	var user models.User
	err:=r.ParseForm(&user)
	if err!=nil{
		r.TplName = "error.html"
		return
	}
	//保存用户信息
	_,err = user.SaveUser()
	//返回前端结果
	if err!=nil{
		r.TplName = "error.html"
		return
	}
	//用户注册成功
	r.TplName = "login.html"
}
