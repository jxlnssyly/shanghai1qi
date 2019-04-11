package main

import (
	_ "shanghai1qi/routers"
	"github.com/astaxie/beego"
	_ "shanghai1qi/models"
)

func main() {
	beego.AddFuncMap("prepage",ShowPrePage)
	beego.AddFuncMap("nextpage",ShowNextPage)
	beego.Run()
}

// 后台定义一个函数
func ShowPrePage(page int) int {
	if page <= 1 {
		return 1
	}
	return page - 1
}

func ShowNextPage(page int,pageCount int) int {
	if page >= pageCount {
		return pageCount
	}

	return page + 1
}
