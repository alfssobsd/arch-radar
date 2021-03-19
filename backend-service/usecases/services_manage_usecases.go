package usecases

import (
	"arch-radar/backend-service/backend-service/dataproviders/pg_provider"
)

type ServicesManageUseCase interface {
	CreateService() error
	ListByPage(pageNumber int) ([]ServiceItemDto, error)
}
type servicesManageUseCase struct {
	serviceRepo pg_provider.ServiceRepo
}

func NewServiceManageUseCase(serviceRepo pg_provider.ServiceRepo) *servicesManageUseCase {
	return &servicesManageUseCase{serviceRepo: serviceRepo}
}

func (uc *servicesManageUseCase) CreateService() error {
	return nil
}

func (uc *servicesManageUseCase) ListByPage(pageNumber int) ([]ServiceItemDto, error) {
	serviceList, err := uc.serviceRepo.List()
	if err != nil {
		return nil, err
	}

	var resultList []ServiceItemDto

	for _, service := range serviceList {
		resultList = append(resultList, ServiceItemDto{
			UUID:  service.ID,
			Title: service.Title,
		})
	}
	return resultList, nil
}
