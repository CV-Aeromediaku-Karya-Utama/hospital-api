package model

import (
	"gorm.io/gorm"
	"time"
)

type HspDoctor struct {
	ID               uint           `gorm:"primaryKey" json:"id"`
	Name             string         `json:"name"`
	HspDisciplinesID uint           `json:"hsp_disciplines_id"`
	HspDisciplines   HspDisciplines `json:"hsp_disciplines"`
	HspUnit          []HspUnit      `gorm:"many2many:hsp_doctors_units" json:"hsp_unit"`
	CreatedAt        time.Time
	UpdatedAt        time.Time
	DeletedAt        gorm.DeletedAt `gorm:"index"`
}
