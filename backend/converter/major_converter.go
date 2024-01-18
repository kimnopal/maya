package converter

import (
	"github.com/kimnopal/maya/entity"
	"github.com/kimnopal/maya/model"
)

func MajorEntityToResponse(entity *entity.Major) *model.MajorResponse {
	var facultyResponse *model.FacultyResponse
	if entity.Faculty != nil {
		facultyResponse = FacultyEntityToResponse(entity.Faculty)
	}

	return &model.MajorResponse{
		ID:        entity.ID,
		Name:      entity.Name,
		CreatedAt: entity.CreatedAt,
		UpdatedAt: entity.UpdatedAt,
		Faculty:   facultyResponse,
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

func MajorListEntityToResponse(majors *[]*entity.Major) []*model.MajorResponse {
	var responses []*model.MajorResponse

	for _, major := range *majors {
		response := &model.MajorResponse{
			ID:        major.ID,
			Name:      major.Name,
			CreatedAt: major.CreatedAt,
			UpdatedAt: major.UpdatedAt,
		}
		responses = append(responses, response)
	}

	return responses
}
