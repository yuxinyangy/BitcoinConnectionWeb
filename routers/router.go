package routers

import (
	"BitcoinConnectionWeb/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
    //
    beego.Router("/register.html",&controllers.LoginController{})
    //
    beego.Router("/user_register",&controllers.RegisterController{})
}
