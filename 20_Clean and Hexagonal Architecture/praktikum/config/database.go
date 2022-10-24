package config

import (
	"belajar-go-echo/entities"
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// func ConnectDB() (*gorm.DB, error) {
// 	return gorm.Open(sqlite.Open("app.db"), &gorm.Config{})
// }

// func MigrateDB(db *gorm.DB) error {
// 	return db.AutoMigrate(
// 		entities.User{},
// 	)
// }

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

	config := Config{
		DB_Username: "root",
		DB_Password: "",
		DB_Port:     "3306",
		DB_Host:     "localhost",
		DB_Name:     "crud_clean",
	}

	connectionString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		config.DB_Username,
		config.DB_Password,
		config.DB_Host,
		config.DB_Port,
		config.DB_Name,
	)

	var err error
	DB, err = gorm.Open(mysql.Open(connectionString), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	return DB
}

func InitialMigration() {
	DB.AutoMigrate(&entities.User{})
}
