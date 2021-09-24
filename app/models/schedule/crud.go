package schedule

import (
	"lichmaker/girlfriend-gift-1/pkg/model"
)

func (schedule *Schedule) Create() (err error) {
	result := model.DB.Create(&schedule)
	if err = result.Error; err != nil {
		return err
	}

	return nil
}
