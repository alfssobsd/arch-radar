package grpc_entrypoints

import (
	"context"
	pb "github.com/alfssobsd/arch-radar/archradar-executor/entrypoints/grpc_entrypoints/gen"
	"github.com/alfssobsd/arch-radar/archradar-executor/usecases"
	"github.com/golang/protobuf/ptypes/empty"
	"github.com/google/uuid"
	"log"
)

type dictionariesGrpcServiceServer struct {
	dictionariesManageUseCase usecases.DictionariesManageUseCase
}

func NewDictionariesGrpcServiceServer(dictionariesManageUseCase usecases.DictionariesManageUseCase) *dictionariesGrpcServiceServer {
	return &dictionariesGrpcServiceServer{dictionariesManageUseCase}
}

func (s *dictionariesGrpcServiceServer) GetDictionaryList(ctx context.Context, req *pb.DictionaryListRequest) (*pb.DictionaryListResponse, error) {
	log.Println(req.DictType.String())

	useCaseResult, _ := s.dictionariesManageUseCase.ShowAreaList()
	grpcResponse := &pb.DictionaryListResponse{
		Items: []*pb.DictionaryResponse{},
		PageInfo: &pb.PageInfoItem{
			PageLimit:  0,
			PageSize:   0,
			PageNumber: 0,
			TotalItems: 0,
		},
	}

	for _, v := range useCaseResult.Items {
		grpcResponse.Items = append(grpcResponse.Items, &pb.DictionaryResponse{
			DictionaryUuid: v.AreaUUID.String(),
			DictType:       pb.DictionaryType_AREA,
			Title:          v.Title,
			Color:          v.Color,
			Description:    "NO TEXT",
		})
	}

	return grpcResponse, nil
}

// Create Area Item
func (s *dictionariesGrpcServiceServer) CreateDictionaryItem(ctx context.Context, req *pb.DictionaryCreateRequest) (*pb.DictionaryResponse, error) {
	s.dictionariesManageUseCase.CreateArea(usecases.CreateAreaInDTO{
		AreaUUID:    uuid.MustParse(req.DictionaryUuid),
		Title:       req.Title,
		Color:       req.Color,
		Description: &req.Description,
	})
	return &pb.DictionaryResponse{
		DictionaryUuid: "",
		DictType:       0,
		Title:          "",
		Color:          "",
		Description:    "",
	}, nil
}

// Delete Area Item
func (s *dictionariesGrpcServiceServer) DeleteDictionaryItem(ctx context.Context, req *pb.DictionaryDeleteRequest) (*empty.Empty, error) {
	return nil, nil
}

// Update Area Item
func (s *dictionariesGrpcServiceServer) UpdateDictionaryItem(ctx context.Context, req *pb.DictionaryUpdateRequest) (*pb.DictionaryResponse, error) {
	return nil, nil
}