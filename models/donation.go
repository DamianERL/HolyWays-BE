package models

type Donation struct {
	ID     int          `json:"id" gorm:"primary_key:auto_increment" `
	Donate int          `json:"donate" gorm:"type:int"`
	UserID int          `json:"-"`
	User   UserResponse `json:"user"`
	FundID int          `json:"-" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"  `
	Fund   FundResponse `json:"fund"  `
	Status string       `json:"status"`
}
