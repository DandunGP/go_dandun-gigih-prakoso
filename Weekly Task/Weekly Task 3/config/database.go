package config

import (
	"fmt"
	"log"
	"os"
	"weekly3/entities"

	"github.com/joho/godotenv"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var (
	DB *gorm.DB
)

func init() {
	InitDB()
	InitialMigration()
}

type Config struct {
	DB_Username string
	DB_Password string
	DB_Port     string
	DB_Host     string
	DB_Name     string
}

func InitDB() *gorm.DB {

	errLoad := godotenv.Load(".env")
	if errLoad != nil {
		log.Fatalf("Error loading .env file")
	}

	config := Config{
		DB_Username: os.Getenv("DB_Username"),
		DB_Password: os.Getenv("DB_Password"),
		DB_Port:     os.Getenv("DB_Port"),
		DB_Host:     os.Getenv("DB_Host"),
		DB_Name:     os.Getenv("DB_Name"),
	}

	connectionString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		config.DB_Username,
		config.DB_Password,
		config.DB_Host,
		config.DB_Port,
		config.DB_Name,
	)

	var err error
	DB, err = gorm.Open("mysql", connectionString)
	if err != nil {
		panic(err)
	}

	return DB
}

func InitialMigration() {
	DB.AutoMigrate(&entities.User{})
	DB.AutoMigrate(&entities.Category{})
	DB.AutoMigrate(&entities.Blog{})
}
