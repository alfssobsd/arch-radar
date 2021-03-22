package usecases

import "github.com/google/uuid"

type ServiceItemDTO struct {
	UUID  uuid.UUID
	Title string
}
