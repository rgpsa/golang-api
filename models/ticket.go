package models

import (
	"time"

	"github.com/lib/pq"
)

type Ticket struct {
	ID          uint           `json:"id" gorm:"primary_key;auto_increment"`
	Type        string         `json:"type"`
	Subject     string         `json:"subject"`
	Description string         `json:"description"`
	Priority    string         `json:"priority"`
	Status      string         `json:"status"`
	Tags        pq.StringArray `json:"tags" gorm:"type:varchar(255)[]"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
}
