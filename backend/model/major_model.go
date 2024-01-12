package model

type MajorResponse struct {
	ID        uint64           `json:"id"`
	Name      string           `json:"name"`
	CreatedAt uint64           `json:"created_at"`
	UpdatedAt uint64           `json:"updated_at"`
	Faculty   *FacultyResponse `json:"faculty,omitempty"`
}

type MajorCreateRequest struct {
	Name      string `json:"name" validate:"required,max=100"`
	FacultyID uint64 `json:"faculty_id" validate:"required"`
}

type MajorUpdateRequest struct {
	ID        uint64 `json:"id" validate:"required"`
	Name      string `json:"name" validate:"required,max=100"`
	FacultyID uint64 `json:"faculty_id" validate:"required"`
}

type MajorDeleteRequest struct {
	ID uint64 `json:"id"`
}

type MajorGetRequest struct {
	ID uint64 `json:"id"`
}
