package seed

import (
	"errors"
	"gorm.io/gorm"
	"hospital-api/pkg/repository/model"
)

func SeedPermission(gormDB *gorm.DB) {
	if gormDB.Migrator().HasTable(&model.CorePermission{}) {
		if err := gormDB.First(&model.CorePermission{}).Error; errors.Is(err, gorm.ErrRecordNotFound) {
			permissions := []model.CorePermission{
				{Name: "manage-user"},
				{Name: "manage-role"},
			}
			gormDB.Create(&permissions)
		}
	}
}
