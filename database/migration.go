package database

import (
	"fmt"
	"holyways/models"
	"holyways/pkg/mysql"
)

func RunMigration() {
	err := mysql.DB.AutoMigrate(
		&models.User{},
		&models.Fund{},
		&models.Donation{},
		&models.Transaction{},
	
	)
	if err != nil {
		fmt.Println(err)
		panic("Migration Failed")
	}

	fmt.Println(("Migration success"))
}
