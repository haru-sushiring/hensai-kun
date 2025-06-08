package models

type Det struct {
	ID uint `gorm:"primaryKey"`
	MonthlyPayment int `gorm:"not null"`
	Principal int `gorm:"not null"`
	Balance int `gorm:"not null"`
	InterestPaid int `gorm:"not null"`
	RemainingInterest int `gorm:"not null"`
	AnnualInterestRate int `gorm:"not null"`
	NumberOfPayments int `gorm:"not null"`
	NumberOfRemainingPayments int `gorm:"not null"`
	CompletionFlag int `gorm:"not null"`
	PaymentType int `gorm:"not null"`
}