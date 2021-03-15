CREATE TABLE dictionaries
(
    "dictionary_uuid" uuid    NOT NULL,
    "title"           varchar NOT NULL,
    "color"           varchar NOT NULL,
    "dict_type"       varchar NOT NULL, -- LANG,AREAS,TECH_LAYER,DATABASE,TEAM,LOCATION
    "description"     text NULL,
    primary key (dictionary_uuid)
);
CREATE INDEX services_title_dict_type_unq
    ON dictionaries(title, dict_type);
CREATE INDEX services_dictionary_uuid_dict_type_idx
    ON dictionaries(dictionary_uuid, dict_type);

CREATE TABLE "services"
(
    "service_uuid" uuid    NOT NULL,
    "title"        varchar NOT NULL,
    "description"  text NULL,
    primary key ("service_uuid")
);
CREATE INDEX services_title_unq
    ON services(title);

CREATE TABLE "services_dictionaries"
(
    "service_uuid"    uuid    NOT NULL references "services",
    "dictionary_uuid" uuid    NOT NULL references "dictionaries"
);

CREATE INDEX services_dictionaries_service_idx
    ON services_dictionaries(service_uuid);
CREATE UNIQUE INDEX services_dictionaries_unq
    ON services_dictionaries(service_uuid,dictionary_uuid);


CREATE TABLE "services_connections" (
    "service_a_uuid" uuid NOT NULL references "services",
    "service_b_uuid" uuid NOT NULL references "services",
    "direction_type" varchar NOT NULL, -- IN, OUT, INOUT
    "description" varchar NOT NULL
);

CREATE INDEX services_connections_unq
    ON services_connections(service_a_uuid,service_b_uuid);
CREATE INDEX services_connections_a_idx
    ON services_connections(service_a_uuid);
CREATE INDEX services_connections_b_idx
    ON services_connections(service_b_uuid);