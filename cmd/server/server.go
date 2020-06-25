package main

import (
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"grpc-calculator/pb"
	"grpc-calculator/servers"
	"log"
	"net"
)
func main(){
	var port string = "58885"
	grpcServer := grpc.NewServer()
	pb.RegisterMathServiceServer(grpcServer, &servers.Math{})
	reflection.Register(grpcServer)

	listener, err := net.Listen("tcp", ":"+port)
	if err != nil {
		log.Fatalf("Cannot start the grpc server: %v", err)
	}

	log.Println("GRPC Server running on port:", port)
	grpcServer.Serve(listener)
}
