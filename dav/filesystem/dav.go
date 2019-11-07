package filesystem

import (
	"net/http"

	"github.com/SpeedVan/go-common/app/web"
	"github.com/SpeedVan/go-common/log/common"

	"golang.org/x/net/webdav"

	"github.com/SpeedVan/go-common/config"
	"github.com/SpeedVan/go-common/log"
)

// DAV todo
type DAV struct {
	Logger log.Logger
	web.Controller
	Handler   http.Handler
	FS        webdav.FileSystem
	Address   string
	URLPrefix string
}

// New todo
func New(cfg config.Config, logger log.Logger) *DAV {
	if logger == nil {
		logger = common.NewCommon(log.Debug)
	}
	prefix := "/" + cfg.Get("NAME") + "/"
	dir := cfg.Get("FILESYSTEM_DIR")

	h := &webdav.Handler{
		Prefix:     prefix,
		FileSystem: webdav.Dir(dir),
		LockSystem: webdav.NewMemLS(),
		Logger: func(r *http.Request, err error) {
			if err != nil {
				logger.Error(err.Error())
			}
		},
	}

	// h2 := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	// 	fmt.Println(r.Method + " " + r.URL.Path + " " + r.URL.RawQuery)
	// 	fmt.Println(r.Header)
	// 	bs, err := ioutil.ReadAll(r.Body)
	// 	if err != nil {
	// 		fmt.Println(err.Error())
	// 	}
	// 	fmt.Println(string(bs))
	// 	// u, p, ok := r.BasicAuth()
	// 	// if !(ok == true && u == wd.Config.WebDav.Username && p == wd.Config.WebDav.Password) {
	// 	// 	w.Header().Set("WWW-Authenticate", `Basic realm="davfs"`)
	// 	// 	http.Error(w, "Unauthorized", http.StatusUnauthorized)
	// 	// 	return
	// 	// }
	// 	h.ServeHTTP(&rwWrapper{rw: w}, r)
	// })

	return &DAV{
		Logger:    logger,
		Handler:   h,
		URLPrefix: prefix,
	}
}

func (s *DAV) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.Handler.ServeHTTP(w, r)
}

// GetRoute todo
func (s *DAV) GetRoute() web.RouteMap {
	return web.RouteMap{
		s.URLPrefix + "{_dummy:.*}": s.Handler,
	}
}

// // NewHandleFunc todo
// func NewHandleFunc(name string, config config.Config) (string, func(http.ResponseWriter, *http.Request)) {

// 	rootPath := "/" + name + "/"

// 	dir := config.Get("FILESYSTEM_DIR")

// 	h := &webdav.Handler{
// 		Prefix:     rootPath,
// 		FileSystem: webdav.Dir(dir),
// 		LockSystem: webdav.NewMemLS(),
// 		Logger: func(r *http.Request, err error) {
// 			if err != nil {
// 				log.Println(err)
// 			}
// 		},
// 	}

// 	h2 := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 		fmt.Println(r.Method + " " + r.URL.Path + " " + r.URL.RawQuery)
// 		fmt.Println(r.Header)
// 		bs, err := ioutil.ReadAll(r.Body)
// 		if err != nil {
// 			fmt.Println(err.Error())
// 		}
// 		fmt.Println(string(bs))
// 		// u, p, ok := r.BasicAuth()
// 		// if !(ok == true && u == wd.Config.WebDav.Username && p == wd.Config.WebDav.Password) {
// 		// 	w.Header().Set("WWW-Authenticate", `Basic realm="davfs"`)
// 		// 	http.Error(w, "Unauthorized", http.StatusUnauthorized)
// 		// 	return
// 		// }
// 		h.ServeHTTP(&rwWrapper{rw: w}, r)
// 	})
// 	return rootPath + `{_dummy:.*}`, h2
// }

// type rwWrapper struct {
// 	rw http.ResponseWriter
// }

// func (s *rwWrapper) Header() http.Header {
// 	return s.rw.Header()
// }

// func (s *rwWrapper) Write(b []byte) (int, error) {
// 	println(string(b))
// 	return s.rw.Write(b)
// }

// func (s *rwWrapper) WriteHeader(statusCode int) {
// 	println(statusCode)
// 	s.rw.WriteHeader(statusCode)
// }

// // Run todo
// func (s *DAV) Run() error {

// 	return http.ListenAndServe(s.Address, s.Handler)
// }
