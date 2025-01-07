package main

import (
	"github.com/jetaimejeteveux/e-wallet-ums/cmd"
	"github.com/jetaimejeteveux/e-wallet-ums/helpers"
)

func main() {

	// load config
	helpers.SetupConfig()

	// load log
	helpers.SetupLogger()

	// load db
	helpers.SetupMySQL()

	// run grpc
	go cmd.ServeGRPC()

	// run http
	cmd.ServeHTTP()
}
