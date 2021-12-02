package allowlistsrv

import (
	"github.com/mbravovaisma/authorizer/internal/core/constants"
	"github.com/mbravovaisma/authorizer/internal/core/domain"
	"github.com/mbravovaisma/authorizer/internal/core/ports"
	"github.com/mbravovaisma/authorizer/pkg/log"
)

type service struct {
	repository ports.AccountRepository
}

func New(repository ports.AccountRepository) ports.AllowListService {
	return &service{
		repository: repository,
	}
}

func (s *service) Set(active bool) (*domain.Movement, error) {
	account, err := s.repository.Get(constants.AccountID)
	if err != nil {
		log.Error("Error getting account", log.ErrorField(err))
		return nil, err
	}

	account.SetAllowList(active)

	err = s.repository.Save(constants.AccountID, *account)
	if err != nil {
		log.Error("Error save account", log.ErrorField(err))
		return nil, err
	}

	return &domain.Movement{
		Account:    account,
		Violations: []string{},
	}, nil
}
