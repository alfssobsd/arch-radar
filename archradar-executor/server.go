package main

import (
	"github.com/alfssobsd/arch-radar/archradar-executor/dataproviders/pg_provider"
	"github.com/alfssobsd/arch-radar/archradar-executor/entrypoints/grpc_entrypoints"
	"github.com/alfssobsd/arch-radar/archradar-executor/usecases"
	pb "github.com/alfssobsd/arch-radar/archradar-grpc"
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
	pb.RegisterServiceManageGrpcServiceServer(server, grpc_entrypoints.NewServiceManageGrpcService(
		usecases.NewServiceManageUseCase(pg_provider.NewServiceRepo(db))),
	)
	pb.RegisterDictionaryManageGrpcServiceServer(server, grpc_entrypoints.NewDictionaryManageGrpcServiceServer(
		usecases.NewDictionaryManageUseCase(pg_provider.NewDictionaryRepo(db))),
	)
	reflection.Register(server)

	lis, err := net.Listen("tcp", ":"+PORT)
	if err != nil {
		log.Fatalf("net.Listen err: %v", err)
	}
	log.Print("Server start (enable reflection) :" + PORT)
	server.Serve(lis)
}
