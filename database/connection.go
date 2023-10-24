package database

import (
	"log"

	"github.com/casikenegro/golang-gin-api/entities"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

type ConfigDB struct {
	Host     string
	User     string
	Password string
	DBName   string
	Port     string
}

func DBConnection(configDB *ConfigDB) {
	var err error
	var DSN = "host=" + configDB.Host + " user=" + configDB.User + " password=" + configDB.Password + " dbname=" + configDB.DBName + " port=" + configDB.Port
	DB, err = gorm.Open(postgres.Open(DSN), &gorm.Config{})
	DB.AutoMigrate(&entities.Product{}, &entities.ProductImage{})
	if err != nil {
		panic("failed to connect db")
	}
	if err != nil {
		log.Fatal(err)
		panic(err)
	} else {
		log.Println("Database connection successful")
	}
}
