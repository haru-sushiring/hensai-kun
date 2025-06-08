package models

type Det struct {
	ID uint `gorm:"primaryKey"`
	DetId int `gorm:"not null"`
}

