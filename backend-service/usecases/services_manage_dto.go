package usecases

import "github.com/google/uuid"

type ServiceItemDto struct {
	UUID  uuid.UUID
	Title string
}
