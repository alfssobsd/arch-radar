package main

import (
	"arch-radar/backend-service/backend-service/entrypoints/grpc_entrypoint"
	pb "arch-radar/backend-service/backend-service/entrypoints/grpc_entrypoint/gen"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
)

const PORT = "9090"

func main() {

	server := grpc.NewServer()
	pb.RegisterServiceRadarServerServer(server, &grpc_entrypoint.ServiceRadarServerServer{})
	reflection.Register(server)

	lis, err := net.Listen("tcp", ":"+PORT)
	if err != nil {
		log.Fatalf("net.Listen err: %v", err)
	}
	log.Print("Server start (enable reflection) :" + PORT)
	server.Serve(lis)
}
