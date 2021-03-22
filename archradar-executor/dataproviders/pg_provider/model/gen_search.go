//nolint
//lint:file-ignore U1000 ignore unused code, it's generated
package model

import (
	"github.com/google/uuid"

	"github.com/go-pg/pg/v10"
	"github.com/go-pg/pg/v10/orm"
)

const condition = "?.? = ?"

// base filters
type applier func(query *orm.Query) (*orm.Query, error)

type search struct {
	appliers []applier
}

func (s *search) apply(query *orm.Query) {
	for _, applier := range s.appliers {
		query.Apply(applier)
	}
}

func (s *search) where(query *orm.Query, table, field string, value interface{}) {

	query.Where(condition, pg.Ident(table), pg.Ident(field), value)

}

func (s *search) WithApply(a applier) {
	if s.appliers == nil {
		s.appliers = []applier{}
	}
	s.appliers = append(s.appliers, a)
}

func (s *search) With(condition string, params ...interface{}) {
	s.WithApply(func(query *orm.Query) (*orm.Query, error) {
		return query.Where(condition, params...), nil
	})
}

// Searcher is interface for every generated filter
type Searcher interface {
	Apply(query *orm.Query) *orm.Query
	Q() applier

	With(condition string, params ...interface{})
	WithApply(a applier)
}

type DictionarySearch struct {
	search

	ID          *uuid.UUID
	Title       *string
	Color       *string
	DictType    *string
	Description *string
}

func (s *DictionarySearch) Apply(query *orm.Query) *orm.Query {
	if s.ID != nil {
		s.where(query, Tables.Dictionary.Alias, Columns.Dictionary.ID, s.ID)
	}
	if s.Title != nil {
		s.where(query, Tables.Dictionary.Alias, Columns.Dictionary.Title, s.Title)
	}
	if s.Color != nil {
		s.where(query, Tables.Dictionary.Alias, Columns.Dictionary.Color, s.Color)
	}
	if s.DictType != nil {
		s.where(query, Tables.Dictionary.Alias, Columns.Dictionary.DictType, s.DictType)
	}
	if s.Description != nil {
		s.where(query, Tables.Dictionary.Alias, Columns.Dictionary.Description, s.Description)
	}

	s.apply(query)

	return query
}

func (s *DictionarySearch) Q() applier {
	return func(query *orm.Query) (*orm.Query, error) {
		return s.Apply(query), nil
	}
}

type ServiceSearch struct {
	search

	ID          *uuid.UUID
	Title       *string
	Description *string
}

func (s *ServiceSearch) Apply(query *orm.Query) *orm.Query {
	if s.ID != nil {
		s.where(query, Tables.Service.Alias, Columns.Service.ID, s.ID)
	}
	if s.Title != nil {
		s.where(query, Tables.Service.Alias, Columns.Service.Title, s.Title)
	}
	if s.Description != nil {
		s.where(query, Tables.Service.Alias, Columns.Service.Description, s.Description)
	}

	s.apply(query)

	return query
}

func (s *ServiceSearch) Q() applier {
	return func(query *orm.Query) (*orm.Query, error) {
		return s.Apply(query), nil
	}
}

type ServicesConnectionSearch struct {
	search

	ServiceAUuid  *uuid.UUID
	ServiceBUuid  *uuid.UUID
	DirectionType *string
	Description   *string
}

func (s *ServicesConnectionSearch) Apply(query *orm.Query) *orm.Query {
	if s.ServiceAUuid != nil {
		s.where(query, Tables.ServicesConnection.Alias, Columns.ServicesConnection.ServiceAUuid, s.ServiceAUuid)
	}
	if s.ServiceBUuid != nil {
		s.where(query, Tables.ServicesConnection.Alias, Columns.ServicesConnection.ServiceBUuid, s.ServiceBUuid)
	}
	if s.DirectionType != nil {
		s.where(query, Tables.ServicesConnection.Alias, Columns.ServicesConnection.DirectionType, s.DirectionType)
	}
	if s.Description != nil {
		s.where(query, Tables.ServicesConnection.Alias, Columns.ServicesConnection.Description, s.Description)
	}

	s.apply(query)

	return query
}

func (s *ServicesConnectionSearch) Q() applier {
	return func(query *orm.Query) (*orm.Query, error) {
		return s.Apply(query), nil
	}
}

type ServicesDictionarySearch struct {
	search

	ServiceUuid    *uuid.UUID
	DictionaryUuid *uuid.UUID
}

func (s *ServicesDictionarySearch) Apply(query *orm.Query) *orm.Query {
	if s.ServiceUuid != nil {
		s.where(query, Tables.ServicesDictionary.Alias, Columns.ServicesDictionary.ServiceUuid, s.ServiceUuid)
	}
	if s.DictionaryUuid != nil {
		s.where(query, Tables.ServicesDictionary.Alias, Columns.ServicesDictionary.DictionaryUuid, s.DictionaryUuid)
	}

	s.apply(query)

	return query
}

func (s *ServicesDictionarySearch) Q() applier {
	return func(query *orm.Query) (*orm.Query, error) {
		return s.Apply(query), nil
	}
}
