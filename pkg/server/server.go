package main

import (
	"fmt"
	"log"
	"net"

	"github.com/SpaghettiMan03/attendance-management-backend/pkg/presentation/handler"
	schema "github.com/SpaghettiMan03/attendance-management-backend/schema/gen/server"
	"google.golang.org/grpc"
)

func main() {
	port := 50051
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		log.Fatalf("failed to listen:%v", err)
	}

	fmt.Println(lis)

	server := grpc.NewServer()

	schema.RegisterEmployeeServiceServer(
		server,
		handler.NewEmployeeHandler(),
		)

	go func() {
		log.Printf("start gRPC server port: %v", port)
		server.Serve(lis)
	}()
}