package memory

import (
	"github.com/mbravovaisma/authorizer/internal/core/domain"
	"github.com/mbravovaisma/authorizer/internal/core/ports"
)

type memory struct {
	mapper map[string]domain.AccountData
}

func NewMemory() ports.AuthorizerRepository {
	return &memory{mapper: make(map[string]domain.AccountData)}
}

func (m *memory) AccountCreate(id string, accountToSave *domain.Account) *domain.AccountData {
	/* Create object to save */
	data := domain.AccountData{
		AccountInfo: &domain.AccountInfo{
			ActiveCard:     accountToSave.Account.ActiveCard,
			AvailableLimit: accountToSave.Account.AvailableLimit,
		},
		AccountMovements: []domain.Response{{
			Account: domain.AccountInfo{
				ActiveCard:     accountToSave.Account.ActiveCard,
				AvailableLimit: accountToSave.Account.AvailableLimit,
			},
			Violations: []string{},
		}},
		TransactionsInfo: []domain.Transaction{},
	}

	m.save(id, data)

	return &data
}

func (m *memory) AccountExist(id string) bool {
	_, exist := m.mapper[id]
	return exist
}

func (m *memory) UpdateAccountData(
	id string,
	accountInfo *domain.Account,
	transaction *domain.Transaction,
	accountMovement *domain.Response,
) *domain.AccountData {

	if !m.AccountExist(id) {
		return &domain.AccountData{}
	}

	data := m.GetAccountData(id)

	data.AccountInfo.ActiveCard = accountInfo.Account.ActiveCard
	data.AccountInfo.AvailableLimit = accountInfo.Account.AvailableLimit
	data.AccountMovements = append(data.AccountMovements, *accountMovement)
	data.TransactionsInfo = append(data.TransactionsInfo, *transaction)

	m.save(id, *data)

	return data
}

func (m *memory) GetAccountData(id string) *domain.AccountData {
	if !m.AccountExist(id) {
		return &domain.AccountData{}
	}

	accountData, _ := m.mapper[id]
	return &accountData
}

func (m *memory) save(id string, data domain.AccountData) {
	m.mapper[id] = data
}
