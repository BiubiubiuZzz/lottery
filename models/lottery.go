package models

type LuckybagLottoryAddress struct {
	Id       int64
	GiftName string
	OpenId   string
	Name     string
	Email    string
	Phone    string
	Address  string
	Date     int64
}

type LuckybagLottoryGifts struct {
	Id           int64
	DeliverId    int64
	GiftName     string
	GiftPic      string
	Fee          int64
	Odds         int64
	OddsBase     int64
	OddsTop      int64
	Valid        int64
	Method       int64
	Quantity     int64 //数量
	Date         int64
	Used         int64 `orm:"-"` //used number
	LeftQuantity int64 `orm:"-"`
	Total         int64 `orm:"-"` //总数
}

type LuckybagLottoryGiftsDisplay struct {
	Id           int64
	DeliverId    int64
	GiftName     string
	GiftPic      string
	Fee          float64
	Odds         float64
	OddsBase     int64
	OddsTop      int64
	Valid        int64
	Method       int64
	Quantity     int64 //数量
	Date         int64
	Used         int64 `orm:"-"` //used number
	LeftQuantity int64 `orm:"-"`
	Total         int64 `orm:"-"`
}

type LuckybagLottoryGiftsLogs struct {
	Id        int64
	DeliverId int64
	GiftId    int64
	GiftName  string
	OpenId    string
	Code      string
	Date      int64
	Method    int64
	CodeId	  int64 `orm:"-"`
	Name     string`orm:"-"`
	Email    string`orm:"-"`
	Phone    string`orm:"-"`
	Address  string`orm:"-"`
	AddressDate     int64`orm:"-"`

}


type LuckybagLottoryRedpack struct {
	Id        int64
	RedPackId string
	DeliverId int64
	Fee       int64
	GiftId    int64
	OpenId    string
	ErrMsg    string
	Code      string
	GiftName  string `orm:"-"`
}

type LuckybagLottoryRedpackDisplay struct {
	Id        int64
	RedPackId string
	DeliverId int64
	Fee       float64
	GiftId    int64
	OpenId    string
	ErrMsg    string
	Code      string
	GiftName  string `orm:"-"`
}

type LuckybagLottory struct {
	Id          int64
	Qx          string
	DeliverId   int64
	BatchNo     int64
	Url         string
	Method      int64
	CreatedDate int64
	UsedDate    int64
	CodeId	  	int64 `orm:"-"`
	GiftName   string`orm:"-"`
	Code		string`orm:"-"`
	Date 		int64 `orm:"-"`
}

type LuckybagLottoryOutput struct {
	Id          int64
	Qx          string
	DeliverId   int64
	BatchNo     int64
	Url         string
	Method      int64
	CreatedDate int64
	UsedDate    int64
}

type LotteryGiftsLogs struct {
	Id			int64
	DeliverId	int64
	GiftId		int64
	GiftName	string
	Quantity	int64
	Date		int64

}
