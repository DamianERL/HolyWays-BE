package models

import "time"

type Fund struct {
	ID        int          `json:"id" gorm:"primary_key:auto_increment"`
	Name      string       `json:"name" gorm:"type:varchar(255)"`
	Image     string       `json:"image" gorm:"type:varchar(255)"`
	Desc      string       `json:"desc" gorm:"type:varchar(255)"`
	Goals     int          `json:"goals" gorm:"type:int"`
	Donated   int          `json:"donated" gorm:"type:int"`
	UserID    int          `json:"-" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	User      UserResponse `json:"user"`
	CreatedAt time.Time    `json:"create_at"`
	Update_at time.Time    `json:"-"`
}

type FundResponse struct {
	ID          int          `json:"id"`
	Name        string       `json:"name"`
	Image       string       `json:"image"`
	Description string       `json:"description"`
	Goals       int          `json:"goals"`
	Donated     int          `json:"donated"`
	UserID      int          `json:"-"`
	User        UserResponse `json:"user"`
}

func (FundResponse) TableName() string {
	return "Funds"
}
