package gitlab

import "net/http"

// DAV todo
type DAV interface {
	Head(http.ResponseWriter, *http.Request) error
	Get(http.ResponseWriter, *http.Request) error
	Post(http.ResponseWriter, *http.Request) error
	Put(http.ResponseWriter, *http.Request) error
	Delete(http.ResponseWriter, *http.Request) error
	Options(http.ResponseWriter, *http.Request) error
	Copy(http.ResponseWriter, *http.Request) error
	Move(http.ResponseWriter, *http.Request) error
	Mkcol(http.ResponseWriter, *http.Request) error
	Lock(http.ResponseWriter, *http.Request) error
	Unlock(http.ResponseWriter, *http.Request) error
	Propfind(http.ResponseWriter, *http.Request) error
	Proppatch(http.ResponseWriter, *http.Request) error
}
