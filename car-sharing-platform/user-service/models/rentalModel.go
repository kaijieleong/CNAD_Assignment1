package models

type RentalHistory struct {
	ID          uint    `gorm:"primaryKey"`
	UserID      uint    `gorm:"column:UserID"`
	VehicleID   uint    `gorm:"column:VehicleID"`
	RentalDate  string  `gorm:"column:RentalDate"`
	AmountSpent float64 `gorm:"column:AmountSpent"`
}

// TableName overrides the default table name
func (RentalHistory) TableName() string {
	return "RentalHistory"
}
