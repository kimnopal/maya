package repository

import "github.com/kimnopal/maya/entity"

type MajorRepository struct {
	Repository[entity.Major]
}

func NewMajorRepository() *MajorRepository {
	return &MajorRepository{}
}
