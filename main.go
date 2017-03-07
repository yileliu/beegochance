package main

import (
	_ "beegochance/routers"

	"github.com/astaxie/beego"
)

func main() {
	beego.SetStaticPath("/public", "static")
	beego.DelStaticPath("/static")
	beego.Run()
}
