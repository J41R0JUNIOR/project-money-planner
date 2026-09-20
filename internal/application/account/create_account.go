package account

import (
	"context"
	domain "money-manager/internal/domain/account"
	root "money-manager/internal/domain"
	"github.com/google/uuid"
	"money-manager/internal/infra/dynamodb/repository"
)


type CreateAccountUseCase struct {
	accountRepository *repository.AccountRepository
}

func NewCreateAccountUseCase(accountRepository *repository.AccountRepository) *CreateAccountUseCase {
	return &CreateAccountUseCase{
		accountRepository: accountRepository,
	}
}

func (uc *CreateAccountUseCase) Execute(
	ctx context.Context,
	userId string,
	name string,
	accountType domain.AccountType,
	balance root.Money,
) error {
	var newAccount = domain.Account{
		Id:     uuid.New().String(),
		UserId: userId,
		Name:   name,
		Type:   accountType,
		Balance: balance,
	}

	err := uc.accountRepository.SaveAccount(newAccount, ctx)
	
	if err != nil {
		return err
	}

	return nil
}