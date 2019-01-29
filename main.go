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
	// maxIdle := 15
	// maxConn := 15
	// err := orm.RegisterDataBase("query", "mysql", "root:123456@tcp(127.0.0.1:3306)/youbon_querys?charset=utf8", maxIdle, maxConn)
	// if err != nil {
	// 	beego.Debug("query db:", err.Error())
	// }
	// beego.Debug("[UPDATE] register query database ok.")
}

func main() {

	fmt.Println("Lotter Version 0.12")

	orm.Debug = true
	orm.RunSyncdb("default", false, true)
	beego.BConfig.WebConfig.Session.SessionOn = true
	beego.BConfig.WebConfig.Session.SessionName = "youbon"
	beego.AddFuncMap("convertt", convertT)

	beego.Run()

}
