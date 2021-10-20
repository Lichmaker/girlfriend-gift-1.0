package pool

import (
	"lichmaker/girlfriend-gift-1/app/models"
)

type Pool struct {
	models.BaseModel

	Md5  string `gorm:"column:md5;type:varchar(64);unique"`
	ModDate string `gorm:"column:mod_date;type:date;index"`
	LocalPath string `gorm:"column:local_path;type:varchar(128);"`
	OssPath string `gorm:"column:oss_path;type:varchar(128);"`
}

type Tabler interface {
	TableName() string
}

func (Pool) TableName() string {
	return "pool"
}
