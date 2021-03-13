package grpc_entrypoint

import (
	"arch-radar/backend-service/backend-service/dataproviders/pg_provider"
	pb "arch-radar/backend-service/backend-service/entrypoints/grpc_entrypoint/gen"
	"context"
	"github.com/google/uuid"
	"log"
)

func NewServiceRadarGrpcServiceServer(serviceRepo pg_provider.ServiceRepo) *serviceRadarGrpcServiceServer {
	return &serviceRadarGrpcServiceServer{serviceRepo}
}

type serviceRadarGrpcServiceServer struct {
	serviceRepo pg_provider.ServiceRepo
}

func (s *serviceRadarGrpcServiceServer) GetServices(ctx context.Context, r *pb.ListServicesRequest) (*pb.ListServicesResponse, error) {
	log.Print("ServiceRadarServerServer.GetServices request -> " + r.String())
	result, _ := s.serviceRepo.List()
	var responses []*pb.ServiceResponse
	for _, v := range result {
		uuid, _ := uuid.NewUUID()
		responses = append(responses, &pb.ServiceResponse{
			Uuid:  uuid.String(),
			Title: v,
		})
	}
	items := pb.ListServicesResponse{
		Items: responses,
	}
	return &items, nil
}
