package service

import (
	"context"
	"fmt"
	"payment_processing_system/internal/domain"
	"payment_processing_system/internal/domain/entity"

	"github.com/shopspring/decimal"
)

type BalanceStorage interface {
	GetByID(ctx context.Context, id int64) (*entity.Balance, error)
	// GetAll(ctx context.Context, limit, offset int) ([]entity.Balance, error)
	// Create(ctx context.Context, balance entity.Balance) (entity.Balance, error)
	// Update(ctx context.Context, id string, amount int64) error
	IncreaseAmount(ctx context.Context, id int64, amount decimal.Decimal) error
	DecreaseAmount(ctx context.Context, id int64, amount decimal.Decimal) error
}

type BalanceService struct {
	storage BalanceStorage
}

func NewBalanceService(storage BalanceStorage) *BalanceService {
	return &BalanceService{storage: storage}
}

func (s *BalanceService) GetByID(ctx context.Context, id int64) (*entity.Balance, error) {
	return s.storage.GetByID(ctx, id)
}

// func (s BalanceService) Create(ctx context.Context, balance expectedTransaction.Balance) (expectedTransaction.Balance, error) {
//	return s.testStorage.Create(ctx, balance)
// }

func (s *BalanceService) ChangeAmount(ctx context.Context, id int64, amount decimal.Decimal) error {
	if amount.IsZero() {
		return fmt.Errorf("id = %q ; amount = %s ; %w", id, amount.String(), domain.ChangeBalanceByZeroAmountErr)
	} else if amount.IsPositive() {
		return s.storage.IncreaseAmount(ctx, id, amount)
	}
	return s.storage.DecreaseAmount(ctx, id, amount.Neg())
}
