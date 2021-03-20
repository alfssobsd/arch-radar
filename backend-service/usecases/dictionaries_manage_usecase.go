package usecases

import (
	"arch-radar/backend-service/backend-service/dataproviders/pg_provider"
)

type DictionariesManageUseCase interface {
	ShowAreaList() (ShowAreaListReplyDTO, error)
}
type dictionariesManageUseCase struct {
	dictionaryRepo pg_provider.DictionaryRepo
}

func NewDictionariesManageUseCase(dictionaryRepo pg_provider.DictionaryRepo) *dictionariesManageUseCase {
	return &dictionariesManageUseCase{dictionaryRepo}
}

func (uc *dictionariesManageUseCase) ShowAreaList() (ShowAreaListReplyDTO, error) {
	dictList, err := uc.dictionaryRepo.List()

	var areaList []AreaItemDTO

	for _, v := range dictList {
		areaList = append(areaList, AreaItemDTO{
			AreaUUID:    v.ID,
			Title:       v.Title,
			Color:       v.Color,
			Description: v.Description,
		})
	}

	return ShowAreaListReplyDTO{Items: areaList}, err
}
