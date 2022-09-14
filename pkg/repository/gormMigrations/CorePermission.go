package gormMigrations

type CorePermission struct {
	ID   uint `gorm:"primaryKey"`
	Name string
}
