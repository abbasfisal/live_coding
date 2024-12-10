package entity

import (
	"github.com/google/uuid"
)

type Address struct {
	ID      uuid.UUID `json:"id" gorm:"type:char(36);primaryKey"`
	UserID  uuid.UUID `json:"user_id" gorm:"type:uuid"`
	Street  string    `json:"street"`
	City    string    `json:"city"`
	State   string    `json:"state"`
	ZipCode string    `json:"zip_code"`
	Country string    `json:"country"`
	User    User      `json:"user"`
}

