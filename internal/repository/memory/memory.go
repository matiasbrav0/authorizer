package memory

import (
	"time"

	"github.com/mbravovaisma/authorizer/internal/core/domain"
	"github.com/mbravovaisma/authorizer/internal/core/ports"
)

type memory struct {
	mapper map[string]domain.AuthorizerData
}

func NewMemory() ports.AuthorizerRepository {
	return &memory{mapper: make(map[string]domain.AuthorizerData)}
}

func (m *memory) AccountCreate(id string, account *domain.Account) *domain.AuthorizerData {
	/* Create object to save */
	data := domain.AuthorizerData{
		Account: account,
		AccountMovements: []domain.Response{{
			Account:    account,
			Violations: []string{},
		}},
		TransactionsInfo: []domain.Transaction{},
		Attempts:         0,
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
	account *domain.Account,
	transaction *domain.Transaction,
	accountMovement *domain.Response,
	authorizationTime time.Time,
	attempts uint8,
) *domain.AuthorizerData {

	if !m.AccountExist(id) {
		return &domain.AuthorizerData{}
	}

	data := m.GetAccountData(id)

	data.Account = account
	data.AccountMovements = append(data.AccountMovements, *accountMovement)
	data.TransactionsInfo = append(data.TransactionsInfo, *transaction)
	data.AuthorizationTime = authorizationTime
	data.Attempts = attempts

	m.save(id, *data)

	return data
}

func (m *memory) GetAccountData(id string) *domain.AuthorizerData {
	if !m.AccountExist(id) {
		return &domain.AuthorizerData{}
	}

	accountData, _ := m.mapper[id]
	return &accountData
}

func (m *memory) save(id string, data domain.AuthorizerData) {
	m.mapper[id] = data
}
