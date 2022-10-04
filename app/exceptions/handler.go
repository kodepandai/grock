package exceptions

import (
	"grock/app/http/response"
	"grock/src/foundation/exception"
	httException "grock/src/http/exception"
	"net/http"
)

var Handler exception.Handler

func init() {
	Handler = exception.Handler{
		Register: func(h *exception.Handler) {
			h.Renderable(func(e interface{}) interface{} {
				switch v := e.(type) {
				case httException.HttpException:
					return response.JsonResponse{
						Code:    v.StatusCode,
						Message: v.Error(),
					}
				case error:
					return response.JsonResponse{
						Code:    http.StatusBadRequest,
						Message: v.Error(),
					}
				default:
					return response.JsonResponse{
						Code:    http.StatusInternalServerError,
						Message: "Something Wrong",
					}
				}
			})
		},
	}
}
