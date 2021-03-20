package pg_provider

import (
	"arch-radar/backend-service/backend-service/dataproviders/pg_provider/model"
	"github.com/go-pg/pg/v10"
)

func NewDictionaryRepo(db *pg.DB) *dictionaryRepo {
	return &dictionaryRepo{db: db}
}

type dictionaryRepo struct {
	db *pg.DB
}

type DictionaryRepo interface {
	List() ([]model.Dictionary, error)
}

func (r *dictionaryRepo) List() ([]model.Dictionary, error) {
	var dictionaryList []model.Dictionary
	err := r.db.Model(&dictionaryList).Select()
	if err != nil {
		return nil, err
	}

	return dictionaryList, err
}
