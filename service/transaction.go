package service

import (
	"fmt"

	"github.com/jinzhu/gorm"
	dbEntity "github.com/taufiqade/gowallet/models"
	"github.com/taufiqade/gowallet/models/http/request"
)

// TransactionService godoc
type TransactionService struct {
	userRepo        dbEntity.IUserRepository
	userBalanceRepo dbEntity.IUserBalanceRepository
	historyRepo     dbEntity.IUserBalanceHistoryRepository
	db              *gorm.DB
}

// NewTransactionService initialize new transaction service
func NewTransactionService(u dbEntity.IUserRepository, ub dbEntity.IUserBalanceRepository, uh dbEntity.IUserBalanceHistoryRepository, dbConn *gorm.DB) *TransactionService {
	return &TransactionService{
		userRepo:        u,
		userBalanceRepo: ub,
		historyRepo:     uh,
		db:              dbConn,
	}
}

// TopUp godoc
func (u *TransactionService) TopUp(email string, payload *request.TransactionRequest) error {
	tx := u.db.Begin()
	if err := tx.Error; err != nil {
		return err
	}

	user, err := u.userRepo.GetUserByEmail(email)
	if err != nil {
		return err
	}
	beneficiary, err := u.userBalanceRepo.GetByUserID(int(user.ID))
	if err != nil {
		return err
	}
	currBalance := beneficiary.Balance

	//insert user balance
	balanceData := &dbEntity.UserBalance{
		UserID:         beneficiary.UserID,
		Balance:        beneficiary.Balance + float64(payload.Amount),
		BalanceAchieve: beneficiary.BalanceAchieve + float64(payload.Amount),
	}
	if err := u.userBalanceRepo.Create(balanceData, tx); err != nil {
		tx.Rollback()
		return err
	}
	//insert balance history
	balanceHistory := &dbEntity.UserBalanceHistory{
		UserBalanceID: beneficiary.ID,
		BalanceBefore: currBalance,
		BalanceAfter:  balanceData.BalanceAchieve,
		Activity:      "TopUp",
		Type:          "debit",
		IP:            payload.IP,
		UserAgent:     payload.UserAgent,
		Location:      payload.Location,
		Author:        payload.Author,
	}
	if bhErr := u.historyRepo.Create(balanceHistory, tx); bhErr != nil {
		tx.Rollback()
		return bhErr
	}
	tx.Commit()
	return nil
}

// Transfer godoc
func (u *TransactionService) Transfer(obligorID int, email string, payload *request.TransactionRequest) error {
	tx := u.db.Begin()
	if err := tx.Error; err != nil {
		return err
	}

	user, err := u.userRepo.GetUserByEmail(email)
	if err != nil {
		return err
	}
	obligor, _ := u.userBalanceRepo.GetByUserID(obligorID)
	beneficiary, bErr := u.userBalanceRepo.GetByUserID(int(user.ID))
	if bErr != nil {
		return bErr
	}

	if obligor.Balance < float64(payload.Amount) {
		return fmt.Errorf("Insuficent Balance")
	}

	// create credit transaction for authenticated user
	ob := new(dbEntity.UserBalance)
	ob.UserID = uint(obligorID)
	ob.Balance = obligor.Balance - float64(payload.Amount)
	ob.BalanceAchieve = obligor.BalanceAchieve - float64(payload.Amount)
	if err := u.userBalanceRepo.Create(ob, tx); err != nil {
		return err
	}

	currBalance := obligor.Balance
	balanceHistory := &dbEntity.UserBalanceHistory{
		UserBalanceID: obligor.ID,
		BalanceBefore: currBalance,
		BalanceAfter:  ob.BalanceAchieve,
		Activity:      "Transfer",
		Type:          "credit",
		IP:            payload.IP,
		UserAgent:     payload.UserAgent,
		Location:      payload.Location,
		Author:        payload.Author,
	}
	bhErr := u.historyRepo.Create(balanceHistory, tx)
	if bhErr != nil {
		tx.Rollback()
		return bhErr
	}
	// send debit transaction
	if err := u.DebitTransaction(&beneficiary, payload, tx); err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()
	return nil
}

// DebitTransaction godoc
func (u *TransactionService) DebitTransaction(beneficiary *dbEntity.UserBalance, payload *request.TransactionRequest, tx *gorm.DB) error {
	// create credit transaction for authenticated user
	userBalance := new(dbEntity.UserBalance)
	userBalance.UserID = uint(beneficiary.UserID)
	userBalance.Balance = beneficiary.Balance + float64(payload.Amount)
	userBalance.BalanceAchieve = beneficiary.BalanceAchieve + float64(payload.Amount)

	// it should be created new one
	if err := u.userBalanceRepo.Create(userBalance, tx); err != nil {
		return err
	}
	currBalance := beneficiary.Balance
	balanceHistory := &dbEntity.UserBalanceHistory{
		UserBalanceID: beneficiary.ID,
		BalanceBefore: currBalance,
		BalanceAfter:  userBalance.BalanceAchieve,
		Activity:      "Transfer",
		Type:          "debit",
		IP:            payload.IP,
		UserAgent:     payload.UserAgent,
		Location:      payload.Location,
		Author:        payload.Author,
	}
	if err := u.historyRepo.Create(balanceHistory, tx); err != nil {
		return err
	}

	return nil
}
