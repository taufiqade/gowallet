package db

import (
	"github.com/jinzhu/gorm"
	"github.com/taufiqade/gowallet/models"
)

type UserBalanceHistoryRepository struct {
	DB *gorm.DB
}

func NewUserBalanceHistoryRepository(db *gorm.DB) models.IUserBalanceHistoryRepository {
	return &UserBalanceHistoryRepository{DB: db}
}

func (repo *UserBalanceHistoryRepository) GetBalanceID(id int) (models.UserBalanceHistory, error) {
	var uh models.UserBalanceHistory
	query := repo.DB.Table("user_balance_history").Where("user_balance_id=?", id).First(&uh)
	return uh, query.Error
}

func (repo *UserBalanceHistoryRepository) Create(data *models.UserBalanceHistory) error {
	query := repo.DB.Table("user_balance_history").Create(data)
	return query.Error
}