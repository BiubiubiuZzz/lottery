package controllers


type IndexController struct {
	baseController
}

func (this *IndexController) Index() {
	Logined := this.GetSession("login")
	if Logined != true {
		this.Redirect("/login", 302)
		return
	}
	this.display()
}