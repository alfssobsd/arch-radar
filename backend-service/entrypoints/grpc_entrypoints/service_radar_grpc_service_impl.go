package grpc_entrypoints

import (
	pb "arch-radar/backend-service/backend-service/entrypoints/grpc_entrypoints/gen"
	"arch-radar/backend-service/backend-service/usecases"
	"context"
	"log"
)

func NewServiceRadarGrpcServiceServer(servicesUseCase usecases.ServicesManageUseCase) *serviceRadarGrpcServiceServer {
	return &serviceRadarGrpcServiceServer{servicesUseCase}
}

type serviceRadarGrpcServiceServer struct {
	servicesUseCase usecases.ServicesManageUseCase
}

func (s *serviceRadarGrpcServiceServer) GetServices(ctx context.Context, r *pb.ListServicesRequest) (*pb.ListServicesResponse, error) {
	log.Print("ServiceRadarServerServer.GetServices request -> " + r.String())

	result, _ := s.servicesUseCase.ListByPage(1)
	var responses []*pb.ServiceResponse
	for _, v := range result {
		responses = append(responses, &pb.ServiceResponse{
			ServiceUUID: &pb.UUIDOptional{Value: v.UUID.String()},
			Title:       v.Title,
		})
	}
	items := pb.ListServicesResponse{
		Items: responses,
	}

	return &items, nil
}