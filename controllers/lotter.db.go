package controllers

import (
	"Lottery/models"
	"fmt"
	"strconv"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)



// func GetAddress() []*models.LuckybagLottoryAddress {
// 	var address []*models.LuckybagLottoryAddress
// 	o := orm.NewOrm()
// 	o.Using("update")
// 	_, err := o.Raw("SELECT distinct(logs.gift_name),luck.id,luck.phone,luck.name,luck.address,luck.email,luck.date FROM " +
// 		" luckybag_lottory_gifts_logs as logs left JOIN " +
// 		" luckybag_lottory_address as luck on luck.open_id=logs.open_id").QueryRows(&address)
// 	if err != nil {
// 		beego.Debug("[ADMIN REPORT] GET a address manager error:", err.Error())
// 		return nil
// 	}
// 	beego.Debug("[ADMIN REPORT] get a AddressManager：", len(address))
// 	return address
// }

//查询地址
func GetAddress() []*models.LuckybagLottoryGiftsLogs {
	var giftlogs []*models.LuckybagLottoryGiftsLogs
	o := orm.NewOrm()
	o.Using("update")
	_, err := o.Raw("SELECT distinct(gift_name),open_id FROM luckybag_lottory_gifts_logs").QueryRows(&giftlogs)
	if err != nil {
		beego.Debug("[ADMIN REPORT] GET a address manager error:", err.Error())
		return nil
	}
	beego.Debug("[ADMIN REPORT] get a AddressManager：", len(giftlogs))
	for i := 0; i <len(giftlogs); i++ {
		o := orm.NewOrm()
		o.Using("update")
		var address *models.LuckybagLottoryAddress
		logs := giftlogs[i]

		err1 := o.Raw("SELECT * from luckybag_lottory_address where open_id =?",logs.OpenId).QueryRow(&address)
		if err1 != nil{
			beego.Debug("[ADMIN REPORT] get error:",err1)
			return nil
		}
		logs.OpenId = address.OpenId
		logs.Address = address.Address
		logs.Name =address.Name
		logs.Phone =address.Phone
		logs.Email =address.Email
		logs.AddressDate = address.Date
	}

	return giftlogs
}

//***注：所有QR表示抽奖码；
//全部抽奖码显示
func GetQR()[]*models.LuckybagLottory  {
	var QR []*models.LuckybagLottory
	o := orm.NewOrm()
	o.Using("update")
	_,err := o.Raw("select *  from luckybag_lottory").QueryRows(&QR)
	if err != nil{
		beego.Debug("[ADMIN REPORT] get a error:",err.Error())
		return nil
	}
	beego.Debug("[ADMIN REPORT] get winning:",len(QR))
	return QR
}

//Qr通过Id 查询
func GetQRcode(id string) []*models.LuckybagLottory {
	var QR []*models.LuckybagLottory
	o := orm.NewOrm()
	o.Using("update")
	search := fmt.Sprintf("select * from luckybag_lottory where id REGEXP '%s'", id)
	_, err := o.Raw(search).QueryRows(&QR)
	if err != nil {
		beego.Debug("[ADMIN REPORT]get a QR-code error:", err.Error(), "Id:", id)
		return nil
	}
	beego.Debug("[ADMIN REPORT] get a QR-code ：", len(QR), "Id:", id)
	if QR == nil {
		search := fmt.Sprintf("select * from luckybag_lottory where qx = '%s'", id)
		_, err = o.Raw(search).QueryRows(&QR)
	}
	return QR
}

//查询use使用总数
func GetGiftUsedByGiftID(giftID int64) int64 {
	var result int64 = 0
	o := orm.NewOrm()
	o.Using("update")
	cond := fmt.Sprintf("select count(*) as used from luckybag_lottory_gifts_logs where gift_id=%d ", giftID)
	o.Raw(cond).QueryRow(&result)
	return result
}

//显示剩余数量
func GetLeftQuantity(giftID int64) int64 {
	var totalUsedCount int64
	var gift models.LuckybagLottoryGifts
	o := orm.NewOrm()
	o.Using("update")
	sql := fmt.Sprintf("SELECT * FROM luckybag_lottory_gifts where id = %d", giftID)
	err := o.Raw(sql).QueryRow(&gift)
	if err != nil {
		beego.Debug("[ADMIN REPORT]GET a quantity err:", err.Error())
	}

	sql1 := fmt.Sprintf("SELECT count(*) FROM luckybag_lottory_gifts_logs where gift_id = %d and date >= %d", giftID, gift.Date)
	err1 := o.Raw(sql1).QueryRow(&totalUsedCount)
	if err1 != nil {
		beego.Debug("[ADMIN REPORT] get a use number err:", err1.Error())
	}

	leftQuantity := gift.Quantity - totalUsedCount
	if leftQuantity < 0 {
		leftQuantity = 0
	}

	return leftQuantity
}

//显示活动设置的数据
func GetActivity() []*models.LuckybagLottoryGifts {
	var AC []*models.LuckybagLottoryGifts
	var result []*models.LuckybagLottoryGifts = nil
	o := orm.NewOrm()
	o.Using("update")
	_, err := o.Raw("select * from luckybag_lottory_gifts").QueryRows(&AC)
	if err != nil {
		beego.Debug("[ADMIN REPORT] GET a error:", err.Error())
		return nil
	}
	beego.Debug("[ADMIN REPORT] get activity:", len(AC))

	for _, gift := range AC {
		gift.Used = GetGiftUsedByGiftID(gift.Id)
		gift.LeftQuantity = GetLeftQuantity(gift.Id)
		result = append(result, gift)
	}
	return result

}

//添加活动奖品查询
func GetActivityByName(awardName string) []*models.LuckybagLottoryGifts {
	var AC []*models.LuckybagLottoryGifts
	var result []*models.LuckybagLottoryGifts = nil
	o := orm.NewOrm()
	o.Using("update")

	cond := fmt.Sprintf("select * from luckybag_lottory_gifts where gift_name REGEXP '%s'", awardName)
	_, err := o.Raw(cond).QueryRows(&AC)
	if err != nil {
		beego.Debug("[ADMIN REPORT] GET a error:", err.Error())
		return nil
	}
	beego.Debug("[ADMIN REPORT] get activity:", len(AC))
	for _, gift := range AC {
		gift.Used = GetGiftUsedByGiftID(gift.Id)
		gift.LeftQuantity = GetLeftQuantity(gift.Id)
		result = append(result, gift)

	}
	return result
}

//中奖商品Id查询
func GetWinningByCodeId(id string) []*models.LuckybagLottory {
	var wi []*models.LuckybagLottory
	o := orm.NewOrm()
	o.Using("update")
	_,err := o.Raw("SELECT * FROM luckybag_lottory where id=?",id).QueryRows(&wi)
	if err != nil{
		beego.Debug("[ADMIN REPORT] get a error:",err.Error())
		return nil
	}
	beego.Debug("[ADMIN REPORT]get a luckybag_lottery_gifts_logs logs:",len(wi))
	for i := 0; i <len(wi);i++{
		o := orm.NewOrm()
		o.Using("update")
		var lottory *models.LuckybagLottoryGiftsLogs
		logs := wi[i]

		err1 := o.Raw("select code,gift_name,date from luckybag_lottory_gifts_logs where code=? ",logs.Qx).QueryRow(&lottory)
		if err1 != nil{
			beego.Debug("[ADMIN REPORT]get error1:",err1)
			return nil
		}
		logs.GiftName =lottory.GiftName
		logs.Date = lottory.Date
	}
	return wi
}

//地址中奖活动查询
func GetAddressQuser(giftname string) []*models.LuckybagLottoryAddress {
	var AQ []*models.LuckybagLottoryAddress
	o := orm.NewOrm()
	o.Using("update")
	cond := fmt.Sprintf("SELECT distinct(logs.gift_name),luck.phone,luck.name,luck.address,luck.email,luck.date FROM "+
		" luckybag_lottory_gifts_logs as logs left JOIN "+
		" luckybag_lottory_address as luck on luck.open_id=logs.open_id WHERE logs.gift_name REGEXP '%s'", giftname)
	_, err := o.Raw(cond).QueryRows(&AQ)
	if err != nil {
		beego.Debug("[ADMIN REPORT] get a error:", err.Error())
		return nil
	}
	beego.Debug("[ADMIN REPORT] get addressquser:", len(AQ))
	return AQ

}

//根据deliver ID 获取所有的gift 信息，返回数组
func GetLotteryGiftByDeliverID(id int) ([]models.LuckybagLottoryGifts, error) {
	var gift []models.LuckybagLottoryGifts
	o := orm.NewOrm()
	o.Using("update")

	criter := fmt.Sprintf("select * from luckybag_lottory_gifts where deliver_id=%d order by date ASC", id)
	_, err := o.Raw(criter).QueryRows(&gift)

	if err != nil {
		beego.Debug("[ADMIN REPORT] GET a error:", err.Error())
		return nil, err
	}
	return gift, err

}

//活动编辑根据id
func GetLotteryGiftByID(id int) (*models.LuckybagLottoryGifts, error) {
	var gift models.LuckybagLottoryGifts
	o := orm.NewOrm()
	o.Using("update")

	criter := fmt.Sprintf("select * from luckybag_lottory_gifts where id=%d", id)
	err := o.Raw(criter).QueryRow(&gift)

	if err != nil {
		beego.Debug("[ADMIN REPORT] GET a error:", err.Error())
		return nil, err
	}
	return &gift, err

}

//删除活动
func RemoveLotteryGiftByID(id int64) error {
	o := orm.NewOrm()
	o.Using("update")
	criter := fmt.Sprintf("delete from luckybag_lottory_gifts where id=%d", id)
	_, err := o.Raw(criter).Exec()

	if err != nil {
		beego.Debug("[ADMIN REPORT] GET a error:", err.Error())
		return err
	}
	return err
}

//更新活动
func AddLotteryGifts(gitf *models.LuckybagLottoryGifts) (id int64, err error) {
	o := orm.NewOrm()
	o.Using("update")
	id, err = o.Insert(gitf)
	return
}

//id查询
func GetDeliverIDByUid(uid int) int64 {
	o := orm.NewOrm()
	//根据uid从lotter_user表中查询
	var deiverID string
	criter := fmt.Sprintf("select deliver_id from LotteryUser where id=%d", uid)
	err := o.Raw(criter).QueryRow(&deiverID)

	if err != nil {
		return -1
	}
	if nID, err := strconv.Atoi(deiverID); err == nil {
		return int64(nID)
	}
	fmt.Println("lotteruser data is dirty")
	return -1
}

//编辑/更新活动
func EditLotteryGifts(gitf *models.LuckybagLottoryGifts) (err error) {
	o := orm.NewOrm()
	o.Using("update")
	_, err = o.Update(gitf)
	return

}

//中奖结果显示
func GetWinning() []*models.LuckybagLottoryGiftsLogs {
	var Winning []*models.LuckybagLottoryGiftsLogs
	o := orm.NewOrm()
	o.Using("update")
	_,err := o.Raw("SELECT code,gift_name,date FROM luckybag_lottory_gifts_logs ").QueryRows(&Winning)
	if err != nil{
		beego.Debug("[ADMIN REPORT] get a error:",err.Error())
		return nil
	}
	beego.Debug("[ADMIN REPORT]get a luckybag_lottery_gifts_logs logs:",len(Winning))
	for i := 0; i <len(Winning);i++{
		o := orm.NewOrm()
		o.Using("update")
		var lottory *models.LuckybagLottory
		logs := Winning[i]

		err1 := o.Raw("select id,qx from luckybag_lottory where qx=? ",logs.Code).QueryRow(&lottory)
		if err1 != nil{
			beego.Debug("[ADMIN REPORT]get error1:",err1)
			return nil
		}
		logs.Code = lottory.Qx
		logs.CodeId = lottory.Id
	}
	return Winning
}

//查询中奖名称
func GetLotterywinning(giftname string) (*models.LuckybagLottoryGiftsLogs, error) {
	var winning models.LuckybagLottoryGiftsLogs
	o := orm.NewOrm()
	o.Using("update")

	criter := fmt.Sprintf("select * from luckybag_lottory_gifts_logs where gift_name=%d", giftname)
	err := o.Raw(criter).QueryRow(&winning)

	if err != nil {
		beego.Debug("[ADMIN REPORT] GET a error:", err.Error())
		return nil, err
	}
	return &winning, err

}

//中奖数据导出
func GetWinningQu(startTime, endTime int64) []*models.LuckybagLottoryGiftsLogs {
	var QU []*models.LuckybagLottoryGiftsLogs
	o := orm.NewOrm()
	o.Using("update")
	_, err := o.Raw("select * from luckybag_lottory_gifts_logs where date >= ? and date<?", startTime, endTime).QueryRows(&QU)
	if err != nil {
		beego.Debug("[ADMIN REPORT] get a luckybag gifts error:", err.Error(), "startTime:", startTime, "endTime:", endTime)
		return nil
	}

	return QU
}

//二维码导出
func GetQr(startTime, endTime int64) []*models.LuckybagLottory {
	var QR []*models.LuckybagLottory
	o := orm.NewOrm()
	o.Using("update")
	_, err := o.Raw("select * from luckybag_lottory where used_date>=? and used_date <? ", startTime, endTime).QueryRows(&QR)
	if err != nil {
		beego.Debug("[ADMIN REPORT] get a qr error:", err.Error(), "startTime:", startTime, "endTime:", endTime)
	}
	return QR
}

//地址导出
func GetAddressExcel(startTime, endTime int64) []*models.LuckybagLottoryAddress {
	var Add []*models.LuckybagLottoryAddress
	o := orm.NewOrm()
	o.Using("update")
	_, err := o.Raw("SELECT distinct(logs.gift_name),luck.phone,luck.name,luck.address,luck.email,luck.date FROM "+
		" luckybag_lottory_gifts_logs as logs left JOIN luckybag_lottory_address as luck "+
		" on luck.open_id=logs.open_id where luck.date >= ? and luck.date < ? ", startTime, endTime).QueryRows(&Add)
	if err != nil {
		beego.Debug("[ADMIN REPORT] get a address error:", err.Error(), "startTime:", startTime, "endTime:", endTime)
	}
	return Add

}

//已使用数量查询
func GetUsed(giftname string) []*models.LuckybagLottoryGiftsLogs {
	var useds []*models.LuckybagLottoryGiftsLogs
	o := orm.NewOrm()
	o.Using("upadte")
	_, err := o.Raw("select count(gift_id) from luckybag_lottory_gifts_logs where gift_name = ?", giftname).QueryRows(&useds)
	if err != nil {
		beego.Debug("[ADMIN REPORT] get used gitname err:", err.Error(), "giftname:", giftname)
	}
	beego.Debug("[ADMIN REPORT] get a used err:", err.Error(), "giftname:", giftname)
	return useds
}

//地址修改根据id
func GetAddressById(id int) (*models.LuckybagLottoryAddress , error)  {
	var address models.LuckybagLottoryAddress
	o := orm.NewOrm()
	o.Using("update")
	criter := fmt.Sprintf("select * from luckybag_lottory_address where id=%d",id)
	err := o.Raw(criter).QueryRow(&address)

	if err != nil{
		beego.Debug("[ADMIN REPORT] get a error:",err.Error())
		return nil,err
	}
	return &address,err
}


//删除地址
func RemoveAdderssById(id int) error{
	o := orm.NewOrm()
	o.Using("update")
	criter := fmt.Sprintf("delete from luckybag_lottory_address where id=%d",id)
	_,err := o.Raw(criter).Exec()

	if err != nil{
		beego.Debug("[ADMIN REPORT] Get a error:",err.Error())
		return err
	}
	return err
}

//更新/编辑地址信息
func EditAddress(address *models.LuckybagLottoryAddress)(err error){
	o := orm.NewOrm()
	o.Using("update")
	_,err =o.Update(address,"name","email","phone","address")
	return
}

//更新地址
func AddAddress(address *models.LuckybagLottoryAddress) (id int64,err error) {
	o := orm.NewOrm()
	o.Using("update")
	id,err = o.Insert(address)
	return
}


//记录修改时间
func LotteryGiftLogs(giftlogs *models.LotteryGiftsLogs) (id int64,err error) {
	o := orm.NewOrm()
	o.Using("update")
	id,err = o.Insert(giftlogs)
	return
}

