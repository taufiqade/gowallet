package models

type UserBalance struct {
	ID              uint 	`gorm:"primary_key" json:"id"`
	UserID          uint 	`gorm:"column:user_id" json:"user_id"`
	Balance         float64 `gorm:"column:balance" gorm:"type:decimal(10,2)" json:"balance"`
	BalanceAchieve 	float64 `gorm:"column:balance_achieve" gorm:"type:decimal(10,2)" json:"balance_achieve"`
}


// repository contract
type IUserBalanceRepository interface {
	GetByUserID(userId int) (UserBalance, error)
	Update(userId int, data *UserBalance) error
}