package models

type Vehicle struct {
	ID           uint   `gorm:"primaryKey"`
	LicensePlate string `gorm:"unique"`
	Model        string
	ChargeLevel  int
	Status       string
	Location     string
}
