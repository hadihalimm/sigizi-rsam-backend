package config

import (
	"fmt"
	"log"
	"net/url"
	"os"

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

	DB_HOST := os.Getenv("DB_HOST")
	DB_PORT := os.Getenv("DB_PORT")
	DB_NAME := os.Getenv("DB_NAME")
	DB_USER := os.Getenv("DB_USER")
	DB_PASSWORD := os.Getenv("DB_PASSWORD")
	loc := url.QueryEscape("Asia/Jakarta")

	// dsnLocal := "hadi:hadi@tcp(127.0.0.1:3306)/sigizi_rsam?charset=utf8mb4&parseTime=True&loc=Local"
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=%s",
		DB_USER, DB_PASSWORD, DB_HOST, DB_PORT, DB_NAME, loc)
	fmt.Println(dsn)
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
