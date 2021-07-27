package main

import (
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
	schema "github.com/SpaghettiMan03/attendance-management-backend/schema/gen/server"
)

func main() {
	port := 50051
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		log.Fatalf("failed to listen:%v", err)
	}

	fmt.Println(lis)

	server := grpc.NewServer()

	go func() {
		log.Printf("start gRPC server port: %v", port)
		server.Serve(lis)
	}()
}