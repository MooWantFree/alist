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
func GetDownloadColumn(index int, pagination int, orderBy string, reverse bool, fileName, IPAddress string, httpStatusCode int) ([]model.Counter, error) {
	var counts []model.Counter
	orderClause := sortMethod(reverse, orderBy)
	query := db.Model(&model.Counter{})
	if fileName != "" {
		query = query.Where("file_name LIKE ?", "%"+fileName+"%")
	}
	if IPAddress != "" {
		query = query.Where("request_ip LIKE ?", "%"+IPAddress+"%")
	}
	if httpStatusCode != 0 {
		query = query.Where("http_status_code = ?", httpStatusCode)
	}
	err := query.Order(orderClause).Limit(pagination).Offset((index - 1) * pagination).Find(&counts).Error
	return counts, err
}
