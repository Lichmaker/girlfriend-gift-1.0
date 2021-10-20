package index

import (
	"lichmaker/girlfriend-gift-1/pkg/picker"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type IndexApi struct {
}

func (i *IndexApi) Get(c *gin.Context) {
	url := picker.Do(time.Now().Format("2006-01-02"))

	c.JSON(http.StatusOK, gin.H{
		"url": url,
	})
}
