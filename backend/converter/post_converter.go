package converter

import (
	"github.com/kimnopal/maya/entity"
	"github.com/kimnopal/maya/model"
)

func PostEntityToResponse(post *entity.Post) *model.PostResponse {
	return &model.PostResponse{
		ID:          post.ID,
		Title:       post.Title,
		Code:        post.Code,
		Description: post.Description,
		CreatedAt:   post.CreatedAt,
		UpdatedAt:   post.UpdatedAt,
	}
}

func PostCreateRequestToEntity(createRequest *model.PostCreateRequest) *entity.Post {
	return &entity.Post{
		Title:          createRequest.Title,
		Code:           createRequest.Code,
		Description:    createRequest.Description,
		UserID:         createRequest.UserID,
		PostCategoryID: createRequest.PostCategoryID,
	}
}

func PostUpdateRequestToEntity(post *entity.Post, updateRequest *model.PostUpdateRequest) *entity.Post {
	post.Title = updateRequest.Title
	post.Description = updateRequest.Description
	post.PostCategoryID = updateRequest.PostCategoryID

	return post
}
