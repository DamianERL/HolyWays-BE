package models

import "time"

type User struct {
	ID           int           `json:"id" gorm:"primary_key:auto_increment"`
	Name         string        `json:"name" gorm:"type: varchar(255)"`
	Email        string        `json:"email" gorm:"type: varchar(255)"`
	Password     string        `json:"password" gorm:"type: varchar(255)"`
	Phone        string        `json:"phone" gorm:"type: varchar(255)"`
	Image        string        `json:"image" gorm:"type:varchar(255)"`
	Transactions []Transaction `json:"transactions"`
	Funds        []Fund        `json:"fund"`
	CreatedAt    time.Time     `json:"-"`
	UpdatedAt    time.Time     `json:"-"`
}

type UserResponse struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Phone    string `json:"phone"`
	Image    string `json:"image"`
}

func (UserResponse) TableName() string {
	return "Users"
}
