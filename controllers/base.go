package controllers

import (
	"github.com/astaxie/beego"
	"Lottery/models"

)

type baseController struct {
	beego.Controller
	pager *models.Pager

}


func (this *baseController) Prepare() {

	page, err := this.GetInt("page")
	if err != nil {
		page = 1
	}
	pagesize := 50
	this.pager = models.NewPager(page, pagesize, 0, "")
}

func (this *baseController) display(tplname ...string) {
	this.Data["queryValue"] = ""
	this.TplName = "QRcode_query.html"

}

func (this *baseController) showmsg(msg ...string) {
	this.display("showmsg")
	if len(msg) == 1 {
		msg = append(msg, this.Ctx.Request.Referer())
	}
	this.Data["msg"] = msg[0]
	this.Data["redirect"] = msg[1]
	this.Render()
	this.StopRun()
}




