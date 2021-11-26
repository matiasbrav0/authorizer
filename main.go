package main

import (
	"os"

	"github.com/mbravovaisma/authorizer/cmd/app"
)

func main() {
	app.Start(os.Stdin, os.Stdout)
}
