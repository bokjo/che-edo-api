package handlers

import (
	"net/http"
)

// NotFoundHandler handler struct
type NotFoundHandler struct{}

//NotFound - returns nortfound message
func (nf *NotFoundHandler) NotFound(w http.ResponseWriter, r *http.Request) {
	errorRespond(w, 404, "EDO API is not supporting the specified route")
}
