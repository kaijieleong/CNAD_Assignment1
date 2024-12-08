package models

import "time"

type Booking struct {
    ID         uint      `json:"id"`
    UserID     uint      `json:"user_id"`  
    VehicleID  uint      `json:"vehicle_id"`
    StartTime  time.Time `json:"start_time"`
    EndTime    time.Time `json:"end_time"`
    TotalPrice float64   `json:"total_price"`
}
