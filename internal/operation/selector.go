package operation

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/mbravovaisma/authorizer/pkg/log"
	"go.uber.org/zap"

	"github.com/mbravovaisma/authorizer/internal/core/ports"
)

type selector struct {
	accountService     ports.AccountService
	transactionService ports.TransactionService
}

func NewSelector(accountService ports.AccountService, transactionService ports.TransactionService) *selector {
	return &selector{
		accountService:     accountService,
		transactionService: transactionService,
	}
}

func (s *selector) OperationSelector(request []byte) (Response, error) {
	operation := string(request)

	/* Perform an account operation */
	if strings.Contains(operation, "account") {
		var accountOperation AccountOperation
		if err := json.Unmarshal(request, &accountOperation); err != nil {
			log.Error("can't unmarshal account request", zap.Error(err))
			return Response{}, err
		}

		r, err := s.accountService.Create(accountOperation.Account.ActiveCard, accountOperation.Account.AvailableLimit)
		if err != nil {
			log.Error("error creating account", zap.Error(err))
			return Response{}, err
		}

		return BuildResponse(r), nil
	}

	/* Perform a transaction operation */
	if strings.Contains(operation, "transaction") {
		var transactionOperation TransactionOperation
		if err := json.Unmarshal(request, &transactionOperation); err != nil {
			log.Error("can't unmarshal transaction request", zap.Error(err))
			return Response{}, err
		}

		r, err := s.transactionService.PerformTransaction(
			transactionOperation.Transaction.Amount,
			transactionOperation.Transaction.Merchant,
			transactionOperation.Transaction.Time,
		)
		if err != nil {
			log.Error("error performing transaction", zap.Error(err))
			return Response{}, err
		}

		return BuildResponse(r), nil
	}

	/* Invalid operation */
	err := fmt.Errorf("invalid operation, request: %s", request)
	log.Error("invalid operation", zap.Error(err))

	return Response{}, err
}
