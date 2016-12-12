package routers

import (
	"github.com/astaxie/beego"
	"webApp/controllers"
)

func init() {
	beego.Router("/", &controllers.MainController{})
}
