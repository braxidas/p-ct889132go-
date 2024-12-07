package model

//3-17
type ContentDetail struct {
	ID             int    `gorm:"column:id;primary_key"`
	Title          string `gorm:"column:title"`
	Description    string `gorm:"description"`
	VideoURL       string `gorm:"video_url"`
	Category       string `gorm:"category"`
	ApprovalStatus int    `gorm:"approval_status"`
	Thumbnail      string `gorm:"thumbnail"`
	Format         string `gorm:"format"`
}

func (c ContentDetail) TableName() string {
	return "cms_content.t_content_details"
}
