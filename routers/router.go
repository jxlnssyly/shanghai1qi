package routers

import (
	"shanghai1qi/controllers"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
)

func init() {

	beego.InsertFilter("/article/*",beego.BeforeExec,Filter)

    beego.Router("/", &controllers.ArticleController{},"get:ShowArticleList")

    beego.Router("/register", &controllers.UserController{},"get:ShowRegister;post:HandleResister")
	beego.Router("/login", &controllers.UserController{},"get:ShowLogin;post:HandleLogin")

    // 文章列表页访问
    beego.Router("/article/showArticleList",&controllers.ArticleController{},"get:ShowArticleList")

    // 添加文章
    beego.Router("/article/addArticle",&controllers.ArticleController{}, "get:ShowAddArticle;post:HandleAddArticle")

    beego.Router("/article/showArticleDetail",&controllers.ArticleController{},"get:ShowArticleDetail")

	beego.Router("/article/updateArticle",&controllers.ArticleController{},"get:ShowUpdateArticle;post:HandleUpdateArticle")

    beego.Router("/article/deleteArticle",&controllers.ArticleController{},"get:DeleteArticle")

    beego.Router("/article/addType",&controllers.ArticleController{},"get:AddType;post:HandleAddType")
	beego.Router("/article/delType",&controllers.ArticleController{},"get:DelType")
    beego.Router("/article/logout",&controllers.UserController{},"get:Logout")
    // 给请求指定自定义方法
    //beego.Router("/login",&controllers.LoginController{},"get:ShowLogin;post:PostFunc")
	//
    //// 给多个请求指定方法
    //beego.Router("/index",&controllers.IndexController{},"get,post:HandleFunc")
	//
    //// 给所有请求指定一个方法
	//beego.Router("/index",&controllers.IndexController{},"*:HandleFunc")
	//
    //// 当两种指定方法冲突
	//beego.Router("/index",&controllers.IndexController{},"*:HandleFunc;post:PostFunc")
}

var Filter = func(ctx *context.Context) {
	userName := ctx.Input.Session("userName")

	if userName == "" {
		 ctx.Redirect(302,"/login")
		 return
	}

}