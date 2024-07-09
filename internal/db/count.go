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
func GetDownloadColumn(currentPage int, pageSize int, sortKey string, reverse bool, fileName, IPAddress string) ([]model.Counter, int, error) {
	var counts []model.Counter
	orderClause := sortMethod(reverse, sortKey)
	query := db.Model(&model.Counter{})
	var tempTotalItems int64
	err := db.Model(&model.Counter{}).Count(&tempTotalItems).Error
	totalItems := int(tempTotalItems)
	totalItems = func(a, b int) int {
		if a%b == 0 {
			return a / b
		} else {
			return a/b + 1
		}
	}(totalItems, pageSize)
	if err != nil {
		return nil, 0, err
	}
	if fileName != "" {
		query = query.Where(columnName("file_name")+" LIKE ?", "%"+fileName+"%")
	}
	if IPAddress != "" {
		query = query.Where(columnName("ip_address")+" LIKE ?", "%"+IPAddress+"%")
	}
	err = query.Order(orderClause).Limit(pageSize).Offset((currentPage - 1) * pageSize).Find(&counts).Error
	return counts, totalItems, err
}
