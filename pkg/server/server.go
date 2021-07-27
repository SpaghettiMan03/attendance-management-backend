package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"

	"attendance-management-backend/pkg/presentation/handler"
	schema "attendance-management-backend/schema/gen/server"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
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
	reflection.Register(server)

	go func() {
		log.Printf("start gRPC server port: %v", port)
		server.Serve(lis)
	}()

	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit
	log.Println("stopping gRPC server")
	server.GracefulStop()
}