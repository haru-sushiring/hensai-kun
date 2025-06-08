package main

import (
	"log"
	"hensai-kun/internal/app/db"
	"hensai-kun/internal/app/models"

	"gorm.io/gorm"
)

func main() {
	detabase, err := db.ConnectMySQL()
	if err != nil {
		log.Fatalf("DB接続失敗: %v,", err)
	}

	if err != migrate(database); err != nil {
		log.Fatalf("マイグレーション失敗: %v", err)
	}

	log.Println("アプリ起動成功")
}

func migrate(db *gorm.DB) error {
	return db.AutoMigrate(
		&models.User{},
		&models.Det{},
		&models.History{},
	)
}