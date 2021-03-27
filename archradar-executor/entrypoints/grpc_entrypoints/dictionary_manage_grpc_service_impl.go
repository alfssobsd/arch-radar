package grpc_entrypoints

import (
	"context"
	"github.com/alfssobsd/arch-radar/archradar-executor/usecases"
	pb "github.com/alfssobsd/arch-radar/archradar-grpc"
	"github.com/golang/protobuf/ptypes/empty"
	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log"
)

type dictionaryManageGrpcServiceServer struct {
	dictionariesManageUseCase usecases.DictionariesManageUseCase
}

func NewDictionaryManageGrpcServiceServer(dictionariesManageUseCase usecases.DictionariesManageUseCase) *dictionaryManageGrpcServiceServer {
	return &dictionaryManageGrpcServiceServer{dictionariesManageUseCase}
}

func (s *dictionaryManageGrpcServiceServer) GetDictionaryList(ctx context.Context, req *pb.DictionaryListRequest) (*pb.DictionaryListResponse, error) {
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
func (s *dictionaryManageGrpcServiceServer) CreateDictionaryItem(ctx context.Context, req *pb.DictionaryCreateRequest) (*pb.DictionaryResponse, error) {
	err := req.Validate()

	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

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
func (s *dictionaryManageGrpcServiceServer) DeleteDictionaryItem(ctx context.Context, req *pb.DictionaryDeleteRequest) (*empty.Empty, error) {
	return nil, nil
}

// Update Area Item
func (s *dictionaryManageGrpcServiceServer) UpdateDictionaryItem(ctx context.Context, req *pb.DictionaryUpdateRequest) (*pb.DictionaryResponse, error) {
	return nil, nil
}
