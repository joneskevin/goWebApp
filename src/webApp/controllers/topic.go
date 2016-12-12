package controllers

import (
	"github.com/astaxie/beego"
	// "path"
	"strings"
	"webApp/models"
)

type TopicController struct {
	beego.Controller
}

func (this *TopicController) Get() {
	this.Data["IsTopic"] = true
	this.TplName = "topic.html"
	this.Data["IsLogin"] = checkAccount(this.Ctx)

	topics, err := models.GetAllTopics("", "", false)
	if err != nil {
		beego.Error(err)
	}
	this.Data["Topics"] = topics
}

func (this *TopicController) Add() {
	if !checkAccount(this.Ctx) {
		this.Redirect("/login", 302)
		return
	}

	this.TplName = "topic_add.html"
	this.Data["IsLogin"] = true
}

func (this *TopicController) Post() {
	if !checkAccount(this.Ctx) {
		this.Redirect("/login", 302)
		return
	}

	// 解析表单
	tid := this.Input().Get("tid")
	title := this.Input().Get("title")
	content := this.Input().Get("content")
	category := this.Input().Get("category")
	lable := this.Input().Get("lable")

	// 获取附件
	// _, fh, err := this.GetFile("attachment")
	// if err != nil {
	// 	beego.Error(err)
	// }

	// var attachment string
	// if fh != nil {
	// 	// 保存附件
	// 	attachment = fh.Filename
	// 	beego.Info(attachment)

	// 	err = this.SaveToFile("attachment", path.Join("attachment", attachment))
	// 	if err != nil {
	// 		beego.Error(err)
	// 	}
	// }
	var err error

	if len(tid) == 0 {
		err = models.AddTopic(title, category, lable, content, "")
		// err = models.AddTopic(title, category, lable, content, attachment)
	} else {
		//err = models.ModifyTopic(tid, title, category, lable, content, attachment)
	}

	if err != nil {
		beego.Error(err)
	}

	this.Redirect("/topic", 302)
}
func (this *TopicController) View() {
	this.TplName = "topic_view.html"

	reqUrl := this.Ctx.Request.RequestURI
	i := strings.LastIndex(reqUrl, "/")
	tid := reqUrl[i+1:]

	//获取文章
	topic, err := models.GetTopic(tid)
	if err != nil {
		beego.Error(err)
		this.Redirect("/", 302)
		return
	}
	this.Data["Topic"] = topic
	if len(topic.Attachment) > 0 {
		this.Data["isAttachment"] = true
	}
	this.Data["Lables"] = strings.Split(topic.Lables, " ")

	// 获取留言
	// replies, err := models.GetAllReplies(tid)
	// if err != nil {
	// 	beego.Error(err)
	// 	return
	// }

	// this.Data["Replies"] = replies
	this.Data["IsLogin"] = checkAccount(this.Ctx)
}
