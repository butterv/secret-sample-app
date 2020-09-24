package main

import (
	"fmt"
	"net"
	"os"

	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_recovery "github.com/grpc-ecosystem/go-grpc-middleware/recovery"
	grpc_validator "github.com/grpc-ecosystem/go-grpc-middleware/validator"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	_ "google.golang.org/grpc/encoding/gzip"
	"google.golang.org/grpc/reflection"

	"github.com/istsh/secret-sample-app/app/domain/service"
	healthpb "github.com/istsh/secret-sample-app/app/interface/rpc/v1/health"
)

func main() {
	listenPort, err := net.Listen("tcp", ":8080")
	if err != nil {
		logrus.Fatalln(err)
	}

	if host := os.Getenv("MYSQL_HOST"); host != "" {
		fmt.Printf("MYSQL_HOST: %s\n", host)
	}
	if user := os.Getenv("MYSQL_USER"); user != "" {
		fmt.Printf("MYSQL_USER: %s\n", user)
	}
	if password := os.Getenv("MYSQL_PASSWORD"); password != "" {
		fmt.Printf("MYSQL_PASSWORD: %s\n", password)
	}


	s := newGRPCServer()
	reflection.Register(s)
	_ = s.Serve(listenPort)
	s.GracefulStop()
}

func newGRPCServer() *grpc.Server {
	s := grpc.NewServer(
		grpc.UnaryInterceptor(grpc_middleware.ChainUnaryServer(
			grpc_validator.UnaryServerInterceptor(),
			grpc_recovery.UnaryServerInterceptor(),
		)),
	)

	healthpb.RegisterHealthServer(s, service.NewHealthService())

	return s
}
