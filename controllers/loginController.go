package controllers

import (
	"BitcoinConnectionWeb/models"
	"fmt"
	"github.com/astaxie/beego"
)

type LoginController struct {
	beego.Controller
}

func (l *LoginController) Get() {
	l.TplName = "register.html"
}

func (l *LoginController) Post() {
	var user models.User
	err := l.ParseForm(&user)
	fmt.Println(user)
	if err != nil {

		l.TplName = "error.html"
		return
	}
	_,err = user.QueryUser()
	if err!=nil{
		fmt.Println(err.Error())
		l.TplName = "error.html"
		return
	}
	l.TplName = "home.html"

}
