package transport

import "github.com/LeonidStefanov/HTTP_Service/pkg/models"

func CreateError(err error, status string) *models.ErrorRecponse {
	return &models.ErrorRecponse{
		Error:  err.Error(),
		Status: status,
	}
}

func CreateResponse(info string, status string) *models.Response {
	return &models.Response{
		Status: status,
		Info:   info,
	}
}
