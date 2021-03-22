//nolint
//lint:file-ignore U1000 ignore unused code, it's generated
package model

import (
	"github.com/google/uuid"
)

var Columns = struct {
	Dictionary struct {
		ID, Title, Color, DictType, Description string
	}
	Service struct {
		ID, Title, Description string
	}
	ServicesConnection struct {
		ServiceAUuid, ServiceBUuid, DirectionType, Description string

		ServiceAUuidRel, ServiceBUuidRel string
	}
	ServicesDictionary struct {
		ServiceUuid, DictionaryUuid string

		DictionaryUuidRel, ServiceUuidRel string
	}
}{
	Dictionary: struct {
		ID, Title, Color, DictType, Description string
	}{
		ID:          "dictionary_uuid",
		Title:       "title",
		Color:       "color",
		DictType:    "dict_type",
		Description: "description",
	},
	Service: struct {
		ID, Title, Description string
	}{
		ID:          "service_uuid",
		Title:       "title",
		Description: "description",
	},
	ServicesConnection: struct {
		ServiceAUuid, ServiceBUuid, DirectionType, Description string

		ServiceAUuidRel, ServiceBUuidRel string
	}{
		ServiceAUuid:  "service_a_uuid",
		ServiceBUuid:  "service_b_uuid",
		DirectionType: "direction_type",
		Description:   "description",

		ServiceAUuidRel: "ServiceAUuidRel",
		ServiceBUuidRel: "ServiceBUuidRel",
	},
	ServicesDictionary: struct {
		ServiceUuid, DictionaryUuid string

		DictionaryUuidRel, ServiceUuidRel string
	}{
		ServiceUuid:    "service_uuid",
		DictionaryUuid: "dictionary_uuid",

		DictionaryUuidRel: "DictionaryUuidRel",
		ServiceUuidRel:    "ServiceUuidRel",
	},
}

var Tables = struct {
	Dictionary struct {
		Name, Alias string
	}
	Service struct {
		Name, Alias string
	}
	ServicesConnection struct {
		Name, Alias string
	}
	ServicesDictionary struct {
		Name, Alias string
	}
}{
	Dictionary: struct {
		Name, Alias string
	}{
		Name:  "dictionaries",
		Alias: "t",
	},
	Service: struct {
		Name, Alias string
	}{
		Name:  "services",
		Alias: "t",
	},
	ServicesConnection: struct {
		Name, Alias string
	}{
		Name:  "services_connections",
		Alias: "t",
	},
	ServicesDictionary: struct {
		Name, Alias string
	}{
		Name:  "services_dictionaries",
		Alias: "t",
	},
}

type Dictionary struct {
	tableName struct{} `pg:"dictionaries,alias:t,discard_unknown_columns"`

	ID          uuid.UUID `pg:"dictionary_uuid,pk,type:uuid"`
	Title       string    `pg:"title,use_zero"`
	Color       string    `pg:"color,use_zero"`
	DictType    string    `pg:"dict_type,use_zero"`
	Description *string   `pg:"description"`
}

type Service struct {
	tableName struct{} `pg:"services,alias:t,discard_unknown_columns"`

	ID          uuid.UUID `pg:"service_uuid,pk,type:uuid"`
	Title       string    `pg:"title,use_zero"`
	Description *string   `pg:"description"`
}

type ServicesConnection struct {
	tableName struct{} `pg:"services_connections,alias:t,discard_unknown_columns"`

	ServiceAUuid  uuid.UUID `pg:"service_a_uuid,type:uuid,use_zero"`
	ServiceBUuid  uuid.UUID `pg:"service_b_uuid,type:uuid,use_zero"`
	DirectionType string    `pg:"direction_type,use_zero"`
	Description   string    `pg:"description,use_zero"`

	ServiceAUuidRel *Service `pg:"fk:service_a_uuid"`
	ServiceBUuidRel *Service `pg:"fk:service_b_uuid"`
}

type ServicesDictionary struct {
	tableName struct{} `pg:"services_dictionaries,alias:t,discard_unknown_columns"`

	ServiceUuid    uuid.UUID `pg:"service_uuid,type:uuid,use_zero"`
	DictionaryUuid uuid.UUID `pg:"dictionary_uuid,type:uuid,use_zero"`

	DictionaryUuidRel *Dictionary `pg:"fk:dictionary_uuid"`
	ServiceUuidRel    *Service    `pg:"fk:service_uuid"`
}
