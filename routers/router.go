package routers

import (
	"hello/controllers"
	"github.com/astaxie/beego"
)

func init() {
    //beego.Router("/", &controllers.MainController{})
	beego.Include(&controllers.MainController{})
}
