package converter

import (
	"github.com/kimnopal/maya/entity"
	"github.com/kimnopal/maya/model"
)

func PostCategoryEntityToResponse(postCategory *entity.PostCategory) *model.PostCategoryResponse {
	return &model.PostCategoryResponse{
		ID:   postCategory.ID,
		Name: postCategory.Name,
	}
}

func PostCategoryCreateRequestToEntity(createRequest *model.PostCategoryCreateRequest) *entity.PostCategory {
	return &entity.PostCategory{
		Name: createRequest.Name,
	}
}

func PostCategoryUpdateRequestToEntity(postCategory *entity.PostCategory, updateRequest *model.PostCategoryUpdateRequest) *entity.PostCategory {
	postCategory.ID = updateRequest.ID
	postCategory.Name = updateRequest.Name

	return postCategory
}
