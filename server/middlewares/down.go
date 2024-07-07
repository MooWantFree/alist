package middlewares

import (
	"github.com/alist-org/alist/v3/internal/conf"
	"github.com/alist-org/alist/v3/internal/db"
	"github.com/alist-org/alist/v3/internal/errs"
	"github.com/alist-org/alist/v3/internal/model"
	"github.com/alist-org/alist/v3/internal/op"
	"github.com/alist-org/alist/v3/internal/setting"
	"github.com/alist-org/alist/v3/internal/sign"
	"github.com/alist-org/alist/v3/pkg/utils"
	"github.com/alist-org/alist/v3/server/common"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
	"net/url"
	"path/filepath"
	"strings"
	"time"
)

func Down(c *gin.Context) {
	rawPath := parsePath(c.Param("path"))
	c.Set("path", rawPath)
	meta, err := op.GetNearestMeta(rawPath)
	if err != nil {
		if !errors.Is(errors.Cause(err), errs.MetaNotFound) {
			common.ErrorResp(c, err, 500, true)
			return
		}
	}
	c.Set("meta", meta)
	// verify sign
	if needSign(meta, rawPath) {
		s := c.Query("sign")
		err = sign.Verify(rawPath, strings.TrimSuffix(s, "/"))
		if err != nil {
			common.ErrorResp(c, err, 401)
			c.Abort()
			return
		}
	}
	c.Next()
	insertDownloadCounter(c)

}

// TODO: implement
// path maybe contains # ? etc.
func parsePath(path string) string {
	return utils.FixAndCleanPath(path)
}
func insertDownloadCounter(c *gin.Context) {
	decodedDownloadReqURL, err := url.QueryUnescape(c.Request.URL.String())
	if err != nil {
		log.Fatalf("failed to decode URL: %+v", err)
		return
	}
	splitPathAndQuery := func(path string) (string, string) {
		parts := strings.SplitN(path, "?", 2)
		if len(parts) == 2 {
			return parts[0], parts[1]
		}
		return parts[0], ""
	}
	basePath, _ := splitPathAndQuery(decodedDownloadReqURL)
	fileName := filepath.Base(basePath)
	err = db.InsertDownloadColumn(&model.Counter{
		FileName:   fileName,
		FilePath:   basePath,
		Time:       time.Now(),
		IPAddress:  c.ClientIP(),
		StatusCode: c.Writer.Status(),
	})
	if err != nil {
		log.Fatalf("failed to insert column: %+v", err)
	}
	log.Info("insert 1 column to download counter")
}

func needSign(meta *model.Meta, path string) bool {
	if setting.GetBool(conf.SignAll) {
		return true
	}
	if common.IsStorageSignEnabled(path) {
		return true
	}
	if meta == nil || meta.Password == "" {
		return false
	}
	if !meta.PSub && path != meta.Path {
		return false
	}
	return true
}
