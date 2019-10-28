package dav

import "net/http"

// ReadonlyMethod todo
type ReadonlyMethod interface {
	Head(http.ResponseWriter, *http.Request)
	Get(http.ResponseWriter, *http.Request)
	Options(http.ResponseWriter, *http.Request)
	Propfind(http.ResponseWriter, *http.Request)
}
