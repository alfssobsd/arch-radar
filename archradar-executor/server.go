package main

import (
	"github.com/alfssobsd/arch-radar/archradar-executor/dataproviders/pg_provider"
	"github.com/alfssobsd/arch-radar/archradar-executor/entrypoints/grpc_entrypoints"
	pb "github.com/alfssobsd/arch-radar/archradar-executor/entrypoints/grpc_entrypoints/gen"
	"github.com/alfssobsd/arch-radar/archradar-executor/usecases"
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
	pb.RegisterServiceRadarGrpcServiceServer(server, grpc_entrypoints.NewServiceRadarGrpcServiceServer(
		usecases.NewServiceManageUseCase(pg_provider.NewServiceRepo(db))),
	)
	pb.RegisterDictionariesGrpcServiceServer(server, grpc_entrypoints.NewDictionariesGrpcServiceServer(
		usecases.NewDictionariesManageUseCase(pg_provider.NewDictionaryRepo(db))),
	)
	reflection.Register(server)

	lis, err := net.Listen("tcp", ":"+PORT)
	if err != nil {
		log.Fatalf("net.Listen err: %v", err)
	}
	log.Print("Server start (enable reflection) :" + PORT)
	server.Serve(lis)
}
