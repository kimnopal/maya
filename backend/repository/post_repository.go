package repository

import "github.com/kimnopal/maya/entity"

type PostRepository struct {
	Repository[entity.Post]
}

func NewPostRepository() *PostRepository {
	return &PostRepository{}
}
