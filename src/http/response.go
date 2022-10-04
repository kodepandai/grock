package http

type Response struct {
	Original interface{}
	Status   int
	Headers  map[string]interface{}
}

type Responsable interface {
	ToResponse() Response
}
