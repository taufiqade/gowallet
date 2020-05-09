package db

import (
	"github.com/jinzhu/gorm"
	"github.com/taufiqade/gowallet/models"
)

// UserBalanceHistoryRepository godoc
type UserBalanceHistoryRepository struct {
	DB *gorm.DB
}

// NewUserBalanceHistoryRepository godoc
func NewUserBalanceHistoryRepository(db *gorm.DB) models.IUserBalanceHistoryRepository {
	return &UserBalanceHistoryRepository{DB: db}
}

// GetBalanceID godoc
func (repo *UserBalanceHistoryRepository) GetBalanceID(id int) (models.UserBalanceHistory, error) {
	var uh models.UserBalanceHistory
	query := repo.DB.Table("user_balance_history").Where("user_balance_id=?", id).First(&uh)
	return uh, query.Error
}

// Create godoc
func (repo *UserBalanceHistoryRepository) Create(data *models.UserBalanceHistory) error {
	query := repo.DB.Table("user_balance_history").Create(data)
	return query.Error
}
