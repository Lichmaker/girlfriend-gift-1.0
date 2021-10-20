package pool

import (
	"lichmaker/girlfriend-gift-1/pkg/model"
)

func QueryMd5NotIn(md5Array []string) []string {
	var query []Pool
	queryMd5 := make(map[string]interface{})
	var notInMd5 []string
	model.DB.Where("md5 IN ?", md5Array).Find(&query)
	// fmt.Println(len(query))
	for _, row := range query {
		queryMd5[row.Md5] = 1
	}
	for _, _md5 := range md5Array {
		if queryMd5[_md5] == nil {
			notInMd5 = append(notInMd5, _md5)
		}
	}
	return notInMd5
}

func Create(md5 string, localPath string, ossPath string, modDate string) error {
	var poolModel Pool
	var err error

	poolModel.Md5 = md5
	poolModel.ModDate = modDate
	poolModel.LocalPath = localPath
	poolModel.OssPath = ossPath

	result := model.DB.Create(&poolModel)
	if err = result.Error; err != nil {
		return err
	}
	return nil
}
