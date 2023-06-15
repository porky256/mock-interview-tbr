package models

import (
	"time"
)

type Interview struct {
	Id int64 `json:"id,omitempty"`

	Match *Match `json:"match,omitempty"`
	// Interview Status
	Status string `json:"status,omitempty"`

	Date time.Time `json:"date,omitempty"`
}
