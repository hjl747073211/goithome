package routers

import (
	"gozh/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
    beego.Router("/ask", &controllers.AskController{},"*:GetAsk")
}
