package routers

import (
	"cat/controllers"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/cat/:id", &controllers.SearchByIdController{})
}
