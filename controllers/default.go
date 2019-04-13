package controllers

import (
	"github.com/astaxie/beego"
	"github.com/gomodule/redigo/redis"
	"github.com/astaxie/beego/orm"
	"shanghai1qi/models"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {


	// 连接Redis数据库
	conn, err := redis.Dial("tcp",":6379")
	if err!=nil {
		c.Ctx.WriteString("连接Redis错误")
		return
	}
	// 关闭数据库
	defer conn.Close()

	// 第一种写法
	/*
	// 执行操作数据库语句
	conn.Send("set","liaoy","真棒")

	// 返回结果
	conn.Flush()

	rep, _ := conn.Receive()

	*/

	// 第二种写法
	rep, err := conn.Do("get","liaoy")

	str, err :=redis.String(rep,err)

	if err != nil {
		c.Ctx.WriteString("获取内容出错")
		return
	}

	beego.Info(rep)
	c.Ctx.WriteString("获得：" + str)
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
