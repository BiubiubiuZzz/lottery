package routers

import (
	"Lottery/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.IndexController{},"*:Index")
    beego.Router("/login",&controllers.AccountController{},"*:Login")
    beego.Router("/logout",&controllers.AccountController{},"*:Logout")

    beego.Router("/qrcode/query",&controllers.LotteryController{},"get:GetQrcode")
	beego.Router("/qrcode/query",&controllers.LotteryController{},"post:GetQrcode")
	beego.Router("/qrcode/queryById",&controllers.LotteryController{},"post:QueryById")

	beego.Router("/Winning/result",&controllers.LotteryController{},"get:GetWinning")
	beego.Router("/Winning/result",&controllers.LotteryController{},"post:GetWinning")
	beego.Router("/Winning/Query",&controllers.LotteryController{},"post:WinnQuery")

	beego.Router("/Prize/setting",&controllers.LotteryController{},"get:GetPrize")
	beego.Router("/Prize/setting",&controllers.LotteryController{},"post:GetPrize")
	beego.Router("/Prize/settingQuery",&controllers.LotteryController{},"post:SettingQuery")
	beego.Router("/Prize/setting/redpack",&controllers.LotteryController{},"get:GetRedPackPrize")
	beego.Router("/Prize/setting/redpack",&controllers.LotteryController{},"post:GetRedPackPrize")
	beego.Router("/Prize/settingRedPackQuery",&controllers.LotteryController{},"post:SettingRedPackQuery")


    beego.Router("/address/management",&controllers.LotteryController{},"get:GetAddress")
	beego.Router("/address/management",&controllers.LotteryController{},"post:GetAddress")
	beego.Router("/address/Query",&controllers.LotteryController{},"post:AddressQuery")

    beego.Router("/activity/setting",&controllers.LotteryController{},"*:Setting")
	beego.Router("/activity/SaveGift",&controllers.LotteryController{},"*:SaveGift")
	beego.Router("/activity/remove",&controllers.LotteryController{},"*:RemoveGift")
	beego.Router("/activity/editesetting",&controllers.LotteryController{},"*:EditeSetting")



    beego.Router("/address/setting",&controllers.LotteryController{},"*:SettingAddress")
    beego.Router("/address/save",&controllers.LotteryController{},"*:SaveAddress")
    beego.Router("/address/remove",&controllers.LotteryController{},"*:RemoveAddress")

    beego.Router("/rport/excel",&controllers.LotteryController{},"get:Getqr")

	beego.Router("/redpack/result",&controllers.LotteryController{},"get:GetRedPack")
    beego.Router("/redpack/result",&controllers.LotteryController{},"post:GetRedPack")
	beego.Router("/redpackquery",&controllers.LotteryController{},"*:RedPackQuery")

	beego.Router("/reset",&controllers.LotteryController{},"*:GetReset")

}
