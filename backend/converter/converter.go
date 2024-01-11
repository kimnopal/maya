package converter

import "github.com/kimnopal/maya/model"

func ToWebResponse(statusCode int, message string, data any) *model.WebResponse {
	return &model.WebResponse{
		Status:  statusCode,
		Message: message,
		Data:    data,
	}
}
