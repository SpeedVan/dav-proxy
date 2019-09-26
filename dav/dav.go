package dav

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"golang.org/x/net/webdav"

	"github.com/SpeedVan/go-common/config"
	// "github.com/SpeedVan/go-common/log"
)

// DAV todo
type DAV struct {
	Handler *mux.Router
	Log     log.Logger
	FS      webdav.FileSystem
}

// New todo
func New(path string) *DAV {

	h := &webdav.Handler{
		Prefix:     "/dav",
		FileSystem: webdav.Dir(path),
		LockSystem: webdav.NewMemLS(),
		Logger: func(r *http.Request, err error) {
			if err != nil {
				log.Println(err)
			}
		},
	}

	h2 := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// u, p, ok := r.BasicAuth()
		// if !(ok == true && u == wd.Config.WebDav.Username && p == wd.Config.WebDav.Password) {
		// 	w.Header().Set("WWW-Authenticate", `Basic realm="davfs"`)
		// 	http.Error(w, "Unauthorized", http.StatusUnauthorized)
		// 	return
		// }
		h.ServeHTTP(w, r)
	})

	// h3 := handlers.MethodHandler{
	// 	"GET": h2,
	// }
	router := mux.NewRouter()
	router.HandleFunc("/dav", h2)
	return &DAV{
		Handler: router,
	}
}

// Run todo
func (s *DAV) Run(config config.Config) error {

	return http.ListenAndServe(config.Get("address"), s.Handler)
}
