package exception

import (
	grockHttp "grock/src/http"
	"grock/src/http/response"
	"net/http"
)

type RenderUsing func(e interface{}) interface{}
type Handler struct {
	Register    func(*Handler)
	renderUsing RenderUsing
}

func (h *Handler) Render(e error) grockHttp.Response {
	res := h.renderUsing(e)
	switch val := res.(type) {
	case grockHttp.Response:
		return val
	case grockHttp.Responsable:
		return val.ToResponse()
	default:
		return response.Make(res, http.StatusInternalServerError, nil)
	}
}

func (h *Handler) Renderable(renderUsing func(e interface{}) interface{}) {
	h.renderUsing = renderUsing
}
