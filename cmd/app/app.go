package app

import (
	"bufio"
	"os"

	"github.com/mbravovaisma/authorizer/internal/core/services"
	"github.com/mbravovaisma/authorizer/internal/operation"
	"github.com/mbravovaisma/authorizer/internal/repository/memory"

	"github.com/mbravovaisma/authorizer/pkg/log"
	"go.uber.org/zap"
)

func Start() {
	scanner := bufio.NewScanner(os.Stdin)

	mem := memory.NewMemory()
	trxservice := services.NewTransaction(mem)
	accountService := services.NewAccount(mem)
	selector := operation.NewSelector(accountService, trxservice)

	for scanner.Scan() {
		if err := scanner.Err(); err != nil {
			log.Error("error while scan a new line of operations", zap.Error(err))
		}

		_, _ = selector.OperationSelector(scanner.Bytes())
	}
}
