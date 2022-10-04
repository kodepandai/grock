package response

import (
	"grock/src/http"
	netHttp "net/http"
)

func Make(content interface{}, status int, headers map[string]interface{}) http.Response {
	if status == 0 {
		status = netHttp.StatusOK
	}
	if headers == nil {
		headers = map[string]interface{}{}
	}
	return http.Response{Original: content, Status: status, Headers: headers}
}
