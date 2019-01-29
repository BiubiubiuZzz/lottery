package models

import "github.com/astaxie/beego/orm"

type LotteryUser struct {
	Id int		//主键
	Username 	string `orm:"size(15)"`   //用户名
	Password 	string `orm:"size(32)"`  //密码
	DeliverId	string `orm:size(32)`
}

func (this *LotteryUser) TableName() string {
	return  "LotteryUser"
}

func (user *LotteryUser) Insert() error {
	if _, err := orm.NewOrm().Insert(user); err != nil {
		return err
	}
	return nil
}

func (user *LotteryUser) Read(fields ...string) error {
	orm := orm.NewOrm()

	if err := orm.Read(user, fields...); err != nil {
		return err
	}
	return nil
}

func (user *LotteryUser) Update(fields ...string) error {
	if _, err := orm.NewOrm().Update(user, fields...); err != nil {
		return err
	}
	return nil
}

func (user *LotteryUser) Delete() error {
	if _, err := orm.NewOrm().Delete(user); err != nil {
		return err
	}
	return nil
}
