package pg_provider

import (
	"arch-radar/backend-service/backend-service/dataproviders/pg_provider/model"
	"github.com/go-pg/pg/v9"
)

func NewServiceRepo(db *pg.DB) *serviceRepo {
	return &serviceRepo{db: db}
}

type serviceRepo struct {
	db *pg.DB
}

type ServiceRepo interface {
	List() ([]string, error)
}

func (r *serviceRepo) List() ([]string, error) {
	var services []model.Service
	err := r.db.Model(&services).Select()
	if err != nil {
		return nil, err
	}

	var resultList []string
	for _, service := range services {
		resultList = append(resultList, service.Title)
	}

	return resultList, err
}
