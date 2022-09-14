package model

type CoreRole struct {
	ID         uint `gorm:"primaryKey"`
	Name       string
	Permission []CorePermission `gorm:"many2many:core_roles_permissions"`
}
