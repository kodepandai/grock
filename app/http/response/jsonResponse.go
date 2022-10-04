package response

import (
	"grock/src/http"
	"grock/src/http/response"
)

type JsonResponse struct {
	Success bool        `json:"success"`
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

/* implement Responsable interface */
func (j JsonResponse) ToResponse() http.Response {
	return response.Make(j, j.Code, nil)
}
