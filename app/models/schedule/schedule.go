package schedule

import (
	"lichmaker/girlfriend-gift-1/app/models"
)

type Schedule struct {
	models.BaseModel

	Md5  string `gorm:"column:md5;type:varchar(64);index"`
	Date string `gorm:"column:date;type:date;index"`
	Path string `gorm:"column:path;type:varchar(128);"`
}

type Tabler interface {
	TableName() string
}

// TableName 会将 User 的表名重写为 `profiles`
func (Schedule) TableName() string {
	return "schedule"
}
