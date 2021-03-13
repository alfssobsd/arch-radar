//nolint
//lint:file-ignore U1000 ignore unused code, it's generated
package model

var Columns = struct {
	Dictionary struct {
		ID, Title, Color, DictType, Description string
	}
	Service struct {
		ID, Title, Description string
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
	ServicesDictionary: struct {
		Name, Alias string
	}{
		Name:  "services_dictionaries",
		Alias: "t",
	},
}

type Dictionary struct {
	tableName struct{} `pg:"dictionaries,alias:t,,discard_unknown_columns"`

	ID          string  `pg:"dictionary_uuid,pk,type:uuid"`
	Title       string  `pg:"title,use_zero"`
	Color       string  `pg:"color,use_zero"`
	DictType    string  `pg:"dict_type,use_zero"`
	Description *string `pg:"description"`
}

type Service struct {
	tableName struct{} `pg:"services,alias:t,,discard_unknown_columns"`

	ID          string  `pg:"service_uuid,pk,type:uuid"`
	Title       string  `pg:"title,use_zero"`
	Description *string `pg:"description"`
}

type ServicesDictionary struct {
	tableName struct{} `pg:"services_dictionaries,alias:t,,discard_unknown_columns"`

	ServiceUuid    string `pg:"service_uuid,type:uuid,use_zero"`
	DictionaryUuid string `pg:"dictionary_uuid,type:uuid,use_zero"`

	DictionaryUuidRel *Dictionary `pg:"fk:dictionary_uuid"`
	ServiceUuidRel    *Service    `pg:"fk:service_uuid"`
}
