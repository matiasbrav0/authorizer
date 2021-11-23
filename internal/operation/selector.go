package operation

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/mbravovaisma/authorizer/pkg/log"
	"go.uber.org/zap"

	"github.com/mbravovaisma/authorizer/internal/core/domain"

	"github.com/mbravovaisma/authorizer/internal/core/ports"
)

type selector struct {
	accountService     ports.Account
	transactionService ports.Transaction
}

func NewSelector(accountService ports.Account, transactionService ports.Transaction) ports.Selector {
	return &selector{
		accountService:     accountService,
		transactionService: transactionService,
	}
}

func (s *selector) OperationSelector(request []byte) (domain.Response, error) {
	operation := string(request)

	/* Perform an account operation */
	if strings.Contains(operation, "account") {
		var accountOperation domain.AccountRequest
		if err := json.Unmarshal(request, &accountOperation); err != nil {
			log.Error("can't unmarshal account request", zap.Error(err))
			return domain.Response{}, err
		}

		return s.accountService.CreateAccount(&accountOperation.Account)
	}

	/* Perform a transaction operation */
	if strings.Contains(operation, "transaction") {
		var transactionOperation domain.TransactionRequest
		if err := json.Unmarshal(request, &transactionOperation); err != nil {
			log.Error("can't unmarshal transaction request", zap.Error(err))
			return domain.Response{}, err
		}

		return s.transactionService.PerformTransaction(&transactionOperation.Transaction)
	}

	/* Invalid operation */
	err := fmt.Errorf("invalid operation, request: %s", request)
	log.Error("invalid operation", zap.Error(err))

	return domain.Response{}, err
}
