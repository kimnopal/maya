package converter

import (
	"github.com/kimnopal/maya/entity"
	"github.com/kimnopal/maya/model"
)

func ContactTypeEntityToResponse(contactType *entity.ContactType) *model.ContactTypeResponse {
	return &model.ContactTypeResponse{
		ID:   contactType.ID,
		Name: contactType.Name,
	}
}

func ContactTypeCreateRequestToEntity(createRequest *model.ContactTypeCreateRequest) *entity.ContactType {
	return &entity.ContactType{
		Name: createRequest.Name,
	}
}

func ContactTypeUpdateRequestToEntity(contactType *entity.ContactType, updateRequest *model.ContactTypeUpdateRequest) *entity.ContactType {
	contactType.Name = updateRequest.Name

	return contactType
}
