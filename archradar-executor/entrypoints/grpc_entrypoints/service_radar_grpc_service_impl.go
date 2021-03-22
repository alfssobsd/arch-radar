package grpc_entrypoints

import (
	"context"
	pb "github.com/alfssobsd/arch-radar/archradar-executor/entrypoints/grpc_entrypoints/gen"
	"github.com/alfssobsd/arch-radar/archradar-executor/usecases"
	"log"
)

type serviceRadarGrpcServiceServer struct {
	servicesManageUseCase usecases.ServicesManageUseCase
}

func NewServiceRadarGrpcServiceServer(servicesUseCase usecases.ServicesManageUseCase) *serviceRadarGrpcServiceServer {
	return &serviceRadarGrpcServiceServer{servicesUseCase}
}

func (s *serviceRadarGrpcServiceServer) GetServices(ctx context.Context, r *pb.ListServicesRequest) (*pb.ListServicesResponse, error) {
	log.Print("ServiceRadarServerServer.GetServices request -> " + r.String())

	result, _ := s.servicesManageUseCase.ListByPage(1)
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
