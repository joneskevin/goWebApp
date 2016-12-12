package controllers

import (
	"github.com/astaxie/beego"
	"webApp/models"
)

//主页
type HomeController struct {
	beego.Controller
}

func (this *HomeController) Get() {
	this.Data["IsHome"] = true
	// 配置管理
	this.TplName = "home.html"
	this.Data["IsLogin"] = checkAccount(this.Ctx)

	topics, err := models.GetAllTopics(
		this.Input().Get("cate"), this.Input().Get("lable"), true)
	if err != nil {
		beego.Error(err)
	}
	this.Data["Topics"] = topics

	// categories, err := models.GetAllCategories()
	// if err != nil {
	// 	beego.Error(err)
	// }
	// this.Data["Categories"] = categories

}
