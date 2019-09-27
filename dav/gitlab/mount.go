package gitlab

import "net/http"

// Mount todo
type Mount interface {
	Mount(http.ResponseWriter, *http.Request) error
}
