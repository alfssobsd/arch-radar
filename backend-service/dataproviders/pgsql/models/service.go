package models

import (
	"github.com/satori/go.uuid"
)

type ServiceModelDb struct {
	ID    uuid.UUID `sql:",type:uuid"`
	Title string
}
