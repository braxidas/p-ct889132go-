package dao

import (
	"ContentSystem/internal/model"
	"fmt"

	"gorm.io/gorm"
)

type AccountDao struct {
	db *gorm.DB
}

//新建
func NewAccountDao(db *gorm.DB)*AccountDao{
	return &AccountDao{db:db}
}
//判断数据是否存在
func(a *AccountDao)IsExist(userID string)(bool, error){
	var account model.Account
	err := a.db.Where("user_id=?", userID).First(&account).Error
	if err == gorm.ErrRecordNotFound{
		return false, nil
	}
	if err != nil{
		fmt.Printf("isExist error %v\n", err)
		return false, err
	}
	return true, nil
}

func(a *AccountDao)Create(account model.Account)error{
	if err := a.db.Create(&account).Error; err != nil{
		fmt.Printf("accountDao Create = %v\n", err)
		return err
	}
	return nil
}

func(a *AccountDao)FirstByUserID(userID string)(*model.Account,error){
	var account model.Account
	err := a.db.Where("user_id = ?", userID).First(&account).Error
	if err != nil{
		fmt.Printf("FirstByUserID error :%v\n", err)
		return nil, err
	}
	return &account, nil
}