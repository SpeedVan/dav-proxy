package filesystem

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
	Address string
}

// New todo
func New(config config.Config) (*DAV, error) {

	path := config.Get("WEBDAV_PATH")

	h := &webdav.Handler{
		Prefix:     "/dav/",
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
	router.HandleFunc(`/dav/{_dummy:.*}`, h2)
	return &DAV{
		Handler: router,
		Address: config.Get("address"),
	}, nil
}

// Run todo
func (s *DAV) Run() error {

	return http.ListenAndServe(s.Address, s.Handler)
}
