package main

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"webApp/controllers"
	"webApp/models"
	//_ "webApp/routers"
	// "log"
	// "net/http"
	// "os"
)

func init() {
	//注册数据库
	models.RegisterDB()

}

func main() {
	// 开启 ORM 调试模式
	orm.Debug = false
	//自动建表
	// orm.RunSyncdb("default", false, true)

	// mux := http.NewServeMux()

	// mux.Handle("/", &controllers.IndexController{})

	// 使用函数作为 handler
	// mux.HandleFunc("/bye", sayBye)

	// 静态文件
	// wd, err := os.Getwd()
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// mux.Handle("/static/", http.StripPrefix("/static/",
	// 	http.FileServer(http.Dir(wd))))

	// err = http.ListenAndServe(":8080", mux)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	//注册 beego 路由
	beego.Router("/", &controllers.IndexController{})
	// beego.Router("/music", &controllers.MusicController{})
	beego.Router("/home", &controllers.HomeController{})
	beego.Router("/login", &controllers.LoginController{})
	beego.Router("/topic", &controllers.TopicController{})
	beego.AutoRouter(&controllers.TopicController{}) //自动路由

	// 启动 beego
	beego.Run()
}
