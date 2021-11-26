package accountrepo

import (
	"errors"

	"github.com/mbravovaisma/authorizer/internal/core/domain"
	"github.com/mbravovaisma/authorizer/internal/core/ports"
)

type memKVS struct {
	kvs map[string]domain.Account
}

func New() ports.AccountRepository {
	return &memKVS{
		kvs: make(map[string]domain.Account),
	}
}

func (m *memKVS) Get(id string) (*domain.Account, error) {
	account, exist := m.kvs[id]
	if !exist {
		return nil, errors.New("account not exist")
	}

	return &account, nil
}

func (m *memKVS) Save(id string, account domain.Account) error {
	m.kvs[id] = account
	return nil
}

func (m *memKVS) Exist(id string) bool {
	_, exist := m.kvs[id]
	return exist
}
