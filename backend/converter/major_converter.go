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
	}
}

func MajorCreateRequestToEntity(createRequest *model.MajorCreateRequest) *entity.Major {
	return &entity.Major{
		Name:      createRequest.Name,
		FacultyID: createRequest.FacultyID,
	}
}

func MajorUpdateRequestToEntity(updateRequest *model.FacultyUpdateRequest) *entity.Major {
	return &entity.Major{
		ID:        updateRequest.ID,
		Name:      updateRequest.Name,
		FacultyID: updateRequest.ID,
	}
}

func MajorListEntityToResponse(entities *[]*entity.Major) []*model.MajorResponse {
	var responses []*model.MajorResponse

	for _, entity := range *entities {
		response := &model.MajorResponse{
			ID:        entity.ID,
			Name:      entity.Name,
			CreatedAt: entity.CreatedAt,
			UpdatedAt: entity.CreatedAt,
		}
		responses = append(responses, response)
	}

	return responses
}
