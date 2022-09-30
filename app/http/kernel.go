package http

import "grock/src/foundation/http"

var Kernel *http.Kernel

func init() {
	httpKernel := http.Kernel{}
	Kernel = &httpKernel
}
