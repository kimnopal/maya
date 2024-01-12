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
		UpdatedAt: entity.UpdatedAt,
	}
}

func MajorCreateRequestToEntity(createRequest *model.MajorCreateRequest) *entity.Major {
	return &entity.Major{
		Name:      createRequest.Name,
		FacultyID: createRequest.FacultyID,
	}
}

func MajorUpdateRequestToEntity(major *entity.Major, updateRequest *model.MajorUpdateRequest) *entity.Major {
	major.ID = updateRequest.ID
	major.Name = updateRequest.Name
	major.FacultyID = updateRequest.FacultyID

	return major
}

func MajorListEntityToResponse(entities *[]*entity.Major) []*model.MajorResponse {
	var responses []*model.MajorResponse

	for _, entity := range *entities {
		response := &model.MajorResponse{
			ID:        entity.ID,
			Name:      entity.Name,
			CreatedAt: entity.CreatedAt,
			UpdatedAt: entity.UpdatedAt,
		}
		responses = append(responses, response)
	}

	return responses
}
