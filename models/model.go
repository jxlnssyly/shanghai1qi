package models

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/astaxie/beego/orm"
	)

type User struct {
	Id int
	Name string
	Password string
}

func init()  {
	 // 操作数据库

	 // 1、获取连接对象
	 orm.RegisterDataBase("default","mysql","root:123456@tcp(127.0.0.1:3306)/test?charset=utf8")

	 // 2、创建表
	orm.RegisterModel(new(User))
	 orm.RunSyncdb("default",false,true)
	 // 3、操作表


}