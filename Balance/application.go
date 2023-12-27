package app

import (
	"github.com/your_username/balances/internal/domain"
	"github.com/your_username/balances/internal/infrastructure/persistence"
)

type BalanceService struct {
	balanceRepository persistence.BalanceRepository
}

func NewBalanceService(balanceRepository persistence.BalanceRepository) *BalanceService {
	return &BalanceService{
		balanceRepository: balanceRepository,
	}
}

func (s *BalanceService) UpdateBalance(accountID string, updatedBalance float64) error {
	// Implement logic to update balance
	return s.balanceRepository.UpdateBalance(accountID, updatedBalance)
}

func (s *BalanceService) GetBalance(accountID string) (float64, error) {
	// Implement logic to retrieve balance
	return s.balanceRepository.GetBalance(accountID)
}