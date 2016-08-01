package routers

import (
	"github.com/gagandeepsingh485/learn-go/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
}
