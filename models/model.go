package models

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/astaxie/beego/orm"
	"time"
)

type User struct {
	Id int
	Name string
	Password string
	Articles []*Article `orm:"reverse(many)"`
}

type Article struct {
	Id int `orm:"pk;auto"`
	ArticelName string `orm:"size(20)"`
	Atime time.Time `orm:"auto_now"`
	Acount int `orm:"default(0);null"`
	Acontent string `orm:"size(500)"`
	Aimg string `orm:"size(100)"`

	ArticleType *ArticleType `orm:"rel(fk)"`
	Users [] *User `orm:"rel(m2m)"`
}

// 类型表
type ArticleType struct {
	Id int
	TypeName string `orm:"size(20)"`

	Articles [] *Article `orm:"reverse(many)"`
}



func init()  {
	 // 操作数据库

	 // 1、获取连接对象
	 orm.RegisterDataBase("default","mysql","root:123456@tcp(127.0.0.1:3306)/test?charset=utf8")

	 // 2、创建表
	orm.RegisterModel(new(User), new(Article),new(ArticleType))
	 orm.RunSyncdb("default",false,true)
	 // 3、操作表


}