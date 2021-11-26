package operation

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/mbravovaisma/authorizer/pkg/log"

	"github.com/mbravovaisma/authorizer/internal/core/ports"
)

type Selector struct {
	accountService     ports.AccountService
	transactionService ports.TransactionService
}

func NewSelector(accountService ports.AccountService, transactionService ports.TransactionService) *Selector {
	return &Selector{
		accountService:     accountService,
		transactionService: transactionService,
	}
}

func (s *Selector) OperationSelector(request []byte) (interface{}, error) {
	operation := string(request)

	// Perform an account operation
	if strings.Contains(operation, "account") {
		var accountOperation AccountOperation
		if err := json.Unmarshal(request, &accountOperation); err != nil {
			log.Error("can't unmarshal account request", log.ErrorField(err))
			return nil, err
		}

		response, err := s.accountService.Create(accountOperation.Account.ActiveCard, accountOperation.Account.AvailableLimit)
		if err != nil {
			log.Error("error creating account", log.ErrorField(err))
			return nil, err
		}

		return response, nil
	}

	// Perform a transaction operation
	if strings.Contains(operation, "transaction") {
		var transactionOperation TransactionOperation
		if err := json.Unmarshal(request, &transactionOperation); err != nil {
			log.Error("can't unmarshal transaction request", log.ErrorField(err))
			return nil, err
		}

		response, err := s.transactionService.PerformTransaction(
			transactionOperation.Transaction.Amount,
			transactionOperation.Transaction.Merchant,
			transactionOperation.Transaction.Time,
		)
		if err != nil {
			log.Error("error performing transaction", log.ErrorField(err))
			return nil, err
		}

		return response, nil
	}

	// Invalid operation
	err := fmt.Errorf("invalid operation, request: %s", request)
	log.Error("invalid operation", log.ErrorField(err))

	return nil, err
}
