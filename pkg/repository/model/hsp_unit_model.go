package model

import (
	"gorm.io/gorm"
	"time"
)

type HspUnit struct {
	ID        uint   `gorm:"primaryKey" json:"id"`
	UnitCode  string `json:"unit_code"`
	UnitName  string `json:"unit_name"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
