package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"shanghai1qi/models"
)

type UserController struct {

	beego.Controller
}

// 显示注册页面
func (self *UserController)ShowRegister()  {
	self.TplName = "register.html"
}

// 处理注册数据
func (self *UserController) HandleResister() {
	// 1、获取数据
	userName := self.GetString("userName")
	pwd := self.GetString("password")
	// 2、校验数据

	if userName == "" || pwd == "" {
		self.Data["errmsg"] = "注册数据不完整，请重新注册"
		 beego.Info("注册数据不完整，请重新注册")
		 self.TplName = "register.html"
		 return
	}

	// 3、操作数据
	o := orm.NewOrm()

	// 获取插入对象
	var user models.User
	user.Name = userName
	user.Password = pwd

	// 给插入对象赋值
	o.Insert(&user)


	// 4、返回页面
	//self.Ctx.WriteString("注册成功")
	self.Redirect("/login",302)
}


func (self *UserController) ShowLogin() {
	self.TplName = "login.html"
}

func (self *UserController)HandleLogin()  {
	userName := self.GetString("userName")
	pwd := self.GetString("password")

	if userName == "" || pwd == "" {
		self.Data["errmsg"] = "登录数据不完整"
		self.TplName = "login.html"
		return
	}

	// 操作数据
	o := orm.NewOrm()
	var user models.User
	user.Name = userName
	user.Password = pwd
	err := o.Read(&user,"Name")
	if err != nil {
		self.Data["errmsg"] = "用户不存在"
		self.TplName = "login.html"
		return
	}

	if user.Password != pwd {
		self.Data["errmsg"] = "密码错误"
		self.TplName = "login.html"
		return
	}

	// 返回页面
	//self.Ctx.WriteString("登录成功")
	self.Redirect("/showArticleList",302)

}