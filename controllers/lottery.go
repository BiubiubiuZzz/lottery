package controllers

import (
	"Lottery/models"
	"errors"
	"fmt"
	"github.com/astaxie/beego/orm"
	"strconv"
	"time"

	"github.com/astaxie/beego"
	"github.com/tealeg/xlsx"
)

var FACTOR int = 1000000
var FAC		int = 100

type LotteryController struct {
	baseController
}
//分页
func (this *LotteryController) GetQrcode() {
	Logined := this.GetSession("login")
	if Logined != true {
		this.Redirect("/login", 302)
		return
	}
	list := GetQR()
	o := orm.NewOrm()
	o.Using("update")
	query := o.QueryTable(new(models.LuckybagLottory))
	count,_ := query.Count()
	query.OrderBy("-id").Limit(this.pager.Pagesize,(this.pager.Page-1)*this.pager.Pagesize).All(&list)

	this.Data["qrlist"] =list
	this.pager.SetTotalnum(int(count))
	this.pager.SetUrlpath("/qrcode/query?page=%d")
	this.Data["pagebar"] = this.pager.ToString()
	this.display()

	this.Data["queryValue"] = ""
	this.TplName = "QRcode_query.html"
}

//抽奖码---ID 查询
func (this *LotteryController) QueryById() {
	Logined := this.GetSession("login")
	if Logined != true {
		this.Redirect("/login", 302)
		return
	}
	Id := this.Input().Get("Id")
	if Id != "" {
		qrlist := GetQRcode(Id)
		if qrlist != nil {
			var oList []models.LuckybagLottoryOutput
			for _, v := range qrlist {
				var t models.LuckybagLottoryOutput
				t.Id = v.Id
				t.Qx = v.Qx
				t.DeliverId = v.DeliverId
				t.Url = v.Url
				t.BatchNo = v.Id
				t.Method = v.Method
				t.CreatedDate = v.CreatedDate
				t.UsedDate = v.UsedDate
				//t.GiftName =v.GiftName
				oList = append(oList, t)
			}
			this.Data["qrlist"] = oList
		}
	}
	this.Data["queryValue"] = Id
	this.Data["pagebar"] = this.pager.ToString()
	this.TplName = "QRcode_query.html"

}

//奖品名称查询
func (this *LotteryController) SettingQuery() {
	Logined := this.GetSession("login")
	if Logined != true {
		this.Redirect("/login", 302)
		return
	}
	priceName := this.Input().Get("PriceName")

	if priceName != "" {
		list := GetActivityByName(priceName)
		this.Data["activitylists"] = list
	}
	this.Data["queryValue"] = priceName
	this.TplName = "Prize_setting.html"

}

//中奖结果查询 ---Id查询
func (this *LotteryController) WinnQuery() {
	Logined := this.GetSession("login")
	if Logined != true {
		this.Redirect("/login", 302)
		return
	}
	Id:= this.Input().Get("Id")
	if Id != "" {
		list := GetWinningByCodeId(Id)
		this.Data["winninglist"] = list
	}
	this.Data["queryValue"] = Id
	this.TplName = "winning_query.html"

}

//地址中奖名称查询
func (this *LotteryController) AddressQuery() {
	Logined := this.GetSession("login")
	if Logined != true {
		this.Redirect("/login", 302)
		return
	}
	pricename := this.Input().Get("PriceName")
	if pricename != "" {
		list := GetAddressQuser(pricename)
		this.Data["lists"] = list
	}
	this.Data["queryValue"] = pricename
	this.TplName = "Address_management.html"

}

//中奖结果
func (this *LotteryController) GetWinning() {
	Logined := this.GetSession("login")
	if Logined != true {
		this.Redirect("/login", 302)
		return
	}
	this.TplName = "Winning_result.html"

	list := GetWinning()
	this.Data["winninglist"] = list

}

//活动显示
func (this *LotteryController) GetPrize() {
	Logined := this.GetSession("login")
	if Logined != true {
		this.Redirect("/login", 302)
		return
	}
	list := GetActivity()
	if len(list) > 0 {
		var display []models.LuckybagLottoryGiftsDisplay
		for _, o := range list {
			var t models.LuckybagLottoryGiftsDisplay
			t.Id = o.Id
			t.DeliverId = o.DeliverId
			t.GiftName = o.GiftName
			t.GiftPic = o.GiftPic
			t.Fee = float64(float64(o.Fee) / float64(FAC))
			t.Odds = float64(float64(o.Odds) / float64(FACTOR))
			t.OddsBase = o.OddsBase
			t.OddsTop = o.OddsTop
			t.Valid = o.Valid
			t.Method = o.Method
			t.Quantity = o.Quantity
			t.Date = o.Date
			t.Used = o.Used
			t.LeftQuantity = o.LeftQuantity
			t.Total = int64(int64(o.LeftQuantity) + int64(o.Used))
			display = append(display, t)
		}
		this.Data["activitylists"] = display
	}

	this.TplName = "Prize_setting.html"
}

//地址
func (this *LotteryController) GetAddress() {
	Logined := this.GetSession("login")
	if Logined != true {
		this.Redirect("/login", 302)
		return
	}
	list := GetAddress()
	this.Data["lists"] = list
	this.TplName = "Address_management.html"

}

//显示微信红包
func (this *LotteryController) GetRedPack()  {
	Logined := this.GetSession("login")
	if Logined != true {
		this.Redirect("/login", 302)
		return
	}
	list := RedPack()
	if len(list) > 0{
		var display []models.LuckybagLottoryRedpackDisplay
		for _,o := range list{
			var t models.LuckybagLottoryRedpackDisplay
			t.GiftName = o.GiftName
			t.Fee =float64(float64(o.Fee) / float64(FAC))
			t.Code = o.Code
			t.ErrMsg =o.ErrMsg
			display = append(display,t)
		}
		this.Data["redlist"] = display
	}
	this.TplName ="redpack.html"
}

//code查询
func (this *LotteryController) RedPackQuery()  {
	Logined := this.GetSession("login")
	if Logined != true {
		this.Redirect("/login", 302)
		return
	}
	code := this.Input().Get("code")
	if code != "" {
		list := RedPackQuery(code)
		this.Data["redlist"] = list
	}
	this.Data["radValue"] = code
	this.TplName ="redpack.html"
}


//编辑用户地址信息
func (this *LotteryController) SettingAddress() {
	Logined := this.GetSession("login")
	if Logined != true {
		this.Redirect("/login", 302)
		return
	}
	id,err := this.GetInt("id",-1)
	this.Data["edit"] = false
	if id != -1 && err == nil{
		address,err := GetAddressById(id)
		if err == nil {
			this.Data["name"] = address.Name
			this.Data["phone"] = address.Phone
			this.Data["email"] = address.Email
			this.Data["address"] = address.Address
			this.Data["edit"] =true
			this.Data["id"] = id
		}
	}
	this.TplName = "Address_userEdit.html"

}

//修改地址点击保存过程
func (this *LotteryController) SaveAddress() {
	Logined := this.GetSession("login")
	if Logined != true {
		this.Redirect("/login", 302)
		return
	}
	edit,_ := this.GetBool("edit",false)
	var address models.LuckybagLottoryAddress

	address.Name = this.GetString("name")
	address.Phone = this.GetString("phone")
	address.Email = this.GetString("email")
	address.Address = this.GetString("address")

	if edit == false{
		_,err := AddAddress(&address)
		if err == nil{
			this.Redirect("/address/management",302)
		}else {
			this.TplName ="Address_userEdit.html"
		}
	} else {
		id ,err := this.GetInt64("id",-1)
		if id != -1 && err ==nil{
			address.Id =id
			err = EditAddress(&address)
			if err == nil{
				this.Data["success"] = 1
				this.TplName ="Success_AddressEdit.html"
			}else{
				this.Data["success"] = 0
				this.TplName = "Success_AddressEdit.html"
			}
		}
	}

}

//删除用户地址信息
func (this *LotteryController) RemoveAddress() {
	Logined := this.GetSession("login")
	if Logined != true {
		this.Redirect("/login", 302)
		return
	}
	id,err := this.GetInt("id",-1)
	if id != -1 && err == nil{
		err := RemoveAdderssById(id)
		if err == nil {
			this.Data["success"] = 1
			this.TplName ="Delete_address.html"
		}else {
			this.Data["success"] = 0
			this.TplName ="Delete_address.html"
		}
	}else{
		this.Data["success"] =0
		this.TplName= "Delete_address.html"
	}
}


//添加活动--Show original quantity
func (this *LotteryController) Setting() {
	Logined := this.GetSession("login")
	if Logined != true {
		this.Redirect("/login", 302)
		return
	}
	id, err := this.GetInt("id", -1)
	this.Data["edit"] = false
	if id != -1 && err == nil {
		gift, err := GetLotteryGiftByID(id)
		if err == nil {
			this.Data["giftname"] = gift.GiftName
			this.Data["fee"] = float64(float64(gift.Fee) / float64(FAC))
			this.Data["quantity"] = GetLeftQuantity(gift.Id)
			this.Data["odds"] = float64(float64(gift.Odds) / float64(FACTOR))
			this.Data["valid"] = gift.Valid
			this.Data["method"] = gift.Method
			this.Data["edit"] = true
			this.Data["id"] = id
		}else {
			this.Data["msg"] = err.Error()
		}
	}
	//this.Data["activitylists"] =list
	this.TplName = "Activity_settings.html"
}

//编辑活动---Show remaining quantity
func (this *LotteryController) EditeSetting() {
	Logined := this.GetSession("login")
	if Logined != true {
		this.Redirect("/login", 302)
		return
	}
	id, err := this.GetInt("id", -1)
	this.Data["edit"] = false
	if id != -1 && err == nil {
		gift, err := GetLotteryGiftByID(id)
		if err == nil {
			this.Data["giftname"] = gift.GiftName
			this.Data["fee"] = float64(float64(gift.Fee) / float64(FAC))
			this.Data["quantity"] = GetLeftQuantity(gift.Id)
			this.Data["odds"] = float64(float64(gift.Odds) / float64(FACTOR))
			this.Data["valid"] = gift.Valid
			this.Data["method"] = gift.Method
			this.Data["edit"] = true
			this.Data["id"] = id
		}else {
			this.Data["msg"] = err.Error()
		}
	}
	this.TplName = "Activity_EditeSettings.html"
}

//动态计算概率
func (this *LotteryController) calculateOdd(deliverID int, odds int64) (base int64, top int64, err error) {
	allGifts, _ := GetLotteryGiftByDeliverID(deliverID)
	//第一条记录
	if len(allGifts) == 0 {
		base = 0
		top = odds
		err = nil
		return
	} else { //如果不是第一条记录， 则寻找前面有效的记录
		var i int
		var sum, whole int64
		sum = 0
		for i = 0; i < len(allGifts); i++ {
			//计算所有valid 的数据的概率和
			if allGifts[i].Valid == 1 {
				sum += allGifts[i].Odds
			}

		}
		sum += odds

		//100 % , 如果加上当前的概率值大于 100% 则认定失败
		//whole = int64(100) * int64(FACTOR)
		whole = int64(float64(100.1) * float64(FACTOR))
		if sum > whole {
			base = 0
			top = 0
			err = errors.New("概率值大于 100%")
			return
		}

		n := len(allGifts)
		//寻找前面有效的 gift 记录
		for i = n - 1; i >= 0; i-- {
			if allGifts[i].Valid == 1 {
				break
			}
		}
		//找到了
		if i >= 0 {
			base = allGifts[i].OddsTop
			top = odds + base
			err = nil
			return
		}
		//没有找到前面有效的记录

		base = 0
		top = odds
		err = nil
		return
	}
}

//修改gift 的时候概率处理算法
func (this *LotteryController) modifyGift(deliverID int, gift *models.LuckybagLottoryGifts) error {
	allGifts, _ := GetLotteryGiftByDeliverID(deliverID)
	var i, foundIdx int
	var sum, whole int64
	sum = 0
	for i = 0; i < len(allGifts); i++ {
		if gift.Id == allGifts[i].Id {
			break
		}
	}

	if i < 0 || i >= len(allGifts) {
		return errors.New("没有发现记录")
	}

	//保持时间不变
	allGifts[i].GiftName = gift.GiftName
	allGifts[i].Fee = gift.Fee
	//如果数量不同，则修改修改时间, 否则保持时间不变
	if allGifts[i].Quantity != gift.Quantity {
		allGifts[i].Date = gift.Date
	}

	allGifts[i].Quantity = gift.Quantity
	allGifts[i].Odds = gift.Odds
	allGifts[i].Valid = gift.Valid
	allGifts[i].Method = gift.Method

	idx := i
	foundIdx = i
	for i = 0; i < len(allGifts); i++ {
		//计算所有valid 的数据的概率和
		if allGifts[i].Valid == 1 {
			sum += allGifts[i].Odds
		}

	}

	//100 % , 如果加上当前的概率值大于 100% 则认定失败
	//whole = int64(100) * int64(FACTOR)
	whole = int64(float64(100.1) * float64(FACTOR))
	if sum > whole {
		return errors.New("概率值大于 100%")
	}

	var pi int
	if idx == 0 {
		idx = idx + 1

		allGifts[0].OddsBase = 0
		allGifts[0].OddsTop = allGifts[0].OddsBase + allGifts[0].Odds
	}

	//修改后面每一条记录
	for i = idx; i < len(allGifts); i++ {
		if allGifts[i].Valid == 0 {
			continue
		}

		for pi = i - 1; pi >= 0; pi-- {
			if allGifts[pi].Valid == 1 {
				break
			}
		}

		if pi < 0 {
			allGifts[i].OddsBase = 0
			allGifts[i].OddsTop = allGifts[i].OddsBase + allGifts[i].Odds
		} else {
			allGifts[i].OddsBase = allGifts[pi].OddsTop
			allGifts[i].OddsTop = allGifts[i].OddsBase + allGifts[i].Odds
		}
	}

	for i := foundIdx; i < len(allGifts); i++ {
		err := EditLotteryGifts(&allGifts[i])
		if err != nil {
			return err
		}
	}

	return nil
}

//新增活动---点击过程
func (this *LotteryController) SaveGift() {
	Logined := this.GetSession("login")
	if Logined != true {
		this.Redirect("/login", 302)
		return
	}


	edit, _ := this.GetBool("edit", false)
	var gift models.LuckybagLottoryGifts


	uid, _ := this.GetSession("uid").(int)
	fmt.Println(uid)

	//根据id查询deliver_id
	deliverID := GetDeliverIDByUid(uid)
	gift.GiftName = this.GetString("giftname", "")


	if gift.GiftName == "" {
		this.Data["msg"] = "请输入商品名称"
		this.TplName = "Activity_settings.html"
		return
	}
	gift.Quantity, _ = this.GetInt64("quantity", 0)


	if gift.Quantity == 0 {
		this.Data["msg"] = "请输入商品数量"
		this.TplName = "Activity_settings.html"
		return
	}
	Fee,_ :=this.GetFloat("fee",0.0)
	gift.Fee = int64(float64(Fee) * float64(FAC))

	odds, _ := this.GetFloat("odds", 0.0)
	gift.Odds = int64(float64(odds) * float64(FACTOR))

	gift.Valid, _ = this.GetInt64("valid")

	gift.Method, _ = this.GetInt64("method")
	gift.Date = time.Now().Unix()

	gift.DeliverId = deliverID

	//如果设置的概率为 0 ， 则设置该记录无效
	if gift.Odds == 0 {
		gift.Valid = 0
	}

	//添加gift setting 的记录
	if edit == false {

		if gift.Valid == 0 {
			this.Data["msg"] = "新加活动必须是有效的"
			this.TplName = "Activity_settings.html"
			return
		}

		base, top, err := this.calculateOdd(int(deliverID), gift.Odds)
		if err == nil {

			gift.OddsBase = base
			gift.OddsTop = top

			_, err = AddLotteryGifts(&gift)

			if err == nil {
				this.Redirect("/Prize/setting", 302)
			} else {
				this.Data["msg"] = err.Error()
				this.TplName = "Activity_settings.html"
			}

		} else {
			this.Data["msg"] = err.Error()
			this.TplName = "Activity_settings.html"
		}

	} else {
		id, err := this.GetInt64("id", -1)
		if id != -1 && err == nil {
			gift.Id = id
			//记录商品数量修改logs
			o := orm.NewOrm()
			o.Using("update")
			giftslogs := new(models.LotteryGiftsLogs)
			giftslogs.GiftName = gift.GiftName
			giftslogs.DeliverId = gift.DeliverId
			giftslogs.Quantity = gift.Quantity
			giftslogs.Date = gift.Date
			giftslogs.GiftId = gift.Id
			o.Insert(giftslogs)

			err = this.modifyGift(int(deliverID), &gift)


			if err == nil {
				this.Data["success"] = 1
				this.TplName = "Successful_editing.html"
			} else {
				this.Data["success"] = 0
				this.TplName = "Successful_editing.html"
			}
		}
	}


}



//修改gift 的时候概率处理算法
func (this *LotteryController) removeGift(deliverID int, recordID int64) error {
	allGifts, _ := GetLotteryGiftByDeliverID(deliverID)
	var i int

	for i = 0; i < len(allGifts); i++ {
		if recordID == allGifts[i].Id {
			break
		}
	}

	if i < 0 {
		return errors.New("not found")
	}

	idx := i

	//修改后面每一条记录
	for i = idx + 1; i < len(allGifts); i++ {
		if allGifts[i].Valid == 0 {
			continue
		}

		if i == 1 && idx == 0 {
			allGifts[i].OddsBase = 0
			allGifts[i].OddsTop = allGifts[i].OddsBase + allGifts[i].Odds
		} else {
			var pi int
			for pi = i - 1; pi >= 0; pi-- {
				if allGifts[pi].Valid == 1 && allGifts[pi].Id != recordID {
					break
				}
			}
			if pi >= 0 {
				allGifts[i].OddsBase = allGifts[pi].OddsTop
				allGifts[i].OddsTop = allGifts[i].OddsBase + allGifts[i].Odds
			} else {
				allGifts[i].OddsBase = 0
				allGifts[i].OddsTop = allGifts[i].OddsBase + allGifts[i].Odds
			}

		}
	}

	for i := 0; i < len(allGifts); i++ {
		if allGifts[i].Id == recordID {
			RemoveLotteryGiftByID(recordID)
			continue
		}

		err := EditLotteryGifts(&allGifts[i])
		if err != nil {
			return err
		}
	}

	return nil
}

//删除活动
func (this *LotteryController) RemoveGift() {
	Logined := this.GetSession("login")
	if Logined != true {
		this.Redirect("/login", 302)
		return
	}

	uid, _ := this.GetSession("uid").(int)
	fmt.Println(uid)

	//根据id查询deliver_id
	deliverID := GetDeliverIDByUid(uid)

	id, err := this.GetInt("id", -1)
	if id != -1 && err == nil {
		err := this.removeGift(int(deliverID), int64(id))
		if err == nil {
			//this.Ctx.WriteString("<h1>删除记录 成功</h1>")
			this.Data["success"] = 1
			this.TplName = "Delete_result.html"
		} else {
			this.Data["success"] = 0
			this.TplName = "Delete_result.html"
		}
	} else {
		this.Data["success"] = 0
		this.TplName = "Delete_result.html"
	}

}

//点击按钮导出
func (this *LotteryController) Getqr() {
	this.TplName = "index.html"

	scope := this.Input().Get("scope")
	method := this.Input().Get("method")
	start := this.Input().Get("start") + ":00"
	end := this.Input().Get("end") + ":00"
	s, _ := time.Parse("2006/01/02 15:04:05", start)
	e, _ := time.Parse("2006/01/02 15:04:05", end)
	startTime := s.Unix() - 8*3600
	endTime := e.Unix() - 8*3600
	if startTime >= endTime {
		beego.Debug("[ADMIN REPORT] startTime >= EndTime")
		return
	} else if startTime < endTime {
		beego.Debug("[ADMIN REPORT] startTime < endTime")
	}

	beego.Debug("[ADMIN REPORT] get span report for method:", method, "start:", start, "end:", end, "start Unix:", startTime, "end Unix:", endTime)

	if scope == "winnder" {
		WinnerRow := GetWinningQu(startTime, endTime)
		WinnerFile := GenWinnerExcel(WinnerRow, startTime, endTime)
		beego.Debug("[ADMIN REPORT]get a winning File:", WinnerFile)
		this.Ctx.WriteString(WinnerFile)
		return
	}
	if scope == "qr" {
		QrRow := GetQr(startTime, endTime)
		QrFile := GenQrExcel(QrRow, startTime, endTime)
		beego.Debug("[ADMIN REPORT] get a qr file:", QrFile)
		this.Ctx.WriteString(QrFile)
		return
	}

	if scope == "addr" {
		AddressRow := GetAddressExcel(startTime, endTime)
		AddressFile := GenAddressExcel(AddressRow, startTime, endTime)
		beego.Debug("[ADMIN REPORT] get a Address file:", AddressFile)
		this.Ctx.WriteString(AddressFile)
		return
	}

	this.Ctx.WriteString("empty")
	return

}

//中奖结果导出
func GenWinnerExcel(data []*models.LuckybagLottoryGiftsLogs, startTime, endTime int64) string {
	var file *xlsx.File
	var sheet *xlsx.Sheet
	var row *xlsx.Row
	var timespan *xlsx.Cell
	var code *xlsx.Cell
	//var method		*xlsx.Cell
	var giftname *xlsx.Cell
	var date *xlsx.Cell
	var err error

	file = xlsx.NewFile()
	sheet, err = file.AddSheet("中奖结果")
	if err != nil {
		beego.Debug("[ADMIN REPORT] new xls file err:", err.Error())
		return ""
	}
	row = sheet.AddRow()
	timespan = row.AddCell()
	timespan.Value = "时间段统计"
	code = row.AddCell()
	code.Value = "抽奖编码"
	giftname = row.AddCell()
	giftname.Value = "奖品名称"
	//method = row.AddCell()
	//method.Value ="实物类型"
	date = row.AddCell()
	date.Value = "中奖时间"
	//抽奖编码	是否领取

	t1 := time.Unix(startTime, 0)
	t2 := time.Unix(endTime, 0)
	start := t1.Format("2006/01/02 15:04:05")
	end := t2.Format("2006/01/02 15:04:05")
	selectedTime := start + "-" + end

	for i := 0; i < len(data); i++ {
		row = sheet.AddRow()
		timespan = row.AddCell()
		timespan.Value = selectedTime

		code = row.AddCell()
		code.Value = data[i].Code

		giftname = row.AddCell()
		giftname.Value = data[i].GiftName

		//method = row.AddCell()
		//method.Value = strconv.FormatInt(data[i].Method,10)

		t3 := time.Unix(data[i].Date, 0)
		time := t3.Format("2006/01/02 15:04:05")

		date = row.AddCell()
		date.Value = time

	}
	fileName := "./static/tmp/winning/winning_" + strconv.FormatInt(time.Now().Unix(), 10) + ".xlsx"
	err = file.Save(fileName)
	if err != nil {
		beego.Debug("[ADMIN REPORT] save span report err:", err.Error())
		return ""
	}
	beego.Debug("[ADMIN REPORT] save span report excel file ok!filename:", fileName)
	return fileName[1:]

}

//QRcode导出
func GenQrExcel(data []*models.LuckybagLottory, startTime, endTime int64) string {
	var file *xlsx.File
	var sheet *xlsx.Sheet
	var row *xlsx.Row
	var timespan *xlsx.Cell
	var id *xlsx.Cell
	var qx *xlsx.Cell
	var url *xlsx.Cell
	var createddate *xlsx.Cell
	var useddate *xlsx.Cell
	var method *xlsx.Cell
	var err error

	file = xlsx.NewFile()
	sheet, err = file.AddSheet("中奖结果")
	if err != nil {
		beego.Debug("[ADMIN REPORT] new xls file err:", err.Error())
		return ""
	}
	row = sheet.AddRow()
	timespan = row.AddCell()
	timespan.Value = "时间段统计"
	method = row.AddCell()
	method.Value = "实物类型(1:红包;2:实物)"
	id = row.AddCell()
	id.Value = "Id"
	qx = row.AddCell()
	qx.Value = "抽奖编码"
	url = row.AddCell()
	url.Value = "二维码链接"
	createddate = row.AddCell()
	createddate.Value = "创建时间"
	useddate = row.AddCell()
	useddate.Value = "使用时间"

	t1 := time.Unix(startTime, 0)
	t2 := time.Unix(endTime, 0)
	start := t1.Format("2006/01/02 15:04:05")
	end := t2.Format("2006/01/02 15:04:05")
	selectedTime := start + "-" + end

	for i := 0; i < len(data); i++ {
		row = sheet.AddRow()
		timespan = row.AddCell()
		timespan.Value = selectedTime

		method = row.AddCell()
		method.Value = strconv.FormatInt(data[i].Method, 10)

		id = row.AddCell()
		id.Value = strconv.FormatInt(data[i].Id, 10)

		qx = row.AddCell()
		qx.Value = data[i].Qx

		url = row.AddCell()
		url.Value = data[i].Url

		t3 := time.Unix(data[i].CreatedDate, 0)
		creadtime := t3.Format("2006/01/02 15:04:05")

		createddate = row.AddCell()
		createddate.Value = creadtime

		t4 := time.Unix(data[i].UsedDate, 0)
		usedtime := t4.Format("2006/01/02 15:04:05")

		useddate = row.AddCell()
		useddate.Value = usedtime

	}
	fileName := "./static/tmp/qr/qr_" + strconv.FormatInt(time.Now().Unix(), 10) + ".xlsx"
	err = file.Save(fileName)
	if err != nil {
		beego.Debug("[ADMIN REPORT] save span report err:", err.Error())
		return ""
	}
	beego.Debug("[ADMIN REPORT] save span report excel file ok!filename:", fileName)
	return fileName[1:]

}

//地址导出
func GenAddressExcel(data []*models.LuckybagLottoryAddress, startTime, endTime int64) string {
	var file *xlsx.File
	var sheet *xlsx.Sheet
	var row *xlsx.Row
	var timespan *xlsx.Cell
	var giftname *xlsx.Cell
	var name *xlsx.Cell
	var email *xlsx.Cell
	var phone *xlsx.Cell
	var address *xlsx.Cell
	var date *xlsx.Cell
	var err error

	file = xlsx.NewFile()
	sheet, err = file.AddSheet("中奖结果")
	if err != nil {
		beego.Debug("[ADMIN REPORT] new xls file err:", err.Error())
		return ""
	}
	row = sheet.AddRow()
	timespan = row.AddCell()
	timespan.Value = "时间段统计"
	giftname = row.AddCell()
	giftname.Value = "中奖名称"
	name = row.AddCell()
	name.Value = "名字"
	phone = row.AddCell()
	phone.Value = "电话号码"
	email = row.AddCell()
	email.Value = "邮箱地址"
	address = row.AddCell()
	address.Value = "收货地址"
	date = row.AddCell()
	date.Value = "时间"

	t1 := time.Unix(startTime, 0)
	t2 := time.Unix(endTime, 0)
	start := t1.Format("2006/01/02 15:04:05")
	end := t2.Format("2006/01/02 15:04:05")
	selectedTime := start + "-" + end

	for i := 0; i < len(data); i++ {
		row = sheet.AddRow()
		timespan = row.AddCell()
		timespan.Value = selectedTime

		giftname = row.AddCell()
		giftname.Value = data[i].GiftName

		name = row.AddCell()
		name.Value = data[i].Name

		phone = row.AddCell()
		phone.Value = data[i].Phone

		email = row.AddCell()
		email.Value = data[i].Email

		address = row.AddCell()
		address.Value = data[i].Address

		t3 := time.Unix(data[i].Date, 0)
		time := t3.Format("2006/01/02 15:04:05")

		date = row.AddCell()
		date.Value = time

	}
	fileName := "./static/tmp/address/address_" + strconv.FormatInt(time.Now().Unix(), 10) + ".xlsx"
	err = file.Save(fileName)
	if err != nil {
		beego.Debug("[ADMIN REPORT] save span report err:", err.Error())
		return ""
	}
	beego.Debug("[ADMIN REPORT] save span report excel file ok!filename:", fileName)
	return fileName[1:]

}
