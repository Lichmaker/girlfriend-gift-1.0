package picker

import (
	"lichmaker/girlfriend-gift-1/app/models/schedule"
	"lichmaker/girlfriend-gift-1/pkg/oss_helper"
)

func Do(date string) (string) {
	r, err := schedule.GetRandom(date)
	if err != nil {
		return ""
	}
	
	return oss_helper.GetUrl(r.Path, int64(600))
}