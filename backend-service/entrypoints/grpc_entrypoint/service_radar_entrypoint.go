package grpc_entrypoint

import (
	pb "arch-radar/backend-service/backend-service/entrypoints/grpc_entrypoint/gen"
	"context"
	"log"
)

type ServiceRadarServerServer struct {
	pb.UnimplementedServiceRadarServerServer
}

func (s *ServiceRadarServerServer) GetServices(ctx context.Context, r *pb.ListServicesRequest) (*pb.ListServicesResponse, error) {
	log.Print("ServiceRadarServerServer.GetServices request -> " + r.String())
	items := pb.ListServicesResponse{
		Items: nil,
	}
	return &items, nil
}
