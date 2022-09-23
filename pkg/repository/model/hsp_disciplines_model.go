package model

import (
	"gorm.io/gorm"
	"time"
)

type HspDisciplines struct {
	ID        uint   `gorm:"primaryKey" json:"id"`
	Name      string `json:"name"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
