package db

import (
	"fmt"
	"github.com/alist-org/alist/v3/internal/model"
	"github.com/pkg/errors"
)

func InsertDownloadColumn(c *model.Counter) error {
	return errors.WithStack(db.Create(c).Error)
}

func sortMethod(reverse bool, orderBy string) string {
	const (
		asc  = "ASC"
		desc = "DESC"
	)
	orderDirection := asc
	if reverse {
		orderDirection = desc
	}
	orderClause := fmt.Sprintf("%s %s", orderBy, orderDirection)
	return orderClause
}
func GetDownloadColumn(currentPage int, pageSize int, sortKey string, reverse bool, fileName, IPAddress string, httpStatusCode int) ([]model.Counter, int64, error) {
	var counts []model.Counter
	orderClause := sortMethod(reverse, sortKey)
	query := db.Model(&model.Counter{})
	var totalItems int64
	err := db.Model(&model.Counter{}).Count(&totalItems).Error
	if err != nil {
		return nil, 0, err
	}
	if fileName != "" {
		query = query.Where(columnName("file_name")+" LIKE ?", "%"+fileName+"%")
	}
	if IPAddress != "" {
		query = query.Where(columnName("ip_address")+" LIKE ?", "%"+IPAddress+"%")
	}
	if httpStatusCode != 0 {
		query = query.Where(columnName("status_code")+" = ?", httpStatusCode)
	}
	err = query.Order(orderClause).Limit(pageSize).Offset((currentPage - 1) * pageSize).Find(&counts).Error
	return counts, totalItems, err
}
