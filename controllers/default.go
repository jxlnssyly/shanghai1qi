package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"shanghai1qi/models"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.Data["data"] = "china"

	c.TplName = "test.html"

}

func (c *MainController)Post()  {

	// 1、获取ORM对象
	o := orm.NewOrm()

	// 2、执行某个操作函数, 增删改查
	var user models.User
	// 插入
/*
	user.Name = "Liaoy"
	user.Password = "123456"


	// 3.返回结果
	count, err := o.Insert(&user)

	if err != nil {
		beego.Error(err)
		return
	}
	beego.Info(count)
*/

	//查询
	/*
	user.Id = 2
	err := o.Read(&user,"Id")
	if err != nil {
		beego.Error("查询失败:",err)
	}
	beego.Info(user)
*/

	// 修改
	/*
	user.Id = 2
	err := o.Read(&user)

	if err != nil {
		beego.Error("要更新的数据不存在",err)
	}

	user.Name = "廖延"
	count, err := o.Update(&user)

	if err != nil {
		beego.Error("更新失败")
	}
	beego.Info(count)
	*/

	// 删除
	user.Id = 3
	count, err := o.Delete(&user)

	if err != nil {
		beego.Error("删除失败",err)
	}
	beego.Info(count)

	c.Data["data"] = "POST"
	c.TplName = "test.html"
}
