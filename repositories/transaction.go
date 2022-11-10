package repositories

import (
	"holyways/models"

	"gorm.io/gorm"
)

type TransactionRepository interface {
	FindTransactions(ID int) (models.Transaction, error)
	GetTransaction(ID int) (models.Transaction, error)
	CreateTransaction(transaction models.Transaction) (models.Transaction, error)

	GetOneTransaction(ID string) (models.Transaction, error)
	UpdateTransaction(status string, ID string) error
}

func RepositoryTransaction(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindTransactions(ID int) ([]models.Transaction, error) {
	var transactions []models.Transaction
	err := r.db.Preload("Fund").Preload("Fund.User").Preload("User").Find(&transactions, "donatur_id = ?", ID).Error

	return transactions, err
}

func (r *repository) GetTransaction(ID int) (models.Transaction, error) {
	var transaction models.Transaction
	err := r.db.Preload("Fund").Preload("Fund.User").Preload("User").First(&transaction, "id = ?", ID).Error

	return transaction, err
}

func (r *repository) CreateTransaction(transaction models.Transaction) (models.Transaction, error) {
	err := r.db.Create(&transaction).Error

	return transaction, err
}

func (r *repository) GetOneTransaction(ID string) (models.Transaction, error) {
	var transaction models.Transaction
	err := r.db.Preload("Fund").Preload("Fund.User").Preload("User").First(&transaction, "id = ?", ID).Error

	return transaction, err
}

func (r *repository )UpdateTransaction(status string, ID string) error  {
	var transaction models.Transaction

	r.db.Preload("Fund").Preload("Fund.User").Preload("User").First(&transaction, ID)

	if status != transaction.Status && status == "success" {
		transaction.Status = status
	}

	err := r.db.Save(&transaction).Error

	return err
}
