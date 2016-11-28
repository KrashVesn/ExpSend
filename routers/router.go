package routers

import (
	"ExpSend/controllers"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/", &controllers.MainController{}, "*:Post")
}
