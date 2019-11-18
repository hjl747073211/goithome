package main

import (
	_ "gozh/routers"
	_ "gozh/models"
	"github.com/astaxie/beego"
)

func main() {
	beego.Run()
}

