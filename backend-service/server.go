package main

import (
	"arch-radar/backend-service/backend-service/dataproviders/pg_provider"
	"arch-radar/backend-service/backend-service/entrypoints/grpc_entrypoint"
	pb "arch-radar/backend-service/backend-service/entrypoints/grpc_entrypoint/gen"
	"arch-radar/backend-service/backend-service/usecases"
	"github.com/go-pg/pg/v10"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
)

const PORT = "9090"

func main() {

	db := pg.Connect(&pg.Options{
		User:     "archradar_user",
		Password: "archradar_pass",
		Database: "archradar_db",
	})
	defer db.Close()

	server := grpc.NewServer()
	pb.RegisterServiceRadarGrpcServiceServer(server, grpc_entrypoint.NewServiceRadarGrpcServiceServer(
		usecases.NewServiceManageUseCase(pg_provider.NewServiceRepo(db))),
	)
	reflection.Register(server)

	lis, err := net.Listen("tcp", ":"+PORT)
	if err != nil {
		log.Fatalf("net.Listen err: %v", err)
	}
	log.Print("Server start (enable reflection) :" + PORT)
	server.Serve(lis)
}
