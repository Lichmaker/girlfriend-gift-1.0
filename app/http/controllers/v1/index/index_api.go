package index

import (
	"lichmaker/girlfriend-gift-1/app/models/schedule"
	"net/http"

	"github.com/gin-gonic/gin"
)

type IndexApi struct {
}

func (i *IndexApi) Get (c *gin.Context) {
	var n schedule.Schedule
	n.Md5 = "abc"
	n.Date = "2021-01-01"
	n.Path = "test"
	n.Create()

	c.JSON(http.StatusOK, gin.H{
		"测试": "成功",
	})
}