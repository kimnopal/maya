package converter

import (
	"github.com/kimnopal/maya/entity"
	"github.com/kimnopal/maya/model"
)

func TagEntityToResponse(tag *entity.Tag) *model.TagResponse {
	return &model.TagResponse{
		ID:   tag.ID,
		Name: tag.Name,
	}
}

func TagCreateRequestToEntity(createRequest *model.TagCreateRequest) *entity.Tag {
	return &entity.Tag{
		Name: createRequest.Name,
	}
}

func TagUpdateRequestToEntity(tag *entity.Tag, updateRequest *model.TagUpdateRequest) *entity.Tag {
	tag.ID = updateRequest.ID
	tag.Name = updateRequest.Name

	return tag
}

func TagListEntityToResponse(tags *[]*entity.Tag) []*model.TagResponse {
	var responses []*model.TagResponse

	for _, tag := range *tags {
		response := &model.TagResponse{
			ID:   tag.ID,
			Name: tag.Name,
		}
		responses = append(responses, response)
	}

	return responses
}
