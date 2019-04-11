package routers

import (
	"shanghai1qi/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})

    beego.Router("/register", &controllers.UserController{},"get:ShowRegister;post:HandleResister")
	beego.Router("/login", &controllers.UserController{},"get:ShowLogin;post:HandleLogin")

    // 文章列表页访问
    beego.Router("/showArticleList",&controllers.ArticleController{},"get:ShowArticleList")

    // 添加文章
    beego.Router("/addArticle",&controllers.ArticleController{}, "get:ShowAddArticle;post:HandleAddArticle")

    beego.Router("/showArticleDetail",&controllers.ArticleController{},"get:ShowArticleDetail")

	beego.Router("/updateArticle",&controllers.ArticleController{},"get:ShowUpdateArticle;post:HandleUpdateArticle")
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
