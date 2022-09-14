package gormMigrations

import "gorm.io/gorm"

type CoreUser struct {
	Name       string
	Username   string
	Sex        string
	Email      string
	Password   string
	Status     int
	Permission []CorePermission `gorm:"many2many:core_users_permissions"`
	Role       []CoreRole       `gorm:"many2many:core_users_roles"`
	gorm.Model
}
