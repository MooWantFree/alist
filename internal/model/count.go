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

type Count struct {
	ID             uint      `json:"id" gorm:"primaryKey"`      // unique key
	FileName       string    `json:"file_name" gorm:"not null"` // file name
	FilePath       string    `json:"file_path" gorm:"not null"` // file path
	Operation      string    `json:"operation" gorm:"not null"`
	DownloadTime   time.Time `json:"download_time" gorm:"not null"`    // download time
	RequestIP      string    `json:"request_ip" gorm:"not null"`       // request IP
	HttpStatusCode int       `json:"http_status_code" gorm:"not null"` // HTTP status code
}
