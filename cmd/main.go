package main

import (
	"context"
	"google.golang.org/grpc"
	"log"
	"net"
	"user-management/pkg"
	"user-management/proto"
)

func main() {
	server := &pkg.Server{
		GrpcPort: "50051",
		Opt:      &grpc.EmptyServerOption{},
	}
	lis, err := net.Listen("tcp", ":"+server.GrpcPort)
	if err != nil {
		log.Panicf("failed to listen on port %s: %v", server.GrpcPort, err)
	}
	grpcServer := grpc.NewServer(server.Opt)
	user.RegisterUserServiceServer(grpcServer, server)
	log.Printf("server started on port %s", server.GrpcPort)
	serveGRPC(grpcServer, lis, func() {})
}

func serveGRPC(grpcS *grpc.Server, listener net.Listener, stop context.CancelFunc) {
	defer stop()

	if err := grpcS.Serve(listener); err != nil {
		log.Printf("stopped serving: %s", err.Error())
	}
}
