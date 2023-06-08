package model

import "time"

// User represents the user for this application
// swagger:model

type User struct {
	ID        string    `gorm:"primaryKey"`
	// The name for this user
	// required: true
	// min lenght: 5
	Name      string    `gorm:"column:name"`
	Email     string    `gorm:"column:email"`
	Photo     string    `gorm:"column:photo"`
	Verified  bool      `gorm:"column:verified"`
	Password  string    `gorm:"column:password"`
	Role      string    `gorm:"column:role"`
	CreatedAt time.Time `gorm:"column:created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at"`
}
