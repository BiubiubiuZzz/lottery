package models

import (
	"crypto/md5"
	"fmt"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

func init() {
	RegisterDB()
}
func RegisterDB() {
	maxIdle := 15
	maxConn := 15
	orm.RegisterModel(new(LotteryUser), new(LuckybagLottoryGifts),new(LuckybagLottoryAddress))
	orm.RegisterDriver("mysql", orm.DRMySQL)
	err := orm.RegisterDataBase("default", "mysql", "root:123456@tcp(127.0.0.1:3306)/youbon?charset=utf8", maxIdle, maxConn)
	if err != nil {
		beego.Debug("default db:", err.Error())
	}
	err = orm.RegisterDataBase("update", "mysql", "root:123456@tcp(127.0.0.1:3306)/youbon_update?charset=utf8", maxIdle, maxConn)
	if err != nil {
		beego.Debug("update db:", err.Error())
	}
	beego.Debug("[BOSS] register update database ok.")

}

func Md5(buf []byte) string {
	mymd5 := md5.New()
	mymd5.Write(buf)
	result := mymd5.Sum(nil)
	return fmt.Sprintf("%x", result)
}
