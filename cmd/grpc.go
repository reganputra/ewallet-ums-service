package cmd

import (
	"ewallet-ums/cmd/proto/tokenValidation"
	"ewallet-ums/helpers"
	"log"
	"net"

	"google.golang.org/grpc"
)

func ServerGRPC() {

	deps := InitializeDependencies()

	lis, err := net.Listen("tcp", ":"+helpers.GetEnv("GRPC_PORT", "7000"))
	if err != nil {
		log.Fatal("Failed to start gRPC server:", err)
	}

	s := grpc.NewServer()

	tokenValidation.RegisterTokenValidationServer(s, &deps.TokenValidationAPI)

	if err := s.Serve(lis); err != nil {
		log.Fatal("Failed to start gRPC server:", err)
	}
}
