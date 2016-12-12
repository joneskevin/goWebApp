package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	// "github.com/axgle/mahonia"
	// "html/template"
	"io"
	"io/ioutil"
	// "log"
	"net/http"
	"os"
	"strings"
	"webApp/models"
)

type IndexController struct {
	beego.Controller
}

// ServeHTTP
func (this *IndexController) Get() {
	// func (this *IndexController) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if ajaxRequest(this) {
		return
	}
	if flashUrlRequest(this) {
		return
	}
	if flashBarrage(this) {
		return
	}
	// if r.Method == "GET" {
	// 	t, _ := template.ParseFiles("index.html")
	// 	log.Println(t.Execute(w, nil))
	// }

	// r.URL.Path == "/"
	// r.ParseForm()
	// for k, v := range r.Form {
	// 	fmt.Println("key:", k)
	// 	fmt.Println("val:", strings.Join(v, ""))
	// }

	this.Data["Content_Info"] = "loading..."

	musics, err := models.GetAllMusic()
	if err != nil {
		beego.Error(err)
	}
	//首页上音乐4个就够了
	list := musics[:4]
	this.Data["Musics"] = list

	//6个照片
	this.Data["Photos"] = [6]string{"01", "02", "03", "04", "05", "06"}

	//视频地址	//注意这里暂时没用 放到了app.conf里了
	this.Data["Cid"] = "9059390"
	this.Data["Aid"] = "5577814"

	// 配置管理
	this.TplName = "index.html"
}

//这里设置的是ajax的页面请求
func ajaxRequest(this *IndexController) bool {
	page := this.Input().Get("page")
	switch page {
	case "1":
		io.WriteString(this.Ctx.ResponseWriter, "page_1")
		return true
	case "2":
		io.WriteString(this.Ctx.ResponseWriter, "page_2")
		return true
	case "3":
		io.WriteString(this.Ctx.ResponseWriter, "page_3")
		return true
	case "4":
		io.WriteString(this.Ctx.ResponseWriter, beego.AppConfig.String("bofangqi"))
		return true
	case "5":
		io.WriteString(this.Ctx.ResponseWriter, beego.AppConfig.String("game"))
		return true
	}
	return false
}

//这里请求的是flash的playurl数据
func flashUrlRequest(this *IndexController) bool {
	playurl := this.Input().Get("playurl")
	if len(playurl) > 0 {

		url := "http://interface.bilibili.com/playurl?"

		r := this.Ctx.Request
		r.ParseForm()

		for k, v := range r.Form {
			if k == "playurl" {
				fmt.Println("我明明匹配了啊")
				continue
			}
			url += "&" + k + "="
			url += strings.Join(v, "")
		}
		//url=http://interface.bilibili.com/playurl?&cid=9059390&player=1&ts=1477477275&sign=75dbf022be036a1bc34ac35dba591734
		fmt.Println("flash请求播放数据  转发数据网址为：" + url)

		resp, e := http.Get(url)
		if e != nil {
			beego.Error(e)
		}
		defer resp.Body.Close()
		body, er := ioutil.ReadAll(resp.Body)
		if er != nil {
			beego.Error(er)
		}
		// fmt.Println("数据：" + string(body))
		io.WriteString(this.Ctx.ResponseWriter, string(body))
		return true
	}
	return false
}

//这里请求的是flash的弹幕数据
func flashBarrage(this *IndexController) bool {
	barrage := this.Input().Get("barrage")
	if len(barrage) > 0 {
		cid := this.Input().Get("cid")

		url := "http://comment.bilibili.com/" + cid + ".xml"

		fmt.Println("flash请求弹幕数据  转发数据网址为：" + url)

		fi, err := os.Open("data/danmu.xml")
		if err != nil {
			beego.Error(err)
		}
		defer fi.Close()
		body, err := ioutil.ReadAll(fi)

		//这里呢，有一个问题，就是我通过网上请求xml 读取以后 是乱码，尝试使用了Encoder和Decoder各种编解码也没解决，发现直接读取本地xml文件的话就没有乱码问题，所以先读本地的吧，有时间在研究从网上请求xml的问题
		// resp, e := http.Get(url)
		// if e != nil {
		// beego.Error(e)
		// }
		// defer resp.Body.Close()
		// body, er := ioutil.ReadAll(resp.Body)
		// if er != nil {
		// beego.Error(er)
		// }

		// dec := mahonia.NewDecoder("GBK")
		// xmls := dec.ConvertString(string(body))
		// fmt.Println("弹幕：" + string(body))
		io.WriteString(this.Ctx.ResponseWriter, string(body))
		return true
	}
	return false
}
