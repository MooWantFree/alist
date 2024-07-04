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
	insert(c)

}

// TODO: implement
// path maybe contains # ? etc.
func parsePath(path string) string {
	return utils.FixAndCleanPath(path)
}
func insert(c *gin.Context) {
	decodedDownloadReqURL, err := url.QueryUnescape(c.Request.URL.String())
	basePath, _ := splitPathAndQuery(decodedDownloadReqURL)
	fileName := filepath.Base(basePath)
	err = db.InsertDownloadColumn(&model.Count{
		FileName:       fileName,
		FilePath:       basePath,
		Operation:      c.Request.Method,
		DownloadTime:   time.Now(),
		RequestIP:      c.ClientIP(),
		HttpStatusCode: c.Writer.Status(),
	})
	if err != nil {
		log.Fatalf("failed to create user: %+v", err)
	}
}
func splitPathAndQuery(path string) (string, string) {
	parts := strings.SplitN(path, "?", 2)
	if len(parts) == 2 {
		return parts[0], parts[1]
	}
	return parts[0], ""
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
