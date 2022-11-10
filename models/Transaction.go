package models

import "time"

type Transaction struct {
	ID        int          `json:"id" gorm:"primary_key:auto_increment" `
	Donate    int          `json:"donate"`
	Status    string       `json:"status"`
	UserID    int          `json:"-"`
	User      UserResponse `json:"donatur"`
	FundID    int          `json:"-"`
	Fund      FundResponse `json:"fund"`
	Date      time.Time    `json:"date" gorm:"default:now()" `
	CreatedAt time.Time    `json:"create_at"`
	UpdatedAt time.Time    `json:"-"`
}
