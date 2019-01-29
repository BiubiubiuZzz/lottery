package controllers

import (
	"strings"
	"Lottery/models"
	//"strconv"
)

type AccountController struct {
	baseController
}

func (this *AccountController) Login() {
	if this.GetString("dosubmit") == "yes" {
		account := strings.TrimSpace(this.GetString("account"))
		password := strings.TrimSpace(this.GetString("password"))
		if account != "" && password != "" {
			var user = &models.LotteryUser{}
			user.Username = account
			if user.Read("username") != nil || user.Password != models.Md5([]byte(password)) {
				this.Data["errmsg"] = "账号或密码错误!"
			} else {
				this.SetSession("login", true)
				this.SetSession("uid", user.Id)
				this.SetSession("username", user.Username)
				this.SetSession("password", user.Password)

				this.Redirect("/", 302)
			}
		}
	}
	this.TplName = "account_login.html"
}


func (this *AccountController) Logout() {
	this.DelSession("login")
	this.DelSession("uid")
	this.DelSession("username")
	this.DelSession("password")
	this.Redirect("/login", 302)
}

