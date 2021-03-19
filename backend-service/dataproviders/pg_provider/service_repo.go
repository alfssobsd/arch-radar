package pg_provider

import (
	"arch-radar/backend-service/backend-service/dataproviders/pg_provider/model"
	"github.com/go-pg/pg/v10"
)

func NewServiceRepo(db *pg.DB) *serviceRepo {
	return &serviceRepo{db: db}
}

type serviceRepo struct {
	db *pg.DB
}

type ServiceRepo interface {
	Update(model.Service) error
	Create(model.Service) error
	List() ([]model.Service, error)
}

func (r *serviceRepo) List() ([]model.Service, error) {
	var services []model.Service
	err := r.db.Model(&services).Select()
	if err != nil {
		return nil, err
	}

	return services, err
}

func (r *serviceRepo) Create(model.Service) error {

	return nil
}

func (r *serviceRepo) Update(model.Service) error {

	return nil
}
