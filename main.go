package main

import (
	_ "cat/routers"
	"time"

	"github.com/astaxie/beego"
)

func main() {
	beego.SetStaticPath("/static", "public")
	beego.Run()
	time.Sleep(time.Second * 30)
}
