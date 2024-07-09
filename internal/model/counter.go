package model

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"time"
)

type JSONHeader map[string][]string

func (h *JSONHeader) Value() (driver.Value, error) {
	if h == nil {
		return nil, nil
	}
	return json.Marshal(h)
}

func (h *JSONHeader) Scan(value interface{}) error {
	bytes, ok := value.([]byte)
	if !ok {
		return fmt.Errorf("failed to unmarshal JSONHeader value: %v", value)
	}
	return json.Unmarshal(bytes, h)
}

type Counter struct {
	ID        uint      `json:"id" gorm:"primaryKey"` // unique key
	FileName  string    `json:"file_name"`            // file name
	FilePath  string    `json:"file_path"`            // file path
	Time      time.Time `json:"time"`                 // download time
	IPAddress string    `json:"ip_address"`           // request IP
}
type CounterPage struct {
	CurrentPage int    `json:"current_page" gorm:"primaryKey"`
	PageSize    int    `json:"page_size" gorm:"not null"`
	SortKey     string `json:"sort_key" gorm:"not null"`
	Reverse     bool   `json:"reverse" gorm:"not null"`
	FileName    string `json:"file_name"`
	IPAddress   string `json:"ip_address"`
}

func (c *CounterPage) Validate() {
	if c.CurrentPage < 1 {
		c.CurrentPage = 1
	}
	if c.PageSize < 1 {
		c.PageSize = MaxInt
	}
	if c.SortKey == "" {
		c.SortKey = "download_time"
	}
}
