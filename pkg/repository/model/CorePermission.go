package model

type CorePermission struct {
	ID   uint `gorm:"primaryKey"`
	Name string
}
