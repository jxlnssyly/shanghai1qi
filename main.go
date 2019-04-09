package main

import (
	_ "shanghai1qi/routers"
	"github.com/astaxie/beego"
	_ "shanghai1qi/models"
)

func main() {
	beego.Run()
}

