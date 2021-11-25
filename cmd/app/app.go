package app

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"

	"github.com/mbravovaisma/authorizer/internal/core/services/accountsrv"
	"github.com/mbravovaisma/authorizer/internal/core/services/transactionsrv"
	"github.com/mbravovaisma/authorizer/internal/operation"
	"github.com/mbravovaisma/authorizer/internal/repositories/accountrepo"

	"github.com/mbravovaisma/authorizer/pkg/log"
)

func Start() {
	scanner := bufio.NewScanner(os.Stdin)

	mem := accountrepo.New()
	transactionService := transactionsrv.New(mem)
	accountService := accountsrv.New(mem)
	selector := operation.NewSelector(accountService, transactionService)

	for scanner.Scan() {
		if err := scanner.Err(); err != nil {
			log.Error("error while scan a new line of operations", log.ErrorField(err))
		}

		response, _ := selector.OperationSelector(scanner.Bytes())
		responseJson, _ := json.Marshal(response)
		fmt.Println(string(responseJson))
	}
}
