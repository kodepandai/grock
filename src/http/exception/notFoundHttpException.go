package exception

import "net/http"

func NotFoundHttpException(message string) HttpException {
	return HttpException{Message: message, StatusCode: http.StatusNotFound}
}
