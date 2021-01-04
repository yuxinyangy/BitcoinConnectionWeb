package controllers

import (
	"BitcoinConnectionWeb/models"
	"fmt"
	"github.com/astaxie/beego"
)

type LoginController struct {
	beego.Controller
}

func (l *LoginController) Get()  {
	l.TplName="login.html"
}

func (l *LoginController) Post() {
	var user models.User
	err := l.ParseForm(&user)
	if err != nil {
		l.TplName = "error.html"
		return
	}
	//查询数据库的用户信息
	_, err = user.QueryUser()
	if err != nil {
		fmt.Println(err.Error())
		l.TplName = "error.html"
		return
	}
	//登入成功,跳转项目核心功能页面(home.html)
	l.TplName = "home.html"

}
