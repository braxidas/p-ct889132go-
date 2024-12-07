package demo

import (
	"fmt"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

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

func main() {
	db := connDB()
	var accouts []Account
	if err := db.Table("account").Find(&accouts).Error; err != nil{
		fmt.Println(err)
	}
	fmt.Println(accouts)
}

func connDB() *gorm.DB {
	mysqlDB, err := gorm.Open(mysql.Open("user:password@tcp..."))
	if err != nil {
		panic(err)
	}
	db, err := mysqlDB.DB()
	if err != nil {
		panic(err)
	}
	db.SetMaxOpenConns(4)
	db.SetMaxIdleConns(2)
	mysqlDB = mysqlDB.Debug()

	return mysqlDB
}
