package cmd

import (
	"log"
	"net"

	pb "github.com/jetaimejeteveux/e-wallet-ums/cmd/proto/tokenvalidation"
	"github.com/jetaimejeteveux/e-wallet-ums/helpers"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

func ServeGRPC() {
	dependency := dependencyInject()
	s := grpc.NewServer()
	// list method
	pb.RegisterTokenValidationServer(s, dependency.TokenValidationAPI)

	lis, err := net.Listen("tcp", ":"+helpers.GetEnv("GRPC_PORT", "7000"))
	if err != nil {
		log.Fatal("failed to listen grpc port: ", err)
	}

	// list method
	// pb.ExampleMethod(s, &grpc....)

	logrus.Info("start listening grpc on port:" + helpers.GetEnv("GRPC_PORT", "7000"))
	if err := s.Serve(lis); err != nil {
		log.Fatal("failed to serve grpc port: ", err)
	}
}
