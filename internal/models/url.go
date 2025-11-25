package models

import "gorm.io/gorm"

type URL struct {
	gorm.Model
	LongUrl   string `gorm:"uniqueIndex;size:500"  json:"long_url"`
	ShortCode string `gorm:"uniqueIndex;size:500" json:"short_code"`
	Clicks    uint   `json:"clicks"`
}
