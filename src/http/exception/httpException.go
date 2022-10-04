package exception

type HttpException struct {
	Message    string
	StatusCode int
	Headers    map[string]interface{}
}

func (h HttpException) Error() string {
	return h.Message
}
