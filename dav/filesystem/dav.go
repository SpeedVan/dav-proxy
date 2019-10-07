package filesystem

import (
	"fmt"
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

// NewHandleFunc todo
func NewHandleFunc(rootPath string, config config.Config) (string, func(http.ResponseWriter, *http.Request)) {

	dir := config.Get("FILESYSTEM_DIR")

	h := &webdav.Handler{
		Prefix:     rootPath,
		FileSystem: webdav.Dir(dir),
		LockSystem: webdav.NewMemLS(),
		Logger: func(r *http.Request, err error) {
			if err != nil {
				log.Println(err)
			}
		},
	}

	h2 := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println(r.Method + " " + r.URL.Path)
		// u, p, ok := r.BasicAuth()
		// if !(ok == true && u == wd.Config.WebDav.Username && p == wd.Config.WebDav.Password) {
		// 	w.Header().Set("WWW-Authenticate", `Basic realm="davfs"`)
		// 	http.Error(w, "Unauthorized", http.StatusUnauthorized)
		// 	return
		// }
		h.ServeHTTP(w, r)
	})
	return rootPath + `{_dummy:.*}`, h2
}

// Run todo
func (s *DAV) Run() error {

	return http.ListenAndServe(s.Address, s.Handler)
}
