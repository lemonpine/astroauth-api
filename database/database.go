package database

import (
	"astroauth-api/models"
	"context"

	"github.com/jackc/pgx/v4/pgxpool"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB
var DBB *pgxpool.Pool

var err error

func InitialMigration() {
	dsn := "host=localhost user=postgres password=1234 dbname=postgres port=5432 sslmode=disable TimeZone=Asia/Shanghai"

	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("db error")
	}

	// urlExample := "postgres://username:password@localhost:5432/database_name"
	DBB, err = pgxpool.Connect(context.Background(), "postgres://postgres:1234@localhost:5432/postgres")
	if err != nil {
		panic("db error")
	}

	DB.AutoMigrate(&models.SiteUser{}, &models.App{}, &models.AppUser{}, &models.License{})
	InitializeStore()
}
