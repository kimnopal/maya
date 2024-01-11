package converter

import (
	"github.com/kimnopal/maya/entity"
	"github.com/kimnopal/maya/model"
)

func MajorEntityToResponse(entity *entity.Major) *model.MajorResponse {
	return &model.MajorResponse{
		ID:        entity.ID,
		Name:      entity.Name,
		CreatedAt: entity.CreatedAt,
		UpdatedAt: entity.CreatedAt,
		Faculty:   entity.Faculty,
	}
}
