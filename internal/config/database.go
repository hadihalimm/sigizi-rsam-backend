package config

import (
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Database struct {
	Gorm *gorm.DB
}

var dbInstance *Database

func ConnectToDatabase() *Database {
	if dbInstance != nil {
		return dbInstance
	}

	dsn := "hadi:hadi@tcp(127.0.0.1:3306)/sigizi_rsam?charset=utf8mb4&parseTime=True&loc=Local"
	gormDB, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Successfully connected to the database")
	dbInstance := &Database{
		Gorm: gormDB,
	}

	return dbInstance
}
