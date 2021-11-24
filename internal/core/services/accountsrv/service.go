package accountsrv

import (
	violations "github.com/mbravovaisma/authorizer/internal/core/constants"
	"github.com/mbravovaisma/authorizer/internal/core/domain"
	"github.com/mbravovaisma/authorizer/internal/core/ports"
	"github.com/mbravovaisma/authorizer/pkg/constants"
	"github.com/mbravovaisma/authorizer/pkg/log"
	"go.uber.org/zap"
)

type service struct {
	repository ports.AuthorizerRepository
}

func New(repository ports.AuthorizerRepository) *service {
	return &service{
		repository: repository,
	}
}

func (s *service) Create(activeCard bool, availableLimit int64) (domain.Movement, error) {
	/* Check If account exist */
	if s.repository.Exist(constants.AccountID) {
		account, err := s.repository.Get(constants.AccountID)
		if err != nil {
			log.Error("Error getting account", zap.Error(err))
			return domain.Movement{}, err
		}

		return domain.Movement{
			Account:    &account,
			Violations: []string{violations.AccountAlreadyInitialized},
		}, nil
	}

	/* Create dto to save */
	account := domain.NewAccount(activeCard, availableLimit)

	/* Save an account */
	err := s.repository.Save(constants.AccountID, account)
	if err != nil {
		log.Error("Error save account", zap.Error(err))
		return domain.Movement{}, err
	}

	return domain.Movement{
		Account:    &account,
		Violations: []string{},
	}, nil
}
