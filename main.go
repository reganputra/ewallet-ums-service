package main

import (
	"ewallet-ums/cmd"
	"ewallet-ums/helpers"
)

func main() {

	// Setup logger
	helpers.SetupLogger()
	// Setup config
	helpers.SetupConfig()
	// load database
	helpers.SetupMySql()

	// start http server
	cmd.ServerHttp()

	// start grpc server
	cmd.ServerGRPC()

}
