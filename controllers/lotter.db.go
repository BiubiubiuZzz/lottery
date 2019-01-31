package controllers

import (
	"Lottery/models"
	"fmt"
	"strconv"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)



func GetAddress(deliverId int64) []*models.LuckybagLottoryAddress {
	var address []*models.LuckybagLottoryAddress
	o := orm.NewOrm()
	o.Using("update")
	_, err := o.Raw("SELECT distinct(logs.gift_name),logs.express_no,luck.id,luck.phone,luck.name,luck.address,luck.email,luck.date FROM " +
		" luckybag_lottory_gifts_logs as logs left JOIN " +
		" luckybag_lottory_address as luck on luck.open_id=logs.open_id where luck.deliver_id=?",deliverId).QueryRows(&address)
	if err != nil {
		beego.Debug("[ADMIN REPORT] GET a address manager error:", err.Error())
		return nil
	}
	beego.Debug("[ADMIN REPORT] get a AddressManager：", len(address))
	return address
}


//查询地址

//func GetAddress() []*models.LuckybagLottoryAddress {
//	var address []*models.LuckybagLottoryAddress
//	o := orm.NewOrm()
//	o.Using("update")
//	_, err := o.Raw("SELECT open_id,phone,name,email,date,address from luckybag_lottory_address").QueryRows(&address)
//	if err != nil {
//		beego.Debug("[ADMIN REPORT] GET a address manager error:", err.Error())
//		return nil
//	}
//	beego.Debug("[ADMIN REPORT] get a AddressManager：", len(address))
//	for i := 0; i <len(address); i++ {
//		o := orm.NewOrm()
//		o.Using("update")
//		var giftlogs *models.LuckybagLottoryGiftsLogs
//		addr := address[i]
//
//		err1 := o.Raw("SELECT distinct(gift_name),open_id FROM luckybag_lottory_gifts_logs where open_id =?",addr.OpenId).QueryRow(&giftlogs)
//		if err1 != nil{
//			beego.Debug("[ADMIN REPORT] get error:",err1)
//			return nil
//		}
//		addr.OpenId = giftlogs.OpenId
//		addr.GiftName = giftlogs.GiftName
//	}
//
//	return address
//}


//***注：所有QR表示抽奖码；
//全部抽奖码显示
func GetQR(deliverID int64)[]*models.LuckybagLottory  {
	var QR []*models.LuckybagLottory
	o := orm.NewOrm()
	o.Using("update")
	_,err := o.Raw("select *  from luckybag_lottory where deliver_id =? ",deliverID).QueryRows(&QR)
	if err != nil{
		beego.Debug("[ADMIN REPORT] get a error:",err.Error())
		return nil
	}
	beego.Debug("[ADMIN REPORT] get winning:",len(QR))
	return QR
}

//显示红包
func RedPack(deliverID int64) []*models.LuckybagLottoryRedpack  {
	var red []*models.LuckybagLottoryRedpack
	o := orm.NewOrm()
	o.Using("update")
	_,err := o.Raw("SELECT red.fee,red.code,red.err_msg,red.date,gift.gift_name from luckybag_lottory_redpack as red left join luckybag_lottory_gifts as gift on red.gift_id=gift.id where red.deliver_id= ? ",deliverID).QueryRows(&red)
	if err != nil{
		beego.Debug("[ADMIN REPORT] get a error:",err.Error())
		return nil
	}
	beego.Debug("[ADMIN REPORT] get redpack:",len(red))
	return red
}

//查询code显示red result
func RedPackQuery(code string) []*models.LuckybagLottoryRedpack   {
	var Rcode []*models.LuckybagLottoryRedpack
	o := orm.NewOrm()
	o.Using("update")
	_,err := o.Raw("SELECT red.date,red.fee,red.code,red.err_msg,gift.gift_name from luckybag_lottory_redpack as red " +
		" left join luckybag_lottory_gifts as gift on red.gift_id=gift.id where red.code = ? ",code).QueryRows(&Rcode)
	if err != nil {
		beego.Debug("[ADMIN REPORT] get a error:",err.Error())
		return nil
	}
	beego.Debug("[ADMNIN REPORT] get a recode:",code)
	return Rcode
}

//func RedPack() []*models.LuckybagLottoryRedpack  {
//	var red []*models.LuckybagLottoryRedpack
//	o := orm.NewOrm()
//	o.Using("update")
//	_,err := o.Raw("SELECT gift_id,fee,code from luckybag_lottory_redpack").QueryRows(&red)
//	if err != nil{
//		beego.Debug("[ADMIN REPORT] get a error:",err.Error())
//		return nil
//	}
//	beego.Debug("[ADMIN REPORT] get redpack:",len(red))
//	for i := 0; i <len(red); i++{
//		o := orm.NewOrm()
//		o.Using("update")
//		var gift *models.LuckybagLottoryGifts
//		redpack := red[i]
//
//		err1 := o.Raw("select id,gift_name from luckybag_lottory_gifts where id=? ",redpack.GiftId).QueryRow(&gift)
//		if err1 != nil{
//			beego.Debug("[ADMIN REPORT] get error1:",err1)
//			return nil
//		}
//		redpack.GiftId = gift.Id
//	}
//	return red
//}
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

//查询实物使用总数
func GetGiftUsedByGiftID(giftID int64) int64 {
	var result int64 = 0
	o := orm.NewOrm()
	o.Using("update")
	sql := fmt.Sprintf("select count(*) as used from luckybag_lottory_gifts_logs where gift_id=%d ", giftID)
	err := o.Raw(sql).QueryRow(&result)
	if err != nil{
		beego.Debug("[ADMIN REPORT]Get a gifts used err:",err.Error())
	}
	sql1 := fmt.Sprintf("select count(*) as used from luckybag_lottory_redpack where gift_id=%d ",giftID)
	err1 := o.Raw(sql1).QueryRow(&result)
	if err != nil{
		beego.Debug("[ADMIN REPORT]get a red use err:",err1.Error())
	}

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

//显示活动设置的数据--根据用户deliver_id显示
func GetActivity(deliverID int64) []*models.LuckybagLottoryGifts {
	var AC []*models.LuckybagLottoryGifts
	var result []*models.LuckybagLottoryGifts = nil
	o := orm.NewOrm()
	o.Using("update")
	_, err := o.Raw("select * from luckybag_lottory_gifts where deliver_id =? ",deliverID).QueryRows(&AC)
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

//添加活动活动
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

//更新活动
func EditLotteryGifts(gitf *models.LuckybagLottoryGifts) (err error) {
	o := orm.NewOrm()
	o.Using("update")
	_, err = o.Update(gitf)
	return

}

//中奖结果显示
func GetWinning(deliverID int64) []*models.LuckybagLottoryGiftsLogs {
	var Winning []*models.LuckybagLottoryGiftsLogs
	o := orm.NewOrm()
	o.Using("update")
	_,err := o.Raw("SELECT code,gift_name,date FROM luckybag_lottory_gifts_logs where deliver_id=?",deliverID).QueryRows(&Winning)
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
	criter := fmt.Sprintf("SELECT distinct(logs.gift_name),logs.express_no,luck.id," +
		"luck.phone,luck.name,luck.address,luck.email,luck.date " +
		" FROM luckybag_lottory_gifts_logs as logs left JOIN  " +
		"luckybag_lottory_address as luck on luck.open_id=logs.open_id where luck.id=%d",id)
	err := o.Raw(criter).QueryRow(&address)

	if err != nil{
		beego.Debug("[ADMIN REPORT] get a error:",err.Error())
		return nil,err
	}
	return &address,err
}

//func GetGiftLogsExpressNo(openid string) (*models.LuckybagLottoryGiftsLogs,error)  {
//	var expressno models.LuckybagLottoryGiftsLogs
//	o := orm.NewOrm()
//	o.Using("update")
//	criter := fmt.Sprintf("select * from luckbag_lottory_address where open_id = ?",openid)
//	err := o.Raw(criter).QueryRow(&expressno)
//	if err != nil{
//		beego.Debug("[ADMIN REPORT] get a error:",err.Error())
//		return nil,err
//	}
//	return &expressno,err
//
//}


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
	err =o.Raw("UPDATE `luckybag_lottory_address` as address LEFT JOIN `luckybag_lottory_gifts_logs` as logs "+
	" on address.open_id=logs.open_id " +
	" SET address.`name` = ?, address.`email` = ? , address.`phone` = ? , address.`address` = ?,  logs.express_no =? " +
	"WHERE address.`id` = ?",address.Name,address.Email,address.Phone,address.Address,address.ExpressNo,address.Id).QueryRow(&address)
	if err!= nil {
		beego.Debug("[ADMIN REPORT] Get a error",err.Error())
		return nil
	}
	return err
}



//更新地址
func AddAddress(address *models.LuckybagLottoryAddress) (id int64,err error) {
	o := orm.NewOrm()
	o.Using("update")
	id,err = o.Insert(address)
	return
}

//重置语句--暂时不用
//func Reset(Id *models.LuckybagLottory) (err error) {
//	o := orm.NewOrm()
//	o.Using("upadte")
//	_,err = o.Update(Id,"Method","UsedDate")
//	return
//}

//导出QR已使用的编码
func QRused(starTime,endTime int64) []*models.LuckybagLottory {
	var use []*models.LuckybagLottory
	o := orm.NewOrm()
	o.Using("update")
	_,err := o.Raw("select * from luckybag_lottory where used_date>0 and used_date>=? and used_date<? ",starTime,endTime).QueryRows(&use)
	if err != nil{
		beego.Debug("[ADMIN REPORT] get a qruse error:",err.Error(),"startTime:",starTime,"endTime:",endTime)
	}
	return use
}

//导出QR未使用的编码
func QRNotUsed(starTime,endTime int64) []*models.LuckybagLottory {
	var notuse []*models.LuckybagLottory
	o := orm.NewOrm()
	o.Using("update")
	_,err := o.Raw("select * from luckybag_lottory where used_date=0 and created_date>=? and created_date<? ",starTime,endTime).QueryRows(&notuse)
	if err != nil{
		beego.Debug("[ADMIN REPORT] get a qruse error:",err.Error(),"startTime:",starTime,"endTime:",endTime)
	}
	return notuse
}



