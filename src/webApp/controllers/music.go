package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
)

type MusicController struct {
	beego.Controller
}

func (this *MusicController) Get() {
	fmt.Println("进入了music")
	// 配置管理
	this.TplName = "/music/index.html"
}
