package account

import (
	"context"
	domain "money-manager/internal/domain/account"
	"money-manager/internal/infra/dynamodb/repository"
)

type GetAccountsByUserIdUseCase struct {
	accountRepository *repository.AccountRepository
 
}

func NewGetAccountsByUserIdUseCase(accountRepository *repository.AccountRepository) *GetAccountsByUserIdUseCase {
	return &GetAccountsByUserIdUseCase{
		accountRepository: accountRepository,
	}
}

func (uc *GetAccountsByUserIdUseCase) Execute(
	ctx context.Context,
	userId string,
	) ([]domain.Account, error) {

	accounts, err := uc.accountRepository.GetAccountsByUserId(userId, ctx)
	
	if err != nil {
		return nil, err
	}

	return accounts, nil
}