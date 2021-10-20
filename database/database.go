package database

import (
	"astroauth-api/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

var err error

func InitialMigration() {
	dsn := "host=localhost user=postgres password=1234 dbname=db3 port=5432 sslmode=disable TimeZone=Asia/Shanghai"

	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("db error")
	}
	DB.AutoMigrate(&models.SiteUser{}, &models.App{}, &models.AppUser{}, &models.License{})
	InitializeStore()
}
