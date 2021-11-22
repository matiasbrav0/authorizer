package app

import (
	"bufio"
	"fmt"
	"os"

	"github.com/mbravovaisma/authorizer/pkg/log"
	"go.uber.org/zap"
)

func Start() {
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		if err := scanner.Err(); err != nil {
			log.Error("error while scan a new line of operations", zap.Error(err))
		}

		fmt.Println(string(scanner.Bytes()))
	}
}
