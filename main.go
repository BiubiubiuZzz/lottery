package main

import (
	_ "Lottery/routers"
	"fmt"
	"time"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

func convertT(in int64) (out string) {
	tm := time.Unix(in, 0)
	out = tm.Format("2006/01/02 15:04:05")
	return
}

func init() {

}

func main() {

	fmt.Println("Lotter Version 0.12")

	orm.Debug = true
	orm.RunSyncdb("default", false, false)
	beego.BConfig.WebConfig.Session.SessionOn = true
	beego.BConfig.WebConfig.Session.SessionName = "youbon"
	beego.AddFuncMap("convertt", convertT)

	beego.Run()

}
