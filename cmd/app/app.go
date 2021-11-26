package app

import (
	"bufio"
	"encoding/json"
	"io"

	"github.com/mbravovaisma/authorizer/cmd/dependencies"

	"github.com/mbravovaisma/authorizer/pkg/log"
)

func Start(stdin io.Reader, stdout io.Writer) {
	scanner := bufio.NewScanner(stdin)
	writer := bufio.NewWriter(stdout)

	d := dependencies.NewDependencies()

	for scanner.Scan() {
		if err := scanner.Err(); err != nil {
			log.Error("error while scan a new line of operations", log.ErrorField(err))
		}

		response, err := d.Selector.OperationSelector(scanner.Bytes())
		if err != nil {
			log.Error("error while trying to process the request", log.ErrorField(err))
		}

		responseJson, err := json.Marshal(response)
		if err != nil {
			log.Error("error unmarshalling the response", log.ErrorField(err))
		}

		_, _ = writer.Write(responseJson)
		_, _ = writer.Write([]byte("\n"))
		_ = writer.Flush()
	}
}
