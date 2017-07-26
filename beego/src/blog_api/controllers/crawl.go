package controllers

import (


	"github.com/astaxie/beego"
	"blog_api/models"

	"github.com/astaxie/beego/logs"
)

// Operations about object
type CrawlController struct {
	beego.Controller
}




// @Title Get
// @Description get input
// @Param	input		query 	string	true		"the input you want to get"
// @Success 200 {string} models.Crawl
// @Failure 403 :string is empty
// @router / [get]
func (o *CrawlController) Get() {
	input := o.Ctx.Input.Query("input")
	logs.SetLogger(logs.AdapterFile,`{"filename":"crawl.log","level":7,"maxlines":0,"maxsize":0,"daily":true,"maxdays":10}`)
	logs.Info(input)
	if input != "" {
		ob:= models.Crawl(input)
		o.Data["json"] = ob

	}
	o.ServeJSON()
}
