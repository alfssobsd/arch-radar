package usecases

import "github.com/google/uuid"

type CreateAreaInDTO struct {
	AreaUUID    uuid.UUID
	Title       string
	Color       string
	Description *string
}
type CreateAreaReplyDTO struct {
}
type ShowAreaListReplyDTO struct {
	Items []AreaItemDTO
}
type AreaItemDTO struct {
	AreaUUID    uuid.UUID
	Title       string
	Color       string
	Description *string
}
