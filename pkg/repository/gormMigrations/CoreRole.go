package gormMigrations

type CoreRole struct {
	ID         uint `gorm:"primaryKey"`
	Name       string
	Permission []CorePermission `gorm:"many2many:core_roles_permissions"`
}
