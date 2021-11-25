package accountrepo

import (
	"github.com/mbravovaisma/authorizer/internal/core/domain"
)

type memKVS struct {
	kvs map[string]domain.Account
}

func New() *memKVS {
	return &memKVS{
		kvs: make(map[string]domain.Account),
	}
}

func (m *memKVS) Get(id string) (domain.Account, error) {
	return m.kvs[id], nil
}

func (m *memKVS) Save(id string, account domain.Account) error {
	m.kvs[id] = account
	return nil
}

func (m *memKVS) Exist(id string) bool {
	_, exist := m.kvs[id]
	return exist
}
