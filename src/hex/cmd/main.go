package main

import (
	"hex/internal/adapters/app/api"
	"hex/internal/adapters/core/arithmetic"
	"hex/internal/adapters/framework/left/grpc"
	db "hex/internal/adapters/framework/right/db"
	"hex/internal/ports"
	"log"
	"os"
)

func main() {

	var err error

	var dbaseAdapter ports.DbPort
	var core ports.ArithmeticPort
	var appAdapter ports.APIport
	var gRPCAdapter ports.GRPCPort
	//ports
	// var core ports.ArithmeticPort
	dbaseDriver := os.Getenv("DB_DRIVER")
	dsourceName := os.Getenv("DS_NAME")
	dbaseAdapter, err = db.NewAdapter(dbaseDriver, dsourceName)

	if err != nil {
		log.Fatalf("failed to initiate dbase connection: %v", err)
	}
	defer dbaseAdapter.CloseDbConnection()
	// core = arithmetic.NewAdapter()

	core = arithmetic.NewAdapter()

	appAdapter = api.NewAdapter(dbaseAdapter, core)

	gRPCAdapter = grpc.NewAdapter(appAdapter)

	gRPCAdapter.Run()
}
