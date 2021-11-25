package accountsrv

import (
	"github.com/mbravovaisma/authorizer/internal/core/constants"
	"github.com/mbravovaisma/authorizer/internal/core/domain"
	"github.com/mbravovaisma/authorizer/internal/core/ports"
	"github.com/mbravovaisma/authorizer/pkg/log"
)

type service struct {
	repository ports.AccountRepository
}

func New(repository ports.AccountRepository) ports.AccountService {
	return &service{
		repository: repository,
	}
}

func (s *service) Create(activeCard bool, availableLimit int64) (domain.Movement, error) {
	// Check If account exist
	if s.repository.Exist(constants.AccountID) {
		account, err := s.repository.Get(constants.AccountID)
		if err != nil {
			log.Error("Error getting account", log.ErrorField(err))
			return domain.Movement{}, err
		}

		return domain.Movement{
			Account:    &account,
			Violations: []string{constants.AccountAlreadyInitialized},
		}, nil
	}

	// Create dto to save
	account := domain.NewAccount(activeCard, availableLimit)

	// Save an account
	err := s.repository.Save(constants.AccountID, account)
	if err != nil {
		log.Error("Error save account", log.ErrorField(err))
		return domain.Movement{}, err
	}

	return domain.Movement{
		Account:    &account,
		Violations: []string{},
	}, nil
}
