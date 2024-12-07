package dao

import (
	"ContentSystem/internal/model"
	"fmt"
	"log"

	"gorm.io/gorm"
)

type ContentDao struct {
	db *gorm.DB
}

func NewContentDao(db *gorm.DB) *ContentDao {
	return &ContentDao{db: db}
}

// 右键函数名 Generate unit test for function
func (c *ContentDao) Create(detail model.ContentDetail) error {
	if err := c.db.Create(&detail).Error; err != nil {
		log.Printf("content create error = %v", err)
		return err
	}
	return nil
}

func (c *ContentDao) IsExist(contentID int) (bool, error) {
	var content model.ContentDetail
	err := c.db.Where("id=?", contentID).First(&content).Error
	if err == gorm.ErrRecordNotFound {
		return false, nil
	}
	if err != nil {
		fmt.Printf("ContentDao isExist = %v\n", err)
		return false, err
	}
	return true, nil
}

func (c *ContentDao) Update(id int, detail model.ContentDetail) error {
	if err := c.db.Where("id=?", id).Updates(&detail).Error; err != nil {
		log.Printf("contentDao update = %v", err)
		return err
	}
	return nil
}

func(c *ContentDao)UpdateByID(id int, column string, value interface{})error{
	if err := c.db.Where("id = ?", id).UpdateColumn(column, value).Error; err!= nil{
		log.Printf("contentDao updatebyID = %v", err)
		return err
	}
	return nil
}

func (c *ContentDao) Delete(id int) error {
	if err := c.db.Where("id=?", id).Delete(&model.ContentDetail{}).Error; err != nil {
		log.Printf("contentDao delete = %v", err)
		return err
	}
	return nil
}

type FindParams struct {
	ID       int
	Page     int
	PageSize int
}

func (c *ContentDao) Find(params *FindParams) ([]*model.ContentDetail, int64, error) {
	query := c.db.Model(&model.ContentDetail{})
	if params.ID != 0 {
		query = query.Where("id = ?", params.ID)
	}
	var total int64
	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}
	var page, pageSize = 1, 10
	if params.Page > 0 {
		page = params.Page
	}
	if params.PageSize > 0 {
		pageSize = params.PageSize
	}
	offset := (page - 1) * pageSize
	var data []*model.ContentDetail
	if err := query.Offset(offset).Limit(pageSize).Find(&data).Error; err != nil {
		return nil, 0, err
	}
	return data, total, nil

}

func (c *ContentDao) First(id int) (*model.ContentDetail, error) {
	var detail model.ContentDetail
	if err := c.db.Where("id = ?", id).First(&detail).Error; err != nil {
		log.Printf("content first err = %v", err)
	}
	return &detail, nil
}
