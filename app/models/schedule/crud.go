package schedule

import (
	"lichmaker/girlfriend-gift-1/pkg/model"

	"gorm.io/gorm"
)

func (schedule *Schedule) Create() (err error) {
	result := model.DB.Create(&schedule)
	if err = result.Error; err != nil {
		return err
	}

	return nil
}

func Exists(md5 string) bool {
	var r Schedule
	result := model.DB.Where("md5 = ?", md5).First(&r)
	if result.Error == gorm.ErrRecordNotFound {
		return false
	} else {
		return true
	}
}

func GetRandom(date string) (Schedule, error) {
	var r Schedule
	result := model.DB.Where("date = ?", date).Order("RAND()").First(&r)
	if result.Error != nil {
		return r, result.Error
	} else {
		return r, nil
	}
}