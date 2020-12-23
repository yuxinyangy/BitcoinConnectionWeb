package routers

import (
	"BitcoinConnectionWeb/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
    beego.Router("/register.html",&controllers.RegisterController{})
    beego.Router("/user_register",&controllers.RegisterController{})
	beego.Router("/user_login", &controllers.LoginController{})
}
