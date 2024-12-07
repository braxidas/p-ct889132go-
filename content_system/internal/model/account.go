package model

import "time"

////可以自动转换类型(不推荐)
type Account struct {
	ID       int64     `gorm:"column:id;primary_key"`
	UserID   string    `gorm:"column:user_id`
	Password string    `gorm:"column:password"`
	Nickname string    `gorm:"column:nickname"`
	Ct       time.Time `gorm:"column:created_at"`
	Ut       time.Time `gorm:"column:updated_at"`
}

//实现gorm的接口，可以动态确定表名
func(a Account)TableName()string{
	//指定 库名.表名
	return "cms_account.account"
}
