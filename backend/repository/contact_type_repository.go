package repository

import "github.com/kimnopal/maya/entity"

type ContactTypeRepository struct {
	Repository[entity.ContactType]
}

func NewContactTypeRepository() *ContactTypeRepository {
	return &ContactTypeRepository{}
}
