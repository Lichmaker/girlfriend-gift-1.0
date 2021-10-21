package picker

import (
	"lichmaker/girlfriend-gift-1/app/models/schedule"
	"lichmaker/girlfriend-gift-1/pkg/oss_helper"
	"time"
)

func Do(date string) (string, string) {
	r, err := schedule.GetRandom(date)
	if err != nil {
		return "", ""
	}
	dateTime, _ := time.Parse(time.RFC3339, r.ModDate)
	dateString := dateTime.Format("2006-01-02")
	return oss_helper.GetUrl(r.Path, int64(600)), dateString
}
