package models

import (
	"time"

	"gorm.io/gorm"
)

type URL struct {
	gorm.Model
	LongUrl    string     `gorm:"uniqueIndex;size:500"  json:"long_url"`
	ShortCode  string     `gorm:"uniqueIndex;size:500" json:"short_code"`
	Clicks     uint       `json:"clicks"`
	FirstClick *time.Time `json:"first_click,omitempty"`
	LastClick  *time.Time `json:"last_click,omitempty"`
}
