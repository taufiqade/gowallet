package service

import (
	dbEntity "github.com/taufiqade/gowallet/models"
	"github.com/taufiqade/gowallet/models/http/request"
)

// TransactionService godoc
type TransactionService struct {
	userRepo        dbEntity.IUserRepository
	userBalanceRepo dbEntity.IUserBalanceRepository
	historyRepo     dbEntity.IUserBalanceHistoryRepository
}

// NewTransactionService initialize new transaction service
func NewTransactionService(u dbEntity.IUserRepository, ub dbEntity.IUserBalanceRepository, uh dbEntity.IUserBalanceHistoryRepository) *TransactionService {
	return &TransactionService{
		userRepo:        u,
		userBalanceRepo: ub,
		historyRepo:     uh,
	}
}

// TopUp godoc
func (u *TransactionService) TopUp(userID int, payload *request.TopUpRequest) error {
	beneficiary, err := u.userBalanceRepo.GetByUserID(userID)
	if err != nil {
		return err
	}
	currBalance := beneficiary.Balance
	// update balanceData
	balanceData := &dbEntity.UserBalance{
		ID:             beneficiary.ID,
		UserID:         beneficiary.UserID,
		Balance:        beneficiary.Balance + float64(payload.Amount),
		BalanceAchieve: beneficiary.BalanceAchieve + float64(payload.Amount),
	}
	ubErr := u.userBalanceRepo.Update(userID, balanceData)
	if ubErr != nil {
		return ubErr
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
	bhErr := u.historyRepo.Create(balanceHistory)
	if bhErr != nil {
		return bhErr
	}
	return bhErr
}
