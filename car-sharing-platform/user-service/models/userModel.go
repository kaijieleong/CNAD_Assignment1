package models

// User represents the schema of the Users table
type User struct {
    ID              uint    `gorm:"primaryKey;column:ID"`
    Name            string  `gorm:"column:Name"`
    Email           string  `gorm:"column:Email"`
    Phone           string  `gorm:"column:Phone"`
    Password        string  `gorm:"column:Password"`
    MembershipTier  string  `gorm:"column:MembershipTier"`
    TotalSpending   float64 `gorm:"column:TotalSpending"`
}
