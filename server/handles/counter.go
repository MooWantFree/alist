package handles

import (
	"github.com/alist-org/alist/v3/internal/db"
	"github.com/alist-org/alist/v3/internal/model"
	"github.com/alist-org/alist/v3/server/common"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

func GetCounter(c *gin.Context) {
	var counterPage model.CounterPage
	if err := c.ShouldBindJSON(&counterPage); err != nil {
		common.ErrorResp(c, err, 400)
		return
	}
	counterPage.Validate()
	log.Debugf("%+v", counterPage)
	downloadCounterColumn, err := db.GetDownloadColumn(counterPage.CurrentPage, counterPage.PageSize, counterPage.SortKey, counterPage.Reverse, counterPage.FileName, counterPage.IPAddress, counterPage.StatusCode)
	if err != nil {
		common.ErrorResp(c, err, 500)
		return
	}
	common.SuccessResp(c, common.PageResp{
		Content: downloadCounterColumn,
		Total:   int64(len(downloadCounterColumn)),
	})
}
