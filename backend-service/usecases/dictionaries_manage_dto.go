package usecases

import "github.com/google/uuid"

type ShowAreaListReplyDTO struct {
	Items []AreaItemDTO
}
type AreaItemDTO struct {
	AreaUUID    uuid.UUID
	Title       string
	Color       string
	Description *string
}
