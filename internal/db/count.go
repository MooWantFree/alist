package db

import (
	"fmt"
	"github.com/alist-org/alist/v3/internal/model"
	"github.com/pkg/errors"
)

func InsertDownloadColumn(c *model.Counter) error {
	return errors.WithStack(db.Create(c).Error)
}

const (
	asc  = "ASC"
	desc = "DESC"
)

func GetDownloadColumn(index int, pagination int, orderBy string, reverse bool) ([]model.Counter, error) {
	var counts []model.Counter
	orderDirection := asc
	if reverse {
		orderDirection = desc
	}
	orderClause := fmt.Sprintf("%s %s", orderBy, orderDirection)
	if err := db.Order(orderClause).Limit(pagination).Offset((index - 1) * pagination).Find(&counts).Error; err != nil {
		return nil, err
	}
	return counts, nil
}
