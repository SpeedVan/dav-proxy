package dav

import "net/http"

// Method todo
type Method interface {
	ReadonlyMethod
	// Head(http.ResponseWriter, *http.Request)
	// Get(http.ResponseWriter, *http.Request)
	// Propfind(http.ResponseWriter, *http.Request)
	// Options(http.ResponseWriter, *http.Request)
	Post(http.ResponseWriter, *http.Request)
	Put(http.ResponseWriter, *http.Request)
	Delete(http.ResponseWriter, *http.Request)
	Copy(http.ResponseWriter, *http.Request)
	Move(http.ResponseWriter, *http.Request)
	Mkcol(http.ResponseWriter, *http.Request)
	Lock(http.ResponseWriter, *http.Request)
	Unlock(http.ResponseWriter, *http.Request)
	Proppatch(http.ResponseWriter, *http.Request)
}
