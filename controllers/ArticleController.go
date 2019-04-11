package controllers

import (
	"github.com/astaxie/beego"
	"path"
	"time"
	"github.com/astaxie/beego/orm"
	"shanghai1qi/models"
	"math"
)

type ArticleController struct {
	beego.Controller
}

// 展示文章列表页
func (self *ArticleController) ShowArticleList() {

	// 获取数据
	o := orm.NewOrm()
	qs := o.QueryTable(&models.Article{})
	var articles []models.Article
	//_, err := qs.All(&articles)

	//if err != nil {
	//	beego.Info("查询数据错误")
	//
	//}

	// 查询总记录数
	count, _ := qs.Count()

	// 获取总页数
	pageSize := 2
	pageCount := math.Ceil(float64(count) / float64(pageSize))

	// 获取页码
	pageIndex, err := self.GetInt("pageIndex")
	//beego.Info(pageIndex)
	if err != nil {
		pageIndex = 1
	}

	qs.Limit(pageSize, (pageIndex-1)*pageSize).All(&articles)

	// 传递数据
	self.Data["articles"] = articles
	self.Data["pageIndex"] = pageIndex
	self.Data["count"] = count
	self.Data["pageCount"] = int(pageCount)

	self.TplName = "index.html"
}

// 展示添加文章页面
func (self *ArticleController) ShowAddArticle() {
	self.TplName = "add.html"
}

// 处理文章数据
func (self *ArticleController) HandleAddArticle() {
	// 1、获取数据
	articleName := self.GetString("articleName")
	content := self.GetString("content")

	// 2、校验数据
	if articleName == "" || content == "" {
		self.Data["errmsg"] = "添加数据不完整"
		self.TplName = "add.html"
		return
	}



	beego.Info(articleName, content)
	// 3、处理数据
	// 插入数据
	o := orm.NewOrm()

	var article models.Article
	article.ArticelName = articleName
	article.Acontent = content
	article.Aimg = Uploadfile(&self.Controller,"uploadname","add.html")

	o.Insert(&article)

	// 4、返回页面

	self.Redirect("/showArticleList", 302)
}

// 显示详情
func (self *ArticleController) ShowArticleDetail() {
	var model models.Article
	id, err := self.GetInt("articleId")

	if err != nil {
		beego.Info("传递的ID错误")
	}

	model.Id = id
	o := orm.NewOrm()
	o.Read(&model)

	model.Acount = model.Acount + 1

	o.Update(&model)
	self.Data["article"] = model

	self.TplName = "content.html"
}

func (self *ArticleController) ShowUpdateArticle() {
	var model models.Article
	id, err := self.GetInt("articleId")

	if err != nil {
		beego.Info("传递的ID错误")
	}
	model.Id = id
	o := orm.NewOrm()
	o.Read(&model)

	// 返回视图
	self.Data["article"] = model
	self.TplName = "update.html"
}

func (self *ArticleController) HandleUpdateArticle() {
	id, err := self.GetInt("articleId")
	aritcleName := self.GetString("articleName")
	content := self.GetString("content")

	filePath := Uploadfile(&self.Controller,"uploadname","update.html")

	if err != nil || aritcleName == "" || content == "" || filePath == "" {
		beego.Info("请求错误")
		beego.Info("err : ",err)
		beego.Info("aritcleName : ",aritcleName)
		beego.Info("content : ",content)
		beego.Info("filePath : ",filePath)
		return
	}

	if err != nil {
		beego.Info("传递的ID错误")
		return
	}

	var model models.Article
	model.Id = id
	o := orm.NewOrm()
	err = o.Read(&model)

	if err != nil {
		beego.Info("更新的文章不存在")
		return
	}

	model.ArticelName = aritcleName
	model.Acontent = content
	if filePath != "NoImg" {
		model.Aimg = filePath

	}
	o.Update(&model)

	self.Redirect("/showArticleList",302)

}

func (self *ArticleController) DeleteArticle() {
	id, err := self.GetInt("articleId")
	if err != nil  {
		beego.Info("请求错误")
		return
	}

	var model models.Article
	model.Id = id
	o := orm.NewOrm()
	err = o.Read(&model)

	o.Delete(&model)
	self.Redirect("/showArticleList", 302)
}

func (self *ArticleController) AddType() {
	o := orm.NewOrm()
	var types []models.ArticleType
	o.QueryTable("ArticleType").All(&types)
	self.Data["types"] = types
	self.TplName = "addType.html"
}

func (self *ArticleController) HandleAddType() {
	// 获取数据
	typeName := self.GetString("typeName")


	// 校验数据
	if typeName == "" {
		 beego.Info("信息不完整，请重新输入")
		 return
	}

	o := orm.NewOrm()
	var articleType models.ArticleType
	articleType.TypeName = typeName

	o.Insert(&articleType)

	// 返回数据
	self.Redirect("/addType", 302)

}

func (self *ArticleController) DelType() {
	typeId,err := self.GetInt("typeId")

	if err != nil {
		self.Data["errmsg"] = "ID不能为空"
		self.Redirect("/addType", 302)
		return
	}

	tp := models.ArticleType{Id:typeId}

	o := orm.NewOrm()
	_, err = o.Delete(&tp)

	if err != nil {
		self.Data["errmsg"] = "删除失败"
		self.Redirect("/addType", 302)
		return
	}

	self.Redirect("/addType", 302)

}



// 封装上传文件函数
func Uploadfile(self *beego.Controller,filePath string, tplName string) string {
	// 处理文件上传
	file, head, err := self.GetFile(filePath)

	if head.Filename == "" {
		return "NoImg"
	}

	if err != nil {
		self.Data["errmsg"] = "文件上传失败"
		self.TplName = tplName
		return ""
	}
	defer file.Close()

	// 1、文件大小
	if head.Size > 5000000 {
		self.Data["errmsg"] = "文件太大，请重新上传"
		self.TplName = tplName
		return ""
	}

	// 2、文件格式
	//
	ext := path.Ext(head.Filename)
	if ext != ".jpg" && ext != ".png" && ext != ".jpeg" {
		self.Data["errmsg"] = "文件格式错误，请重新上传"
		self.TplName = tplName
		return ""
	}

	// 3、防止重名
	fileName := time.Now().Format("2006-01-02-15:04:05") + ext

	self.SaveToFile(filePath, "./static/img/"+fileName)
	return "/static/img/" + fileName
}


