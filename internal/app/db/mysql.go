package db

import (
	"fmt"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func ConnectMySQL() (*gorm.DB, error) {
	user := getEnv("DB_USER", "user")
	pass := getEnv("DB_PASSWORD", "password")
	host := getEnv("DB_HOST", "localhost")
	port := getEnv("DB_PORT", "3306")
	name := getEnv("DB_NAME", "app_db")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parsetTime=True&loc=Local",
		user, pass, host, port, name)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	return db, nil
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}