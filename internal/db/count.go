package db

import (
	"github.com/alist-org/alist/v3/internal/model"
	"github.com/pkg/errors"
)

func InsertDownloadColumn(c *model.Count) error {
	return errors.WithStack(db.Create(c).Error)
}
