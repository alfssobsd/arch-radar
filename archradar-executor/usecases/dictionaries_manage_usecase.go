package usecases

import (
	"github.com/alfssobsd/arch-radar/archradar-executor/dataproviders/pg_provider"
	"github.com/alfssobsd/arch-radar/archradar-executor/dataproviders/pg_provider/model"
)

type DictionariesManageUseCase interface {
	CreateArea(CreateAreaInDTO) (CreateAreaReplyDTO, error)
	ShowAreaList() (ShowAreaListReplyDTO, error)
}

type dictionaryManageUseCase struct {
	dictionaryRepo pg_provider.DictionaryRepo
}

func NewDictionaryManageUseCase(dictionaryRepo pg_provider.DictionaryRepo) *dictionaryManageUseCase {
	return &dictionaryManageUseCase{dictionaryRepo}
}
func (uc *dictionaryManageUseCase) CreateArea(dto CreateAreaInDTO) (CreateAreaReplyDTO, error) {
	err := uc.dictionaryRepo.Create(&model.Dictionary{
		ID:          dto.AreaUUID,
		Title:       dto.Title,
		Color:       dto.Color,
		DictType:    "AREA",
		Description: dto.Description,
	})

	return CreateAreaReplyDTO{}, err
}

func (uc *dictionaryManageUseCase) ShowAreaList() (ShowAreaListReplyDTO, error) {
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
