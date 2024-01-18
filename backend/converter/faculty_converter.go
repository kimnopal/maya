package converter

import (
	"github.com/kimnopal/maya/entity"
	"github.com/kimnopal/maya/model"
)

func FacultyEntityToResponse(faculty *entity.Faculty) *model.FacultyResponse {
	var majors []*model.MajorResponse
	if faculty.Majors != nil {
		majors = MajorListEntityToResponse(&faculty.Majors)
	}

	return &model.FacultyResponse{
		ID:        faculty.ID,
		Name:      faculty.Name,
		CreatedAt: faculty.CreatedAt,
		UpdatedAt: faculty.UpdatedAt,
		Majors:    majors,
	}
}

func FacultyCreateRequestToEntity(createRequest *model.FacultyCreateRequest) *entity.Faculty {
	return &entity.Faculty{
		Name: createRequest.Name,
	}
}

func FacultyUpdateRequestToEntity(faculty *entity.Faculty, updateRequest *model.FacultyUpdateRequest) *entity.Faculty {
	faculty.Name = updateRequest.Name

	return faculty
}

func FacultyListEntityToResponse(entities *[]*entity.Faculty) []*model.FacultyResponse {
	var responses []*model.FacultyResponse
	for _, entity := range *entities {
		response := &model.FacultyResponse{
			ID:        entity.ID,
			Name:      entity.Name,
			CreatedAt: entity.CreatedAt,
			UpdatedAt: entity.UpdatedAt,
		}
		responses = append(responses, response)
	}

	return responses
}
