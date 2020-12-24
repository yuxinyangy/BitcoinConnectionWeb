package routers

import (
	"BitcoinConnectionWeb/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
    //跳转注册页面
    beego.Router("/register.html",&controllers.LoginController{})
    //把用户信息提交到数据库中
    beego.Router("/user_register",&controllers.RegisterController{})
    //跳转主页面
    beego.Router("/user_home",&controllers.LoginController{})
}
