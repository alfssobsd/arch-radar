package grpc_entrypoints

import (
	"context"
	"github.com/alfssobsd/arch-radar/archradar-executor/usecases"
	pb "github.com/alfssobsd/arch-radar/archradar-grpc"
	"log"
)

type serviceManageGrpcServiceServer struct {
	servicesManageUseCase usecases.ServicesManageUseCase
}

func NewServiceManageGrpcService(servicesUseCase usecases.ServicesManageUseCase) *serviceManageGrpcServiceServer {
	return &serviceManageGrpcServiceServer{servicesUseCase}
}

func (s *serviceManageGrpcServiceServer) GetServices(ctx context.Context, r *pb.ListServicesRequest) (*pb.ListServicesResponse, error) {
	log.Print("ServiceServerServer.GetServices request -> " + r.String())

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
