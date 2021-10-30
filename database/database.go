package database

import (
	"astroauth-api/models"
	"context"
	"fmt"
	"log"
	"os"

	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB
var DBB *pgxpool.Pool

var err error

func InitialMigration() {
	type DBConfig struct {
		Host     string `env:"DBHOST"`
		User     string `env:"DBUser"`
		Password string `env:"DBPASSWORD"`
		DBName   string `env:"DBDBNAME"`
		Port     string `env:"DBPORT"`
	}

	//Get enviroment variables
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	fmt.Println(os.Getenv("DBHOST"))
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Shanghai",
		os.Getenv("DBHOST"),
		os.Getenv("DBUSER"),
		os.Getenv("DBPASSWORD"),
		os.Getenv("DBNAME"),
		os.Getenv("DBPORT"))

	// dsn1 := "host=localhost user=postgres password=1234 dbname=postgres port=5432 sslmode=disable TimeZone=Asia/Shanghai"

	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("db error")
	}

	// urlExample := "postgres://username:password@localhost:5432/database_name"
	DBB, err = pgxpool.Connect(context.Background(), fmt.Sprintf("postgres://%s:%s@%s:%s/%s",
		os.Getenv("DBUSER"),
		os.Getenv("DBPASSWORD"),
		os.Getenv("DBHOST"),
		os.Getenv("DBPORT"),
		os.Getenv("DBNAME")))
	if err != nil {
		panic("db error")
	}

	DB.AutoMigrate(&models.SiteUser{}, &models.App{}, &models.AppUser{}, &models.License{})
	InitializeStore()
}
