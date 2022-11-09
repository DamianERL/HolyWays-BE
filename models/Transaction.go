package models

import "time"

type Transaction struct {
	ID         int       `json:"id" gorm:"primary_key:auto_increment" `
	Total      int       `json:"total"`
	UserID     int       `json:"-"`
	User       string    `json:"user"`
	Status     string    `json:"status"`
	Donate     int       `json:"donate"`
	DonationID int       `json:"-" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"  `
	Donation   Donation  `json:"donation"`
	CreatedAt  time.Time `json:"create_at"`
	UpdatedAt  time.Time `json:"-"`
}
